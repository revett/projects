package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bnkamalesh/webgo/v4"
	"github.com/revett/projects/pkg/uci"
)

type request struct {
	Depth    int    `json:"depth"`
	FEN      string `json:"fen"`
	MoveTime int64  `json:"moveTime"`
	MultiPV  int    `json:"multiPV"`
}

const maxMoveTime = 7000

// Calculate is a http.HandlerFunc which runs the 'go' UCI command with a set
// of options.
func Calculate(stockfishPath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get(webgo.HeaderContentType)
		if contentType != webgo.JSONContentType {
			msg := "Content-Type header is not application/json"
			webgo.SendError(w, msg, http.StatusUnsupportedMediaType)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var req request
		err := dec.Decode(&req)
		if err != nil {
			webgo.R400(w, err)
			return
		}

		if req.MoveTime > maxMoveTime {
			req.MoveTime = maxMoveTime
		}

		e, err := uci.NewEngine(stockfishPath)
		if err != nil {
			webgo.R500(w, err)
			return
		}

		engOpts := uci.Options{
			Hash:    1024,
			MultiPV: req.MultiPV,
			OwnBook: true,
			Threads: 1,
		}
		err = e.SetOptions(engOpts)
		if err != nil {
			webgo.R500(w, err)
			return
		}

		err = e.SetFEN(req.FEN)
		if err != nil {
			webgo.R500(w, err)
			return
		}

		res, err := e.Go(req.Depth, "", req.MoveTime)
		if err != nil {
			webgo.R500(w, err)
			return
		}

		webgo.R200(w, res)
	}
}

// CalculateRoute is the webgo route definition.
func CalculateRoute(stockfishPath string) *webgo.Route {
	return &webgo.Route{
		Name:    "calculate",
		Method:  http.MethodPost,
		Pattern: "/calculate",
		Handlers: []http.HandlerFunc{
			Calculate(stockfishPath),
		},
	}
}
