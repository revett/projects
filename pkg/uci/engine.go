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

// Command is an exported function to allow unit tests to monkey patch how the
// program will be executed.
var Command = exec.Command

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

	cmd := Command(p)
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

// Debug is an option for the NewEngine function which logs any commands sent to
// the engine, and all output received.
func Debug(e *Engine) error {
	e.debug = true
	return nil
}

// InitialiseGame is an option for the NewEngine function which tells the engine
// to use UCI and start a new game. It then checks if the engine is ready.
func InitialiseGame(e *Engine) error {
	if err := e.UCI(); err != nil {
		return err
	}

	if err := e.UCINewGame(); err != nil {
		return err
	}

	return e.IsReady()
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

// Go searches for the best move(s).
func (e Engine) Go() error {
	if err := e.sendCommand(goCmd); err != nil {
		return err
	}

	_, err := e.readUntil("bestmove")
	if err != nil {
		return err
	}

	return err
}

// IsReady sends the `isready` command to the engine, to check that it is alive.
func (e Engine) IsReady() error {
	err := e.sendCommand(isReadyCmd)
	if err != nil {
		return err
	}

	_, err = e.readUntil("readyok")

	return err
}

// Position sends the `position` command to the engine with a givin FEN, setting
// the internal board position.
func (e Engine) Position(f string) error {
	return e.sendCommand(positionCmd, f)
}

// UCI sends the `uci` command to the engine, to tell the engine to use UCI.
func (e Engine) UCI() error {
	err := e.sendCommand(uciCmd)
	if err != nil {
		return err
	}

	_, err = e.readUntil("uciok")

	return err
}

// UCINewGame sends the `ucinewgame` command to the engine, to tell the engine
// that the next search command will be from a different game.
func (e Engine) UCINewGame() error {
	return e.sendCommand(uciNewGameCmd)
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
