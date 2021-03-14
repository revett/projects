package uci

const (
	goCmd         = "go depth 10"
	isReadyCmd    = "isready"
	positionCmd   = "position %s"
	uciCmd        = "uci"
	uciNewGameCmd = "ucinewgame"
)

type IsReadyCommand struct{}

func (i IsReadyCommand) execute(e *Engine) error {
	if err := e.sendCommand("isready"); err != nil {
		return err
	}

	_, err := e.readUntil("readyok")

	return err
}
