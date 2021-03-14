package uci

type command interface {
	execute(*Engine) error
}
