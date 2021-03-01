package handler

import (
	"net/http"

	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

// Handler is the exported http.HandlerFunc for Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	stockfishPath := "/var/task/templates/stockfish_13_linux_x64_avx2"
	h := handlers.Calculate(stockfishPath)
	h(w, r)
}
