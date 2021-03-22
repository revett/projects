package uci_test

import (
	"fmt"
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestGoCommandString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		c    uci.Command
		want string
	}{
		"EmptyOptions": {
			c:    uci.GoCommand(),
			want: "go",
		},
		"MultipleOptions": {
			c: uci.GoCommand(
				uci.WithSearchMoves("e2e4", "e7e5"),
				uci.WithDepth(10),
				uci.WithMoveTime(1000),
			),
			want: "go depth 10 movetime 1000 searchmoves e2e4 e7e5",
		},
		"WithDepth": {
			c:    uci.GoCommand(uci.WithDepth(10)),
			want: "go depth 10",
		},
		"WithInfinite": {
			c:    uci.GoCommand(uci.WithInfinite),
			want: "go infinite",
		},
		"WithMoveTime": {
			c:    uci.GoCommand(uci.WithMoveTime(1000)),
			want: "go movetime 1000",
		},
		"WithSearchMoves": {
			c:    uci.GoCommand(uci.WithSearchMoves("e2e4", "e7e5")),
			want: "go searchmoves e2e4 e7e5",
		},
	}

	for n, tc := range tests {
		tc := tc

		t.Run(n, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, tc.c.String())
		})
	}
}

func TestIsReadyCommandString(t *testing.T) {
	assert.Equal(t, "isready", uci.IsReadyCommand().String())
}

func TestPositionCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(uci.PositionCommand())
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestPositionCommandString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		c    uci.Command
		want string
	}{
		"EmptyOptions": {
			c:    uci.PositionCommand(),
			want: "position",
		},
		"MultipleOptions": {
			c: uci.PositionCommand(
				uci.WithMoves("e2e4", "e7e5"),
				uci.WithFEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12"),
			),
			want: "position fen r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12 moves e2e4 e7e5",
		},
		"WithFEN": {
			c: uci.PositionCommand(
				uci.WithFEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12"),
			),
			want: "position fen r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12",
		},
		"WithMoves": {
			c: uci.PositionCommand(
				uci.WithMoves("e2e4", "e7e5"),
			),
			want: "position moves e2e4 e7e5",
		},
	}

	for n, tc := range tests {
		tc := tc

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.c.String())
		})
	}
}

func TestSetOptionCommand(t *testing.T) {
	e, err := uci.NewEngine(mockCommander{}.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(uci.SetOptionCommand("threads", "2"))
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestSetOptionCommandString(t *testing.T) {
	name := "threads"
	value := "2"

	want := fmt.Sprintf("setoption name %s value %s", name, value)
	c := uci.SetOptionCommand(name, value)

	assert.Equal(t, want, c.String())
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
