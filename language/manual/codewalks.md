# Codewalks

# Go Doc Comments
* https://go.dev/doc/comment

contents
- packages
- commands
- types
- funcs
- consts
- vars
- syntax: 
  - paragraphs: notes `// MARKER(uid)`, deprecations `// Deprecated: `
  - headings: `// # ...`
  - links: `// ...[RFC 7159]`, `// [RFC 7159]: https://tools.ietf.org/html/rfc7159`
    - doc links: `[Name]`, `[Name1.Name2]`
  - lists: `*, +, -, •`
  - code blocks

See Also
* [Godoc: documenting Go code](https://go.dev/blog/godoc): Note, June 2022: For updated guidelines about documenting Go code, see “Go Doc Comments.”

# Data Race Detector
* https://go.dev/doc/articles/race_detector

Go includes a built-in data race detector. To use it, add the `-race` flag to the go command:

```shell
$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
```

The `GORACE` environment variable sets race detector options.
- `log_path` (default `stderr`): The race detector writes its report to a file named `log_path.pid`. The special names `stdout` and `stderr` cause reports to be written to standard output and standard error, respectively.
- `exitcode` (default 66): The exit status to use when exiting after a detected race.
- `strip_path_prefix` (default ""): Strip this prefix from all reported file paths, to make reports more concise.
- `history_size` (default 1): The per-goroutine memory access history is 32K * 2**history_size elements. Increasing this value can avoid a "failed to restore the stack" error in reports, at the cost of increased memory usage.
- `halt_on_error` (default 0): Controls whether the program exits after reporting first data race.
- `atexit_sleep_ms` (default 1000): Amount of milliseconds to sleep in the main goroutine before exiting.

```shell
# format
GORACE="option1=val1 option2=val2"

# example
$ GORACE="log_path=/tmp/race/report strip_path_prefix=/my/go/sources/" go test -race
```

requirements
- Windows: use [MSYS2] MINGW64 environment to install [mingw-w64-gcc](https://packages.msys2.org/base/mingw-w64-gcc), add `$MSYS2_HOME/mingw64/bin` to `PATH`. see also: [go Issue 61058](https://github.com/golang/go/issues/61058).

runtime overhead: The cost of race detection varies by program, but for a typical program, memory usage may increase by 5-10x and execution time by 2-20x.