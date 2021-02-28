package main

import (
	"fmt"
	"log"

	"github.com/freeeve/uci"
)

const enginePath = "/usr/local/bin/stockfish"

func main() {
	e, err := uci.NewEngine(enginePath)
	if err != nil {
		log.Fatal(err)
	}

	e.SetOptions(uci.Options{
		Hash:    128,
		Ponder:  false,
		OwnBook: true,
		MultiPV: 4,
	})

	e.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")

	opts := uci.HighestDepthOnly | uci.IncludeUpperbounds | uci.IncludeLowerbounds
	r, err := e.GoDepth(10, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}
