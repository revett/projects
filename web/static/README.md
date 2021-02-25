# static

Basic static website built using [Tailwind](https://tailwindcss.com) and
[Parcel](https://parceljs.org).

## Stockfish Usage

Install via Brew:

```
brew install stockfish
```

Start:

```
stockfish
```

Check if ready (output should be `readyok`):

```
isready
```

Output current UCI option configuration:

```
uci
```

Set supported UCI option:

```
uci name <x> value <y>
```

Set position:

```
position startpos
position fen <x>
```

Search:

```
go infinite
go depth <n>
go movetime <ms>
```

