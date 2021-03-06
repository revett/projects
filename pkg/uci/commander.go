package uci

import "os/exec"

type commander interface {
	Command(s string, a ...string) *exec.Cmd
}

type defaultCommander struct{}

func (d defaultCommander) Command(s string, a ...string) *exec.Cmd {
	return exec.Command(s, a...)
}

// DefaultCommand creates a commander{} interface which wraps exec.Command().
func DefaultCommand() commander {
	return defaultCommander{}
}
