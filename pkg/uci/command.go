package uci

import "fmt"

type Command interface {
	fmt.Stringer
	processOutput(*Engine) error
}
