package main

import (
	"github.com/bnkamalesh/webgo/v4"
	"github.com/bnkamalesh/webgo/v4/middleware"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

const stockfishPath = "/usr/local/bin/stockfish"

func routes() []*webgo.Route {
	return []*webgo.Route{
		handlers.CalculateRoute(stockfishPath),
	}
}

func main() {
	c := webgo.Config{
		Port: "1323",
	}

	r := webgo.NewRouter(&c, routes())
	r.Use(middleware.AccessLog)
	r.Start()
}
