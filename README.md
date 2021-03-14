# Projects

🧪 Personal projects and experiments from [@revcd](https://twitter.com/revcd).

## Layout

This repo follows the structure outlined in
[golang-standards/project-layout](https://github.com/golang-standards/project-layout).
Apart from `api/` which is a
[requirement from Vercel](https://vercel.com/docs/serverless-functions/introduction#deploying-serverless-functions)
for deploying serverless functions 😢, however these are very lightweight
handlers which wrap applications in `cmd/`.

### `pkg/uci`

`uci` is a package for interacting with chess engines that support the
[Universal Chess Interface](http://wbec-ridderkerk.nl/html/UCIProtocol.html)
protocol, such as [Stockfish](https://github.com/official-stockfish/Stockfish).

[Full Documentation →](https://github.com/revett/projects/tree/main/pkg/uci)

```go
package main

import (
	"log"

	"github.com/revett/projects/pkg/uci"
)

func main() {
	e, err := uci.NewEngine("/usr/local/bin/stockfish", uci.LogOutput)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()

	err = e.Run(
		uci.UCICommand(),
		uci.UCINewGameCommand(),
		uci.IsReadyCommand(),
		uci.PositionCommand(uci.WithFEN("r3kb1r/pp1q1ppp/4p3/8/3P4/8/P1P2PPP/R1BQ1RK1 b kq - 1 12")),
		uci.GoCommand(uci.WithDepth(18)),
	)
	if err != nil {
		log.Fatal(err)
	}
}
```
