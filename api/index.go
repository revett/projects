package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/freeeve/uci"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/calculate", handlers.Calculate)

	e.Logger.Fatal(e.Start(":1323"))
}

// Handler is required by Vercel.
func Handler(w http.ResponseWriter, req *http.Request) {
	e, err := uci.NewEngine("/var/task/templates/stockfish_13_linux_x64_bmi2")
	if err != nil {
		log.Fatal(err)
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
