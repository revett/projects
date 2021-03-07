package uci_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

const mockEnginePath = "/path/to/engine"

func TestClose(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}, mockEnginePath)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestInitialiseGame(t *testing.T) {
	m := mockCommander{
		out: []string{
			"Stockfish 13 by the Stockfish developers (see AUTHORS file)",
			"id name Stockfish 13",
			"id author the Stockfish developers (see AUTHORS file)",
			"uciok",
			"readyok",
		},
	}

	_, err := uci.NewEngine(m, mockEnginePath, uci.InitialiseGame)
	assert.NoError(t, err)
}

func TestIsReady(t *testing.T) {
	m := mockCommander{
		out: []string{
			"Stockfish 13 by the Stockfish developers (see AUTHORS file)",
			"readyok",
		},
	}

	e, err := uci.NewEngine(m, mockEnginePath)
	assert.NoError(t, err)

	err = e.IsReady()
	assert.NoError(t, err)
}

func TestNewEngine(t *testing.T) {
	_, err := uci.NewEngine(mockCommander{}, mockEnginePath)
	assert.NoError(t, err)
}

func TestUCI(t *testing.T) {
	m := mockCommander{
		out: []string{
			"id name Stockfish 13",
			"id author the Stockfish developers (see AUTHORS file)",
			"uciok",
		},
	}

	e, err := uci.NewEngine(m, mockEnginePath)
	assert.NoError(t, err)

	err = e.UCI()
	assert.NoError(t, err)
}

func TestUCINewGame(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}, mockEnginePath)
	assert.NoError(t, err)

	err = e.UCINewGame()
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
		time.Sleep(1 * time.Millisecond)
	}

	os.Exit(0)
}
