---
title: "Chess"
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

## Detecting Cheating

[clarkerubber/irwin](https://github.com/clarkerubber/irwin)

> **Irwin** (named after Steve Irwin, the Crocodile Hunter) is the AI that learns
> cheating patterns, marks cheaters, and assists moderators in assessing
> potential cheaters.

## Difference Between Blunder, Mistake and Inaccuracy?

Each describe a move which results in a loss of centipawns within a range:

- `?!` Inaccuracy: `0-40`
- `?` Mistake: `40-90`
- `??` Blunder: `90-200`

## Engines

- [Houdini](https://www.cruxis.com/chess/houdini.htm)
- [Komodo](https://komodochess.com/)
- [Maia](https://maiachess.com/)
- [Stockfish](https://github.com/official-stockfish/Stockfish)
- [Sunfish](https://github.com/thomasahle/sunfish)

## Game Database

[Online Chess Database (chessgames.com)](https://www.chessgames.com/index.html) -
Over 981,000 games.

## Manipulate PGN

[pgn-extract (cs.kent.ac.uk)](https://www.cs.kent.ac.uk/people/staff/djb/pgn-extract/)

> A command-line program for searching, manipulating and formatting chess games
> recorded in the Portable Game Notation (PGN).

## Moves and Plies

- One `move` consists of a turn by each player
- A `ply` consists of a turn taken by a single player
- E.g. 20 moves would mean 40 plies have been completed (20 by white and 20 by
  black)
- The word `turn` can have different local meanings, so should be avoided if
  possible

Links:

- [Ply (wikipedia.org)](<https://en.wikipedia.org/wiki/Ply_(game_theory)>)

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

Search for the best move(s), 10 plies deep:

```bash
go depth 10
```

Search for the best move(s) for 5 seconds:

```bash
go movetime 5000
```

Rate the move `e2e4`
([long algebraic notation (LAN)](https://en.wikipedia.org/wiki/Chess_notation#Notation_systems_for_humans)),
searching 10 plies deep:

```bash
go depth 10 searchmoves e2e4
```

Links:

- [Full UCI Protocol Documentation by Stefan-Meyer Kahlen (2004)](http://wbec-ridderkerk.nl/html/UCIProtocol.html)
- [How to Use a Chess Engine (decodechess.com)](https://decodechess.com/how-to-use-a-chess-engine-guide/)
