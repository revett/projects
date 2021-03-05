package uci

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type execContext func(name string, arg ...string) *exec.Cmd

// Engine holds the properties required to communicate with a UCI-compatible
// chess engine executable.
type Engine struct {
	cmd *exec.Cmd
	in  *bufio.Writer
	out *bufio.Reader
}

// NewEngine returns an Engine.
func NewEngine(p string, cmdContext execContext) (*Engine, error) {
	cmd := cmdContext(p)

	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, errors.Wrap(err, "failed to return stdin pipe from command")
	}

	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "failed to return stdout pipe from command")
	}

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "failed to start command")
	}

	return &Engine{
		cmd: cmd,
		in:  bufio.NewWriter(in),
		out: bufio.NewReader(out),
	}, nil
}

// Stop ends the chess engine command.
func (e Engine) Stop() error {
	err := e.sendCommand("quit")
	if err != nil {
		return err
	}

	return e.cmd.Wait()
}

// IsReady checks if the engine is ready for a command
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

	for {
		l, err := e.out.ReadString('\n')
		if err != nil {
			return nil, errors.Wrap(err, "unable to read output from engine")
		}

		l = strings.Trim(l, "\n")
		lines = append(lines, l)

		if l == s {
			break
		}
	}

	return lines, nil
}

func (e Engine) sendCommand(s string) error {
	c := fmt.Sprintf("%s\n", s)
	_, err := e.in.WriteString(c)
	if err != nil {
		return errors.Wrap(err, "error creating command to send")
	}

	err = e.in.Flush()
	if err != nil {
		return errors.Wrap(err, "error sending command to engine")
	}

	return nil
}
