package uci_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestEngineIsReady(t *testing.T) {
	e, err := uci.NewEngine("/path/to/engine", fakeExecContext)
	assert.NoError(t, err)

	ready, err := e.IsReady()
	assert.NoError(t, err)
	assert.True(t, ready)

	err = e.Stop()
	assert.NoError(t, err)
}

// TestProcess is a function acts as a subsitute for an actual shell command.
// GO_TEST_PROCESS flag ensures that if it is called as part of the test suite,
// it is skipped.
func TestProcess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	fmt.Fprintln(
		os.Stdout, "Stockfish 13 by the Stockfish developers (see AUTHORS file)",
	)
	fmt.Fprintln(os.Stdout, "readyok")
}

// fakeExecCommandSuccess is a function that initialises a new exec.Cmd, one
// which will simply calls TestProcess rather than the command it is provided.
func fakeExecContext(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}
