package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/revett/projects/pkg/uci"
)

// Search returns a http.HandlerFunc which uses the `go` UCI command to find the
// next best move for a given board position.
func Search(stockfishPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		e, err := uci.NewEngine(
			stockfishPath, uci.LogOutput, uci.WithCommandTimeout(3*time.Second),
		)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":  err.Error(),
					"status": http.StatusInternalServerError,
				},
			)

			return
		}

		fen := c.Query("fen")

		err = e.Run(
			uci.UCICommand(),
			uci.UCINewGameCommand(),
			uci.IsReadyCommand(),
			uci.PositionCommand(uci.WithFEN(fen)),
			uci.GoCommand(),
		)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":  err.Error(),
					"status": http.StatusInternalServerError,
				},
			)

			return
		}

		// GoCommand needs to return search results, will remove this hard-coding.
		res := struct {
			BestMove string `json:"bestMove"`
		}{
			BestMove: "e2e4",
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"data":   res,
				"status": http.StatusOK,
			},
		)
	}
}
