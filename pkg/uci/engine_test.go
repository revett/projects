package uci_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestEngineIsReady(t *testing.T) {
	m := mockCommander{
		out: []string{
			"Stockfish 13 by the Stockfish developers (see AUTHORS file)",
			"readyok",
		},
	}

	e, err := uci.NewEngine("/path/to/engine", m)
	assert.NoError(t, err)

	ready, err := e.IsReady()
	assert.NoError(t, err)
	assert.True(t, ready)

	err = e.Stop()
	assert.NoError(t, err)
}

type mockCommander struct {
	out []string
}

func (m mockCommander) Command(s string, a ...string) *exec.Cmd {
	out := fmt.Sprintf("TEST_CMD_OUTPUT=%s", strings.Join(m.out, ","))
	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "TEST_MAIN=1", out)
	return cmd
}

func TestMain(m *testing.M) {
	if os.Getenv("TEST_MAIN") != "1" {
		os.Exit(m.Run())
	}

	l := strings.Split(os.Getenv("TEST_CMD_OUTPUT"), ",")
	for _, s := range l {
		fmt.Fprintln(os.Stdout, s)
	}

	os.Exit(0)
}
