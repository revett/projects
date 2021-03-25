package uci

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type commander func(name string, arg ...string) *exec.Cmd

const defaultCommandTimeout = 1 * time.Second

// Engine holds the properties required to communicate with a UCI chess engine
// process.
type Engine struct {
	cmd       *exec.Cmd
	logOutput bool
	in        *io.PipeWriter
	out       *io.PipeReader
	timeout   time.Duration
	Results   Results
}

// NewEngine returns an Engine, after starting the chess engine executable at
// the given path. Zero or more functional options can be passed to configure
// the Engine.
func NewEngine(c commander, p string, opts ...func(e *Engine)) (*Engine, error) {
	rIn, wIn := io.Pipe()
	rOut, wOut := io.Pipe()

	cmd := c(p)
	cmd.Stdin = rIn
	cmd.Stdout = wOut

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "failed to start command")
	}

	e := &Engine{
		cmd:     cmd,
		timeout: defaultCommandTimeout,
		in:      wIn,
		out:     rOut,
	}

	for _, o := range opts {
		o(e)
	}

	return e, nil
}

// LogOutput is a functional option for configuring NewEngine so that commands
// sent to an Engine and output from the Engine itself are logged.
func LogOutput(e *Engine) {
	e.logOutput = true
}

// CommandTimeout is a functional option for configuring NewEngine which
// sets the duration a command must complete in, the default is 1 second.
func CommandTimeout(d time.Duration) func(*Engine) {
	return func(e *Engine) {
		e.timeout = d
	}
}

// Close ends the underlying Engine process.
func (e Engine) Close() error {
	if err := e.sendCommand("quit"); err != nil {
		return err
	}

	if err := e.in.Close(); err != nil {
		return err
	}

	if err := e.out.Close(); err != nil {
		return err
	}

	return e.cmd.Process.Kill()
}

// Run sends one or more UCI commands to the engine and processes the engine
// output. The commands are sent to the engine in the order given.
func (e *Engine) Run(cmds ...Command) error {
	for _, c := range cmds {
		if err := e.sendCommand(c.String()); err != nil {
			return err
		}

		if err := c.processOutput(e); err != nil {
			return err
		}
	}

	return nil
}

func (e Engine) readUntil(prefix string) ([]string, error) {
	scanner := bufio.NewScanner(e.out)
	c := make(chan []string, 1)

	go scanOutput(c, scanner, &e, prefix)

	var lines []string

	select {
	case res := <-c:
		lines = res
	case <-time.After(e.timeout):
		return nil, errors.Errorf(
			"timed out after %d seconds, waiting for '%s' response from engine",
			e.timeout,
			prefix,
		)
	}

	if scanner.Err() != nil {
		return nil, errors.Wrap(scanner.Err(), "error reading output from engine")
	}

	return lines, nil
}

func (e Engine) sendCommand(s string, a ...interface{}) error {
	s = fmt.Sprintf(s, a...)

	if e.logOutput {
		log.Printf("> %s", s)
	}

	_, err := fmt.Fprintln(e.in, s)
	if err != nil {
		return errors.Wrap(err, "error creating command to send")
	}

	return nil
}

func scanOutput(c chan []string, scanner *bufio.Scanner, e *Engine, prefix string) {
	var lines []string

	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)

		if e.logOutput {
			log.Println(l)
		}

		if strings.HasPrefix(l, prefix) {
			break
		}
	}

	c <- lines
}
