package uci_test

import (
	"fmt"
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestGoCommand(t *testing.T) {
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

func TestIsReadyCommand(t *testing.T) {
	mc := mockCommander{
		out: []string{
			"readyok",
		},
	}

	e, err := uci.NewEngine(mc.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.IsReadyCommand(),
	)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestPositionCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.PositionCommand(),
	)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestSetOptionCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.SetOptionCommand("threads", "2"),
	)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestUCICommand(t *testing.T) {
	mc := mockCommander{
		out: []string{
			"id name Stockfish 13",
			"id author the Stockfish developers (see AUTHORS file)",
			"",
			"option name Debug Log File type string default",
			"option name Contempt type spin default 24 min -100 max 100",
			"option name Threads type spin default 1 min 1 max 512",
			"uciok",
		},
	}

	e, err := uci.NewEngine(mc.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.UCICommand(),
	)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestUCINewGameCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.UCINewGameCommand(),
	)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}
