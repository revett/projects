package handler

import (
	"fmt"
	"net/http"
)

// Handler is the exported http.HandlerFunc for Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
