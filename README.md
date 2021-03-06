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

	"github.com/revett/projects/pkg/uci"
)

func main() {
	e, err := uci.NewEngine(uci.DefaultCommand(), "/usr/local/bin/stockfish")
	if err != nil {
		log.Fatal(err)
	}

	_, err = e.IsReady()
	if err != nil {
		log.Fatal(err)
	}

	err = e.Stop()
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
