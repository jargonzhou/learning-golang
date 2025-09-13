# Codewalks

# Data Race Detector
* https://go.dev/doc/articles/race_detector

Go includes a built-in data race detector. To use it, add the `-race` flag to the go command:

```shell
$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
```