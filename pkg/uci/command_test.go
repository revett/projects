package uci_test

import (
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

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
