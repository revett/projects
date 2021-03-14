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
	uci.XCommand = mockCommander{}.Command
	e, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestGo(t *testing.T) {
	tests := map[string]struct {
		cmdOutput []string
		want      error
	}{
		"Success": {
			cmdOutput: []string{
				"info string NNUE evaluation using nn-62ef826d1a6d.nnue enabled",
				// nolint:lll
				"info depth 10 seldepth 12 multipv 1 score cp 38 nodes 10144 nps 563555 tbhits 0 time 18 pv e2e4 c7c5 g1f3 e7e6 d2d4 c5d4 f3d4 g8f6",
				"bestmove e2e4 ponder c7c5",
			},
			want: nil,
		},
		"TimeOut": {
			cmdOutput: []string{
				"info string NNUE evaluation using nn-62ef826d1a6d.nnue enabled",
			},
			want: uci.CommandTimeoutError{},
		},
	}

	for n, tc := range tests {
		t.Run(n, func(t *testing.T) {
			m := mockCommander{
				out: tc.cmdOutput,
			}
			uci.XCommand = m.Command

			e, err := uci.NewEngine(
				mockEnginePath, uci.WithCommandTimeout(100*time.Millisecond),
			)
			assert.NoError(t, err)

			err = e.Go()
			assert.IsType(t, tc.want, err)
		})
	}
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
	uci.XCommand = m.Command

	_, err := uci.NewEngine(mockEnginePath, uci.InitialiseGame)
	assert.NoError(t, err)
}

func TestIsReady(t *testing.T) {
	tests := map[string]struct {
		cmdOutput []string
		want      error
	}{
		"Success": {
			cmdOutput: []string{
				"Stockfish 13 by the Stockfish developers (see AUTHORS file)",
				"readyok",
			},
			want: nil,
		},
		"TimeOut": {
			cmdOutput: []string{
				"Stockfish 13 by the Stockfish developers (see AUTHORS file)",
			},
			want: uci.CommandTimeoutError{},
		},
	}

	for n, tc := range tests {
		t.Run(n, func(t *testing.T) {
			m := mockCommander{
				out: tc.cmdOutput,
			}
			uci.XCommand = m.Command

			_, err := uci.NewEngine(
				mockEnginePath, uci.WithCommandTimeout(100*time.Millisecond),
			)
			assert.NoError(t, err)

			// err = e.IsReady()
			// assert.IsType(t, tc.want, err)
		})
	}
}

func TestNewEngine(t *testing.T) {
	uci.XCommand = mockCommander{}.Command
	_, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)
}

func TestPosition(t *testing.T) {
	uci.XCommand = mockCommander{}.Command
	e, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)

	err = e.Position(uci.StartingPosition)
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
	uci.XCommand = m.Command

	e, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)

	err = e.UCI()
	assert.NoError(t, err)
}

func TestUCINewGame(t *testing.T) {
	uci.XCommand = mockCommander{}.Command
	e, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)

	err = e.UCINewGame()
	assert.NoError(t, err)
}

type mockCommander struct {
	out []string
}

func (m mockCommander) Command(s string, a ...string) *exec.Cmd {
	// nolint:gosec
	cmd := exec.Command(os.Args[0])
	out := fmt.Sprintf("TEST_CMD_OUTPUT=%s", strings.Join(m.out, ","))
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
