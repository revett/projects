# Projects

🧪 Personal projects and experiments from [@revcd](https://twitter.com/revcd).

## Layout

This repo follows the structure outlined in
[golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## Active

### `pkg/uci`

Package for interacting with a chess engine that supports the
[Universal Chess Interface](http://wbec-ridderkerk.nl/html/UCIProtocol.html)
(UCI) protocol. It builds upon the excellent
[freeeve/uci](https://github.com/freeeve/uci) repo.

```go
package main

import (
	"log"
	"time"

	"github.com/revett/projects/pkg/uci"
)

func main() {
	e, err := uci.NewEngine(
		uci.DefaultCommand(),
		"/usr/local/bin/stockfish",
		uci.Debug,
		uci.InitialiseGame,
		uci.WithCommandTimeout(100*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()

	err = e.Position("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Pronounceable Gibberish

Making use of syllable structure to generate nonsense which can easily be
understood or spoken.

### Serverless UCI Wrapper

Creating a generic function for interacting with a
[UCI](https://en.wikipedia.org/wiki/Universal_Chess_Interface) compliant chess
engine.
