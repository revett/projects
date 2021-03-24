/*
uci is a package for interacting with chess engines that support the
Universal Chess Interface (UCI) protocol. Full documentation of the protocol can
be found: http://wbec-ridderkerk.nl/html/UCIProtocol.html

Setup

The package will require a chess engine to be installed locally, which supports
UCI commands. The best and easiest option is to use Stockfish. Install it via
Homebrew:

	brew install stockfish

Download page: https://stockfishchess.org/download

Basic Example

The following example will start the engine process and search for the best move
looking 10 plies ahead:

	package main

	import (
		"log"
		"os/exec"

		"github.com/revett/projects/pkg/uci"
	)

	func main() {
		e, err := uci.NewEngine(
			exec.Command,
			"/usr/local/bin/stockfish",
		)
		if err != nil {
			log.Fatal(err)
		}

		err = e.Run(
			uci.GoCommand(uci.Depth(10)),
		)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(e.Results.BestMove)

		err = e.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

Multiple Commands

Any number of commands can be passed to the engine to be run in a given order:

	err = e.Run(
		uci.UCICommand(),
		uci.UCINewGameCommand(),
		uci.IsReadyCommand(),
		uci.SetOptionCommand("thread", "2"),
		uci.PositionCommand(uci.FEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12")),
		uci.GoCommand(uci.Depth(10)),
	)
	if err != nil {
		log.Fatal(err)
	}
*/
package uci
