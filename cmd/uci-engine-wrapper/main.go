package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

const (
	port          = ":1323"
	stockfishPath = "/usr/local/bin/stockfish"
)

func main() {
	r := gin.Default()
	r.GET("/search", handlers.Search(stockfishPath))

	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
