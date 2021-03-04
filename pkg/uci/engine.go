package uci

import (
	"bufio"
	"os/exec"

	"github.com/pkg/errors"
)

// Engine holds the properties required to communicate with a UCI-compatible
// chess engine executable.
type Engine struct {
	cmd *exec.Cmd
	in  *bufio.Writer
	out *bufio.Reader
}

// NewEngine returns an Engine, with any options configured.
func NewEngine(p string) (*Engine, error) {
	cmd := exec.Command(p)

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
