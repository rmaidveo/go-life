# go-life

[![GoDoc](https://godoc.org/github.com/rmaidveo/go-life?status.svg)](https://godoc.org/github.com/rmaidveo/go-life)
[![Go Report Card](https://goreportcard.com/badge/github.com/rmaidveo/go-life)](https://goreportcard.com/report/github.com/rmaidveo/go-life)

[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life).

## Installation

```
$ go install github.com/rmaidveo/go-life@latest
```

## Usage

```
$ go-life -h | -help | --help
$ go-life [options]
```

Stdin: grid in the [plaintext](https://www.conwaylife.com/wiki/Plaintext) format.

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-delay DURATION` &mdash; delay between frames (e.g. `72h3m0.5s`; default: `100ms`).

## Output Example

```
........................O.............
......................O.O.............
............OO......OO............OO..
...........O...O....OO............OO..
OO........O.....O...OO................
OO........O...O.OO....O.O.............
..........O.....O.......O.............
...........O...O......................
............OO........................
.......................O..............
........................OO............
.......................OO.............
......................................
......................................
......................................
......................................
......................................
..............................O.O.....
...............................OO.....
...............................O......
```

## License

The MIT License (MIT)

Copyright &copy; 2024 rmaidveo