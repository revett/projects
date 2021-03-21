package uci_test

import (
	"testing"
	"time"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestGoCommand(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		engineOutput []string
		expectError  bool
	}{
		"Happy": {
			engineOutput: []string{
				"info string NNUE evaluation using nn-62ef826d1a6d.nnue enabled",
				"info depth 1 seldepth 1 multipv 1 score cp 29 nodes 20 nps 20000 tbhits 0 time 1 pv d2d4",
				"info depth 2 seldepth 2 multipv 1 score cp 89 nodes 42 nps 4666 tbhits 0 time 9 pv d2d4 a7a6",
				"bestmove d2d4 ponder a7a6",
			},
		},
		"MalformedBestMoveLine": {
			engineOutput: []string{
				"info string NNUE evaluation using nn-62ef826d1a6d.nnue enabled",
				"info depth 1 seldepth 1 multipv 1 score cp 29 nodes 20 nps 20000 tbhits 0 time 1 pv d2d4",
				"info depth 2 seldepth 2 multipv 1 score cp 89 nodes 42 nps 4666 tbhits 0 time 9 pv d2d4 a7a6",
				"bestmove d2d4 ponder",
			},
			expectError: true,
		},
		"Timeout": {
			expectError: true,
		},
	}

	for n, tc := range tests {
		tc := tc

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			mc := mockCommander{
				out: tc.engineOutput,
			}

			e, err := uci.NewEngine(
				mc.Command,
				mockEnginePath,
				uci.WithCommandTimeout(100*time.Millisecond),
			)
			assert.NoError(t, err)

			err = e.Run(uci.GoCommand())
			assert.Equal(t, tc.expectError, err != nil)

			err = e.Close()
			assert.NoError(t, err)
		})
	}
}

func TestGoCommandString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		gc   uci.Command
		want string
	}{
		"EmptyOptions": {
			gc:   uci.GoCommand(),
			want: "go",
		},
		"Multiple": {
			gc: uci.GoCommand(
				uci.WithSearchMoves("e2e4", "e7e5"),
				uci.WithDepth(10),
				uci.WithMoveTime(1000),
			),
			want: "go depth 10 movetime 1000 searchmoves e2e4 e7e5",
		},
		"WithDepth": {
			gc:   uci.GoCommand(uci.WithDepth(10)),
			want: "go depth 10",
		},
		"WithInfinite": {
			gc:   uci.GoCommand(uci.WithInfinite),
			want: "go infinite",
		},
		"WithMoveTime": {
			gc:   uci.GoCommand(uci.WithMoveTime(1000)),
			want: "go movetime 1000",
		},
		"WithSearchMoves": {
			gc:   uci.GoCommand(uci.WithSearchMoves("e2e4", "e7e5")),
			want: "go searchmoves e2e4 e7e5",
		},
	}

	for n, tc := range tests {
		tc := tc

		t.Run(n, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, tc.gc.String())
		})
	}
}

func TestIsReadyCommand(t *testing.T) {
	mc := mockCommander{
		out: []string{
			"readyok",
		},
	}

	e, err := uci.NewEngine(mc.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(uci.IsReadyCommand())
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestPositionCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(uci.PositionCommand())
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestSetOptionCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(uci.SetOptionCommand("threads", "2"))
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

	err = e.Run(uci.UCICommand())
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestUCINewGameCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(uci.UCINewGameCommand())
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}
