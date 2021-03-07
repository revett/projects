package uci

import "fmt"

// CommandTimeoutError is a custom error for when a timeout occurs when waiting
// for a response from the engine, after a command has been sent.
type CommandTimeoutError struct {
	duration int
	response string
}

// Error to satisfy the interface.
func (c CommandTimeoutError) Error() string {
	return fmt.Sprintf(
		"timed out after %d seconds, waiting for '%s' response from engine",
		c.duration,
		c.response,
	)
}
