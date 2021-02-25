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
setoption name <x> value <y>
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

> Makes use of the [UCI Protocol](http://wbec-ridderkerk.nl/html/UCIProtocol.html).
