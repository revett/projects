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

// XCommand is an exported function to allow unit tests to monkey patch how the
// program will be executed.
var XCommand = exec.Command

const defaultCommandTimeout = 1 * time.Second

// Engine holds the properties required to communicate with a UCI-compatible
// chess engine executable.
type Engine struct {
	cmd     *exec.Cmd
	debug   bool
	timeout time.Duration
	in      *io.PipeWriter
	out     *io.PipeReader
}

// NewEngine returns an Engine.
func NewEngine(p string, opts ...func(e *Engine) error) (*Engine, error) {
	rIn, wIn := io.Pipe()
	rOut, wOut := io.Pipe()

	cmd := XCommand(p)
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
		if err := o(e); err != nil {
			return nil, err
		}
	}

	return e, nil
}

// LogOutput is an option for the NewEngine function which logs any commands
// sent to the engine, and all output received.
func LogOutput(e *Engine) error {
	e.debug = true
	return nil
}

// WithCommandTimeout is an option for the NewEngine function which sets the
// duration a command must complete in.
func WithCommandTimeout(d time.Duration) func(*Engine) error {
	return func(e *Engine) error {
		e.timeout = d
		return nil
	}
}

// Close ends the chess engine process.
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

// Run is TODO.
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

func (e Engine) readUntil(s string) ([]string, error) {
	scanner := bufio.NewScanner(e.out)
	c := make(chan []string, 1)

	go func() {
		var lines []string

		for scanner.Scan() {
			l := scanner.Text()
			lines = append(lines, l)

			if e.debug {
				log.Println(l)
			}

			if strings.HasPrefix(l, s) {
				break
			}
		}

		c <- lines
	}()

	var lines []string

	select {
	case res := <-c:
		lines = res
	case <-time.After(e.timeout):
		return nil, CommandTimeoutError{
			duration: 1,
			response: s,
		}
	}

	if scanner.Err() != nil {
		return nil, errors.Wrap(scanner.Err(), "error reading output from engine")
	}

	return lines, nil
}

func (e Engine) sendCommand(s string, a ...interface{}) error {
	s = fmt.Sprintf(s, a...)

	if e.debug {
		log.Printf("> %s", s)
	}

	_, err := fmt.Fprintln(e.in, s)
	if err != nil {
		return errors.Wrap(err, "error creating command to send")
	}

	return nil
}
