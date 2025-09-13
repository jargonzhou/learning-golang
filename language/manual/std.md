# Go Standard Library
* https://pkg.go.dev/std


| Name      | Description                                                                                                                                                                                                           |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| archive   |                                                                                                                                                                                                                       |
| bufio     | Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O. |
| builtin   | Package builtin provides documentation for Go's predeclared identifiers.                                                                                                                                              |
| bytes     | Package bytes implements functions for the manipulation of byte slices.                                                                                                                                               |
| cmp       | Package cmp provides types and functions related to comparing ordered values.                                                                                                                                         |
| compress  |                                                                                                                                                                                                                       |
| container |                                                                                                                                                                                                                       |
| context   | Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.                                                 |
| crypto    | Package crypto collects common cryptographic constants.                                                                                                                                                               |
| database  |                                                                                                                                                                                                                       |
| debug     |                                                                                                                                                                                                                       |
| embed     | Package embed provides access to files embedded in the running Go program.                                                                                                                                            |
| encoding  | Package encoding defines interfaces shared by other packages that convert data to and from byte-level and textual representations.                                                                                    |
| errors    | Package errors implements functions to manipulate errors.                                                                                                                                                             |
| expvar    | Package expvar provides a standardized interface to public variables, such as operation counters in servers.                                                                                                          |
| flag      | Package flag implements command-line flag parsing.                                                                                                                                                                    |
| fmt       | Package fmt implements formatted I/O with functions analogous to C's printf and scanf.                                                                                                                                |
| go        |                                                                                                                                                                                                                       |
| hash      | Package hash provides interfaces for hash functions.                                                                                                                                                                  |
| html      | Package html provides functions for escaping and unescaping HTML text.                                                                                                                                                |
| image     | Package image implements a basic 2-D image library.                                                                                                                                                                   |
| index     |                                                                                                                                                                                                                       |
| io        | Package io provides basic interfaces to I/O primitives.                                                                                                                                                               |
| iter      | Package iter provides basic definitions and operations related to iterators over sequences.                                                                                                                           |
| log       | Package log implements a simple logging package.                                                                                                                                                                      |
| maps      | Package maps defines various functions useful with maps of any type.                                                                                                                                                  |
| math      | Package math provides basic constants and mathematical functions.                                                                                                                                                     |
| mime      | Package mime implements parts of the MIME spec.                                                                                                                                                                       |
| net       | Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.                                                                                    |
| os        | Package os provides a platform-independent interface to operating system functionality.                                                                                                                               |
| path      | Package path implements utility routines for manipulating slash-separated paths.                                                                                                                                      |
| plugin    | Package plugin implements loading and symbol resolution of Go plugins.                                                                                                                                                |
| reflect   | Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.                                                                                                        |
| regexp    | Package regexp implements regular expression search.                                                                                                                                                                  |
| runtime   | Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines.                                                                                                  |
| slices    | Package slices defines various functions useful with slices of any type.                                                                                                                                              |
| sort      | Package sort provides primitives for sorting slices and user-defined collections.                                                                                                                                     |
| strconv   | Package strconv implements conversions to and from string representations of basic data types.                                                                                                                        |
| strings   | Package strings implements simple functions to manipulate UTF-8 encoded strings.                                                                                                                                      |
| structs   | Package structs defines marker types that can be used as struct fields to modify the properties of a struct.                                                                                                          |
| sync      | Package sync provides basic synchronization primitives such as mutual exclusion locks.                                                                                                                                |
| syscall   | Package syscall contains an interface to the low-level operating system primitives.                                                                                                                                   |
| testing   | Package testing provides support for automated testing of Go packages.                                                                                                                                                |
| text      |                                                                                                                                                                                                                       |
| time      | Package time provides functionality for measuring and displaying time.                                                                                                                                                |
| unicode   | Package unicode provides data and functions to test some properties of Unicode code points.                                                                                                                           |
| unique    | The unique package provides facilities for canonicalizing ("interning") comparable values.                                                                                                                            |
| unsafe    | Package unsafe contains operations that step around the type safety of Go programs.                                                                                                                                   |
| weak      | Package weak provides ways to safely reference memory weakly, that is, without preventing its reclamation.                                                                                                            |

# archive
# bufio
# builtin
# bytes

# cmp
* https://pkg.go.dev/cmp
* [cmp_test.go](../codes/snippets/cmp_test.go)
# compress
# container

# context
* https://pkg.go.dev/context
* [context.ipynb](./context.ipynb)

# crypto
# database
# debug
# embed
# encoding
## encoding/json
* [encoding_json_test.go](../codes/snippets/encoding_json_test.go)

# errors
* https://pkg.go.dev/errors
* [`panic`](https://pkg.go.dev/builtin#panic)
* [`recover`](https://pkg.go.dev/builtin#recover)

# expvar
# flag
# fmt
# hash
# html
# image
# index
# io
# iter
# log
* https://pkg.go.dev/log@go1.23.5
## log/slog
structured logging
# maps
# math
# mine
# net
* https://pkg.go.dev/net
* See also: [golang.org/x/net](../tools/x/golang.org_x.md#net)

## net/http
* [net_http_test.go](../codes/snippets/net_http_test.go)
* [github_test.go](../codes/snippets/github_test.go)

# os
* https://pkg.go.dev/os
## os/exec
## os/signal
* https://gobyexample.com/signals
## os/user
# path
# plugin
# reflect
* https://pkg.go.dev/reflect
# regexp
* https://pkg.go.dev/regexp
# runtime
## runtime/cgo
* https://pkg.go.dev/runtime/cgo


## runtime/coverage
## runtime/debug
## runtime/metrics
## runtime/pprof
* https://pkg.go.dev/runtime/pprof
* [Profiling Go programs with pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/) - 2017

Package pprof writes runtime profiling data in the format expected by the pprof visualization tool. 

## runtime/race
## runtime/trace
# slices
# sort
# strconv
# strings
# structs
# sync
* https://pkg.go.dev/sync
# syscall
* https://pkg.go.dev/syscall

# testing
* https://pkg.go.dev/testing

* [tests.ipynb](./tests.ipynb)

# text
## text/template
* https://pkg.go.dev/text/template
* [text_template_test.go](../codes/snippets/text_template_test.go)

# time
* https://pkg.go.dev/time
# unicode
# unique
# unsafe
* https://pkg.go.dev/unsafe