package handlers

import (
	"net/http"

	"github.com/freeeve/uci"
	"github.com/labstack/echo/v4"
)

// Calculate runs the 'go' UCI command with a set of options.
func Calculate(c echo.Context) error {
	e, err := uci.NewEngine("/usr/local/bin/stockfish")
	if err != nil {
		return err
	}

	engOpts := uci.Options{
		Hash:    1024,
		MultiPV: 4,
		OwnBook: true,
		Ponder:  false,
		Threads: 1,
	}
	err = e.SetOptions(engOpts)
	if err != nil {
		return err
	}

	err = e.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")
	if err != nil {
		return err
	}

	r, err := e.Go(10, "", 5000)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}
