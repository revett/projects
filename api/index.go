package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bnkamalesh/webgo/v4"
	"github.com/freeeve/uci"
)

// Handler is the exported http.HandlerFunc for Vercel to use.
func Handler(w http.ResponseWriter, req *http.Request) {
	e, err := uci.NewEngine("/var/task/templates/stockfish_13_linux_x64_bmi2")
	if err != nil {
		webgo.R500(w, err)
		return
	}

	engOpts := uci.Options{
		Hash:    1024,
		MultiPV: 1,
		OwnBook: true,
		Threads: 1,
	}
	err = e.SetOptions(engOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = e.SetFEN("r1bqkb1r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 0 1")
	if err != nil {
		log.Fatal(err)
	}

	r, err := e.Go(0, "", 5000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)

  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
