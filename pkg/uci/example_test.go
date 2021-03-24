package uci_test

import (
	"fmt"

	"github.com/revett/projects/pkg/uci"
)

func ExampleFEN() {
	c := uci.PositionCommand(
		uci.FEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12"),
	)
	fmt.Println(c.String())
	// output:
	// position fen r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12
}

func ExampleMoves() {
	c := uci.PositionCommand(
		uci.Moves("e2e4", "e7e5"),
	)
	fmt.Println(c.String())
	// output:
	// position moves e2e4 e7e5
}
