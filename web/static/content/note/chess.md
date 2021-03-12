---
title: "Chess"
date: 2021-03-12T19:22:47Z
draft: true
---

## Centipawns

Chess engines consider each unit to have a set value:

- Pawn: `1`
- Knight: `3`
- Bishop: `3`
- Rook: `5`
- Queen: `9`

Unit positions are assessed for their relative strength in potential future
exchanges resulting in a score measured in hundredths of a pawn, hence
centipawn.

## Difference Between Blunder, Mistake and Inaccuracy?

Each describe a move which results in a loss of centipawns within a range:

- `?!` Inaccuracy: `0-40`
- `?` Mistake: `40-90`
- `??` Blunder: `90-200`

## UCI

> Uses [Stockfish](https://github.com/official-stockfish/Stockfish) as the
> engine in examples.

Start the engine:

```bash
stockfish
```

Setup engine to use [UCI](http://wbec-ridderkerk.nl/html/UCIProtocol.html):

```bash
uci
ucinewgame
isready
```

Configure the engine to output the top 10 moves instead of just the best move:

```bash
setoption name multipv value 10
```

Links:

- [Full UCI Protocol Documentation by Stefan-Meyer Kahlen (2004)](http://wbec-ridderkerk.nl/html/UCIProtocol.html)
