package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/revett/projects/pkg/uci"
)

// Handler is the exported http.HandlerFunc for Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("binary")
	path := fmt.Sprintf("/var/task/engines/%s", s)

	e, err := uci.NewEngine(path, uci.LogOutput)
	if err != nil {
		log.Println(err)
		return
	}

	err = e.Run(
		uci.UCICommand(),
		uci.UCINewGameCommand(),
		uci.IsReadyCommand(),
		uci.PositionCommand(uci.WithFEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12")),
		uci.GoCommand(uci.WithDepth(5)),
	)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
