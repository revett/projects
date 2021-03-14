package uci

import (
	"fmt"
	"strings"
)

const (
	goCmd         = "go depth 10"
	positionCmd   = "position %s"
	uciCmd        = "uci"
	uciNewGameCmd = "ucinewgame"
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
