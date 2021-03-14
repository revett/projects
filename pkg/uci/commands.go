package uci

import (
	"fmt"
	"strings"
)

const (
	goCmd = "go depth 10"
)

// IsReadyCommand is TODO.
func IsReadyCommand() Command {
	return isReadyCommand{}
}

type isReadyCommand struct{}

func (i isReadyCommand) execute(e *Engine) error {
	if err := e.sendCommand(i.string()); err != nil {
		return err
	}

	_, err := e.readUntil("readyok")

	return err
}

func (i isReadyCommand) string() string {
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

func (p positionCommand) execute(e *Engine) error {
	return e.sendCommand(p.string())
}

func (p positionCommand) string() string {
	if len(p.moves) == 0 {
		return fmt.Sprintf("position fen %s", p.fen)
	}

	return fmt.Sprintf(
		"position fen %s moves %s", p.fen, strings.Join(p.moves, " "),
	)
}

// WithFEN is TODO.
func WithFEN(s string) func(*positionCommand) {
	return func(p *positionCommand) {
		p.fen = s
	}
}

// WithMoves is TODO.
func WithMoves(s ...string) func(*positionCommand) {
	return func(p *positionCommand) {
		p.moves = s
	}
}

// UCICommand is TODO.
func UCICommand() Command {
	return uciCommand{}
}

type uciCommand struct{}

func (u uciCommand) execute(e *Engine) error {
	if err := e.sendCommand(u.string()); err != nil {
		return err
	}

	_, err := e.readUntil("uciok")

	return err
}

func (u uciCommand) string() string {
	return "uci"
}

// UCINewGameCommand is TODO.
func UCINewGameCommand() Command {
	return uciCommand{}
}

type uciNewGameCommand struct{}

func (u uciNewGameCommand) execute(e *Engine) error {
	return e.sendCommand(u.string())
}

func (u uciNewGameCommand) string() string {
	return "ucinewgame"
}
