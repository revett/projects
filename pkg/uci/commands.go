package uci

import (
	"fmt"
	"strings"
)

const defaultSearchDepth = 10

// GoCommand is TODO.
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

func (g goCommand) String() string {
	return fmt.Sprintf("go depth %d", g.depth)
}

// WithDepth is TODO.
func WithDepth(i int) func(*goCommand) {
	return func(c *goCommand) {
		c.depth = i
	}
}

// IsReadyCommand is TODO.
func IsReadyCommand() Command {
	return isReadyCommand{}
}

type isReadyCommand struct{}

func (i isReadyCommand) processOutput(e *Engine) error {
	_, err := e.readUntil("readyok")
	return err
}

func (i isReadyCommand) String() string {
	return "isready"
}

// PositionCommand is TODO.
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

func (p positionCommand) String() string {
	if len(p.moves) == 0 {
		return fmt.Sprintf("position fen %s", p.fen)
	}

	return fmt.Sprintf(
		"position fen %s moves %s", p.fen, strings.Join(p.moves, " "),
	)
}

// WithFEN is TODO.
func WithFEN(s string) func(*positionCommand) {
	return func(c *positionCommand) {
		c.fen = s
	}
}

// WithMoves is TODO.
func WithMoves(s ...string) func(*positionCommand) {
	return func(c *positionCommand) {
		c.moves = s
	}
}

// UCICommand is TODO.
func UCICommand() Command {
	return uciCommand{}
}

type uciCommand struct{}

func (u uciCommand) processOutput(e *Engine) error {
	_, err := e.readUntil("uciok")
	return err
}

func (u uciCommand) String() string {
	return "uci"
}

// UCINewGameCommand is TODO.
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
