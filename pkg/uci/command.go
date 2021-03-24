package uci

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Command interface {
	fmt.Stringer
	processOutput(*Engine) error
}

const requiredBestMoveOutputParts = 4

// GoCommand is used to run the "go" UCI command.
func GoCommand(opts ...func(*goCommand)) Command {
	g := goCommand{}

	for _, o := range opts {
		o(&g)
	}

	return g
}

type goCommand struct {
	depth       int
	infinite    bool
	movetime    int
	searchmoves []string
}

func (g goCommand) processOutput(e *Engine) error {
	l, err := e.readUntil("bestmove")
	if err != nil {
		return err
	}

	lastLine := l[len(l)-1]

	parts := strings.Split(lastLine, " ")
	if len(parts) != requiredBestMoveOutputParts {
		return errors.New("malformed last line from go command")
	}

	e.Results = Results{
		BestMove: parts[1],
	}

	return nil
}

// String implements the Command interface.
func (g goCommand) String() string {
	p := []string{}

	if g.depth > 0 {
		p = append(p, "depth", strconv.Itoa(g.depth))
	}

	if g.infinite {
		p = append(p, "infinite")
	}

	if g.movetime > 0 {
		p = append(p, "movetime", strconv.Itoa(g.movetime))
	}

	if len(g.searchmoves) > 0 {
		p = append(p, "searchmoves", strings.Join(g.searchmoves, " "))
	}

	if len(p) == 0 {
		return "go"
	}

	return fmt.Sprintf("go %s", strings.Join(p, " "))
}

// Depth is a functional option that configures the GoCommand to search to
// a certain depth.
func Depth(i int) func(*goCommand) {
	return func(c *goCommand) {
		c.depth = i
	}
}

// Infinite is a functional option that configures the GoCommand to continue
// searching until the "stop" UCI command is sent.
func Infinite(c *goCommand) {
	c.infinite = true
}

// MoveTime is a functional option that configures the GoCommand to search
// for a given period of time, in milliseconds.
func MoveTime(i int) func(*goCommand) {
	return func(c *goCommand) {
		c.movetime = i
	}
}

// SearchMoves is a functional option that restricts the GoCommand to only
// search using a set of defined moves.
func SearchMoves(s ...string) func(*goCommand) {
	return func(c *goCommand) {
		c.searchmoves = s
	}
}

// IsReadyCommand is used to run the "isready" UCI command.
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

// PositionCommand is used to run the "position" UCI command.
func PositionCommand(opts ...func(*positionCommand)) Command {
	p := positionCommand{}

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
	parts := []string{}

	if p.fen != "" {
		parts = append(parts, "fen", p.fen)
	}

	if len(p.moves) > 0 {
		parts = append(parts, "moves", strings.Join(p.moves, " "))
	}

	if len(parts) == 0 {
		return "position"
	}

	return fmt.Sprintf("position %s", strings.Join(parts, " "))
}

// FEN is a functional option that configures the PositionCommand with a
// specific board position in Forsyth–Edwards Notation (FEN) notation.
func FEN(s string) func(*positionCommand) {
	return func(c *positionCommand) {
		c.fen = s
	}
}

// Moves is a functional option that configures the PositionCommand with a
// series of moves to play.
func Moves(s ...string) func(*positionCommand) {
	return func(c *positionCommand) {
		c.moves = s
	}
}

// SetOptionCommand is used to run the "setoption" UCI command.
func SetOptionCommand(n string, v string) Command {
	return setOptionCommand{
		name:  n,
		value: v,
	}
}

type setOptionCommand struct {
	name  string
	value string
}

func (s setOptionCommand) processOutput(e *Engine) error {
	return nil
}

// String implements the Command interface.
func (s setOptionCommand) String() string {
	return fmt.Sprintf("setoption name %s value %s", s.name, s.value)
}

// UCICommand is used to run the "uci" UCI command.
func UCICommand() Command { // nolint:golint
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

// UCINewGameCommand is used to run the "ucinewgame" UCI command.
func UCINewGameCommand() Command { // nolint:golint
	return uciNewGameCommand{}
}

type uciNewGameCommand struct{}

func (u uciNewGameCommand) processOutput(e *Engine) error {
	return nil
}

func (u uciNewGameCommand) String() string {
	return "ucinewgame"
}
