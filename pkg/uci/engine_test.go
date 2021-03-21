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
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestNewEngine(t *testing.T) {
	_, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)
}

func TestResultsBestMove(t *testing.T) {
	bestMove := "d2d4"
	mc := mockCommander{
		out: []string{
			"info string NNUE evaluation using nn-62ef826d1a6d.nnue enabled",
			"info depth 1 seldepth 1 multipv 1 score cp 29 nodes 20 nps 20000 tbhits 0 time 1 pv d2d4",
			"info depth 2 seldepth 2 multipv 1 score cp 89 nodes 42 nps 4666 tbhits 0 time 9 pv d2d4 a7a6",
			fmt.Sprintf("bestmove %s ponder a7a6", bestMove),
		},
	}

	e, err := uci.NewEngine(mc.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.GoCommand(),
	)
	assert.NoError(t, err)
	assert.Equal(t, bestMove, e.Results.BestMove)

	err = e.Close()
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
