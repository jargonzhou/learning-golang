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