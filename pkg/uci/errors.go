package uci

import "fmt"

type CommandTimeoutError struct {
	duration int
	response string
}

func (c CommandTimeoutError) Error() string {
	return fmt.Sprintf(
		"timed out after %d seconds, waiting for '%s' response from engine",
		c.duration,
		c.response,
	)
}
