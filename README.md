# Advent of Code 2024

This repository includes my [Go](https://golang.org/) code for the [Advent of Code 2024](https://adventofcode.com/2024/about).

**Please keep in mind most of this is hacked together.
I don't claim it's perfect code or even in all cases 100% valid/defined/proper code or whatever.
It's for fun and it got the job done in a short amount of time.** 😀

I'll update this repository now and then to include all my solutions whenever I've got time to properly comment and upload it.

This is mostly for those that are curious.
If you want to look at the code, experiment with it, change it, etc. be my guest.

## Running from command line

All code will take a single command line argument which is the path to the input file, which generally will be stored in the data directory under their specific day, i.e `data/day01/input.txt`.
To run the code for any given day, use the following line (replacing `1` with the number of the desired day):

```bash
go run ./cmd/day01 <path to input file>
```

As an alternative, you can also compile the executables explicitly, using `build` rather than `run`:

```bash
go build ./cmd/day01
./day01 <path to input file>
```

## Running/debugging in VS Code

I'm using [Visual Studio Code](https://code.visualstudio.com/) for writing and debugging.
This should work out of the box, if you open the repository directory directly.

You should be able to pick the source code file of any day and simply hit <kbd>F5</kbd> to run/debug it.
A matching `launch.json` configuration file is included.