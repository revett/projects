package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

// Handler is the exported http.HandlerFunc for Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	e := gin.Default()
	e.GET("/api", handlers.Search(os.Getenv("STOCKFISH_PATH")))
	e.ServeHTTP(w, r)
}
