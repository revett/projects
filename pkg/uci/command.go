package uci

type Command interface {
	execute(*Engine) error
	string() string
}
