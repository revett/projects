package uci

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

// Engine holds the properties required to communicate with a UCI-compatible
// chess engine executable.
type Engine struct {
	cmd *exec.Cmd
	in  *io.PipeWriter
	out *io.PipeReader
}

// NewEngine returns an Engine.
func NewEngine(c commander, p string) (*Engine, error) {
	rIn, wIn := io.Pipe()
	rOut, wOut := io.Pipe()

	cmd := c.Command(p)
	cmd.Stdin = rIn
	cmd.Stdout = wOut

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "failed to start command")
	}

	return &Engine{
		cmd: cmd,
		in:  wIn,
		out: rOut,
	}, nil
}

// Stop ends the chess engine executable.
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

// IsReady checks if the engine is ready for a command.
func (e Engine) IsReady() (bool, error) {
	err := e.sendCommand("isready")
	if err != nil {
		return false, err
	}

	l, err := e.readUntil("readyok")
	if err != nil {
		return false, err
	}

	if l[len(l)-1] != "readyok" {
		return false, errors.New("unknown response from engine")
	}

	return true, nil
}

func (e Engine) readUntil(s string) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(e.out)

	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)

		if l == s {
			break
		}
	}

	if scanner.Err() != nil {
		return nil, errors.Wrap(scanner.Err(), "error reading output from engine")
	}

	return lines, nil
}

func (e Engine) sendCommand(s string) error {
	_, err := fmt.Fprintln(e.in, s)
	if err != nil {
		return errors.Wrap(err, "error creating command to send")
	}

	return nil
}
