package main

import (
	"log"

	"github.com/revett/projects/pkg/uci"
)

func main() {
	// e, err := uci.NewEngine("/usr/local/bin/stockfish", uci.LogOutput)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer e.Close()

	// err = e.Run(
	// 	uci.UCICommand(),
	// 	uci.UCINewGameCommand(),
	// 	uci.IsReadyCommand(),
	// 	uci.PositionCommand(),
	// 	uci.PositionCommand(uci.WithFEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12")),
	// 	uci.PositionCommand(uci.WithMoves("e2e4")),
	// 	uci.PositionCommand(uci.WithMoves("e2e4", "e7e5")),
	// 	uci.GoCommand(),
	// 	uci.GoCommand(uci.WithDepth(5)),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e, err := uci.NewEngine("/usr/local/bin/stockfish", uci.LogOutput)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()

	err = e.Run(
		uci.UCICommand(),
		uci.UCINewGameCommand(),
		uci.IsReadyCommand(),
		uci.PositionCommand(uci.WithFEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12")),
		uci.GoCommand(uci.WithDepth(18)),
	)
	if err != nil {
		log.Fatal(err)
	}
}
