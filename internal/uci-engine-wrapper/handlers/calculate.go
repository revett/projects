package handlers

import (
	"net/http"

	"github.com/freeeve/uci"
	"github.com/labstack/echo/v4"
)

type request struct {
	Depth    int    `json:"depth"`
	FEN      string `json:"fen"`
	MoveTime int64  `json:"moveTime"`
	MultiPV  int    `json:"multiPV"`
}

const maxMoveTime = 7000

// Calculate runs the 'go' UCI command with a set of options.
func Calculate(c echo.Context) error {
	var req request
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.MoveTime > maxMoveTime {
		req.MoveTime = maxMoveTime
	}

	e, err := uci.NewEngine("/usr/local/bin/stockfish")
	if err != nil {
		return err
	}

	engOpts := uci.Options{
		Hash:    1024,
		MultiPV: req.MultiPV,
		OwnBook: true,
		Threads: 1,
	}
	err = e.SetOptions(engOpts)
	if err != nil {
		return err
	}

	err = e.SetFEN(req.FEN)
	if err != nil {
		return err
	}

	r, err := e.Go(req.Depth, "", req.MoveTime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}
