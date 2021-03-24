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
				uci.SearchMoves("e2e4", "e7e5"),
				uci.Depth(10),
				uci.MoveTime(1000),
			),
			want: "go depth 10 movetime 1000 searchmoves e2e4 e7e5",
		},
		"Depth": {
			c:    uci.GoCommand(uci.Depth(10)),
			want: "go depth 10",
		},
		"Infinite": {
			c:    uci.GoCommand(uci.Infinite),
			want: "go infinite",
		},
		"MoveTime": {
			c:    uci.GoCommand(uci.MoveTime(1000)),
			want: "go movetime 1000",
		},
		"SearchMoves": {
			c:    uci.GoCommand(uci.SearchMoves("e2e4", "e7e5")),
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
				uci.Moves("e2e4", "e7e5"),
				uci.FEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12"),
			),
			want: "position fen r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12 moves e2e4 e7e5",
		},
		"FEN": {
			c: uci.PositionCommand(
				uci.FEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12"),
			),
			want: "position fen r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12",
		},
		"Moves": {
			c: uci.PositionCommand(
				uci.Moves("e2e4", "e7e5"),
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

func TestSetOptionCommandString(t *testing.T) {
	name := "threads"
	value := "2"

	want := fmt.Sprintf("setoption name %s value %s", name, value)
	c := uci.SetOptionCommand(name, value)

	assert.Equal(t, want, c.String())
}

func TestUCICommandString(t *testing.T) {
	assert.Equal(t, "uci", uci.UCICommand().String())
}

func TestUCINewGameCommandString(t *testing.T) {
	assert.Equal(t, "ucinewgame", uci.UCINewGameCommand().String())
}
