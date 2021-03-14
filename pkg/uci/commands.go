package uci

import (
	"fmt"
	"strings"
)

const defaultSearchDepth = 10

// GoCommand is used to run the `go` UCI command.
func GoCommand(opts ...func(*goCommand)) Command {
	g := goCommand{}
	if len(opts) == 0 {
		g.depth = defaultSearchDepth
		return g
	}

	for _, o := range opts {
		o(&g)
	}

	return g
}

type goCommand struct {
	depth int
}

func (g goCommand) processOutput(e *Engine) error {
	_, err := e.readUntil("bestmove")
	return err
}

// String implements the Command interface.
func (g goCommand) String() string {
	return fmt.Sprintf("go depth %d", g.depth)
}

// WithDepth is a functional option that configures the GoCommand to search to
// a certain depth.
func WithDepth(i int) func(*goCommand) {
	return func(c *goCommand) {
		c.depth = i
	}
}

// IsReadyCommand is used to run the `isready` UCI command.
func IsReadyCommand() Command {
	return isReadyCommand{}
}

type isReadyCommand struct{}

func (i isReadyCommand) processOutput(e *Engine) error {
	_, err := e.readUntil("readyok")
	return err
}

// String implements the Command interface.
func (i isReadyCommand) String() string {
	return "isready"
}

// PositionCommand is used to run the `position` UCI command.
func PositionCommand(opts ...func(*positionCommand)) Command {
	p := positionCommand{
		fen: StartingPosition,
	}

	for _, o := range opts {
		o(&p)
	}

	return p
}

type positionCommand struct {
	fen   string
	moves []string
}

func (p positionCommand) processOutput(e *Engine) error {
	return nil
}

// String implements the Command interface.
func (p positionCommand) String() string {
	if len(p.moves) == 0 {
		return fmt.Sprintf("position fen %s", p.fen)
	}

	return fmt.Sprintf(
		"position fen %s moves %s", p.fen, strings.Join(p.moves, " "),
	)
}

// WithFEN is a functional option that configures the PositionCommand with a
// specific board position in Forsyth–Edwards Notation (FEN) notation.
func WithFEN(s string) func(*positionCommand) {
	return func(c *positionCommand) {
		c.fen = s
	}
}

// WithMoves is a functional option that configures the PositionCommand with a
// series of moves to play.
func WithMoves(s ...string) func(*positionCommand) {
	return func(c *positionCommand) {
		c.moves = s
	}
}

// UCICommand is used to run the `uci` UCI command.
// nolint:golint
func UCICommand() Command {
	return uciCommand{}
}

type uciCommand struct{}

func (u uciCommand) processOutput(e *Engine) error {
	_, err := e.readUntil("uciok")
	return err
}

// String implements the Command interface.
func (u uciCommand) String() string {
	return "uci"
}

// UCINewGameCommand is used to run the `ucinewgame` UCI command.
// nolint:golint
func UCINewGameCommand() Command {
	return uciCommand{}
}

type uciNewGameCommand struct{}

func (u uciNewGameCommand) processOutput(e *Engine) error {
	return nil
}

func (u uciNewGameCommand) String() string {
	return "ucinewgame"
}
