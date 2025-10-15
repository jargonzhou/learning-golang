# golang.org/x
* https://pkg.go.dev/golang.org/x

> These repositories are part of the Go Project but outside the main Go tree. They are developed under looser compatibility requirements than the Go core.

This list is not exhaustive. The full list of Go repositories can be viewed at go.googlesource.com.

| Repo       | Descrition                                                                           |
| :--------- | :----------------------------------------------------------------------------------- |
| benchmarks | benchmarks to measure Go as it is developed.                                         |
| build      | build.golang.org's implementation.                                                   |
| crypto     | additional cryptography packages.                                                    |
| debug      | an experimental debugger for Go.                                                     |
| exp        | experimental and deprecated packages (handle with care; may change without warning). |
| image      | additional imaging packages.                                                         |
| mobile     | experimental support for Go on mobile platforms.                                     |
| mod        | packages for writing tools that work with Go modules.                                |
| net        | additional networking packages.                                                      |
| oauth2     | a client implementation for the OAuth 2.0 spec                                       |
| perf       | packages and tools for performance measurement, storage, and analysis.               |
| pkgsite    | home of the pkg.go.dev website.                                                      |
| review     | a tool for working with Gerrit code reviews.                                         |
| sync       | additional concurrency primitives.                                                   |
| sys        | packages for making system calls.                                                    |
| term       | Go terminal and console support packages.                                            |
| text       | packages for working with text.                                                      |
| time       | additional time packages.                                                            |
| tools      | godoc, goimports, gorename, and other tools.                                         |
| tour       | tour.golang.org's implementation.                                                    |
| vuln       | packages for accessing and analyzing data from the Go Vulnerability Database.        |
| website    | home of the go.dev and golang.org websites.                                          |

# benchmarks
# build
# crypto
# debug
# exp
# image
# mobile
# mod

# net

Go Networking This repository holds supplementary Go networking packages.

| Name                                                                          | Description                                                                                                                                                             |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| bpf                                                                           | Package bpf implements marshaling and unmarshaling of programs for the Berkeley Packet Filter virtual machine, and provides a Go implementation of the virtual machine. |
| context                                                                       | Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.   |
| dict                                                                          | Package dict implements the Dictionary Server Protocol as defined in RFC 2229.                                                                                          |
| dns                                                                           |                                                                                                                                                                         |
| html                                                                          | Package html implements an HTML5-compliant tokenizer and parser.                                                                                                        |
| http                                                                          |                                                                                                                                                                         |
| http2                                                                         | Package http2 implements the HTTP/2 protocol.                                                                                                                           |
| icmp                                                                          | Package icmp provides basic functions for the manipulation of messages used in the Internet Control Message Protocols, ICMPv4 and ICMPv6.                               |
| idna #46, which defines a standard to deal with the transition from IDNA2003. |
| ipv4                                                                          | Package ipv4 implements IP-level socket options for the Internet Protocol version 4.                                                                                    |
| ipv6                                                                          | Package ipv6 implements IP-level socket options for the Internet Protocol version 6.                                                                                    |
| nettest                                                                       | Package nettest provides utilities for network testing.                                                                                                                 |
| netutil                                                                       | Package netutil provides network utility functions, complementing the more common ones in the net package.                                                              |
| proxy                                                                         | Package proxy provides support for a variety of protocols to proxy network data.                                                                                        |
| publicsuffix                                                                  | Package publicsuffix provides a public suffix list based on data from https://publicsuffix.org/                                                                         |
| quic                                                                          | Package quic implements the QUIC protocol.                                                                                                                              |
| route                                                                         | Package route provides basic functions for the manipulation of packet routing facilities on BSD variants.                                                               |
| trace                                                                         | Package trace implements tracing of requests and long-lived objects.                                                                                                    |
| webdav                                                                        | Package webdav provides a WebDAV server implementation.                                                                                                                 |
| websocket                                                                     | Package websocket implements a client and server for the WebSocket protocol as specified in RFC 6455.                                                                   |
| xsrftoken                                                                     | Package xsrftoken provides methods for generating and validating secure XSRF tokens.                                                                                    |

# oauth2
# perf
# pkgsite
* https://pkg.go.dev/golang.org/x/pkgsite
* https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite

This repository hosts the source code of the pkg.go.dev website, and `pkgsite`, a documentation server program.

# review
# sync
* https://pkg.go.dev/golang.org/x/sync

- `errgroup`: Package errgroup provides synchronization, error propagation, and Context cancelation for groups of goroutines working on subtasks of a common task.
- `semaphore`: Package semaphore provides a weighted semaphore implementation.
- `singleflight`: Package singleflight provides a duplicate function call suppression mechanism.
- `syncmap`: Package syncmap provides a concurrent map implementation.

# sys
# term
# text
# time
* https://pkg.go.dev/golang.org/x/time

- `rate`: Package rate provides a rate limiter.

# time/rate

types
- `type Limit float64`: Limit defines the maximum frequency of some events. Limit is represented as number of events per second. A zero Limit allows no events.
- `Limiter struct`: A Limiter controls how frequently events are allowed to happen. It implements a "token bucket" of size b, initially full and refilled at rate r tokens per second.
- `Reservation struct`: A Reservation holds information about events that are permitted by a Limiter to happen after a delay. A Reservation may be canceled, which may enable the Limiter to permit additional events.
- `Sometimes struct`: Sometimes will perform an action occasionally. The `First`, Ev`ery, and In`terval fields govern the behavior of Do, which performs the action. A zero Sometimes value will perform an action exactly once.

# tools
* https://github.com/golang/tools

This repository provides the `golang.org/x/tools` module, comprising various tools and packages mostly for static analysis of Go programs, some of which are listed below. Use the "Go reference" link above for more information about any package.

It also contains the `golang.org/x/tools/gopls` module, whose root package is a language-server protocol (LSP) server for Go. An LSP server analyses the source code of a project and responds to requests from a wide range of editors such as VSCode and Vim, allowing them to support IDE-like functionality.

Selected commands:

- `cmd/goimports` formats a Go program like go fmt and additionally inserts import statements for any packages required by the file after it is edited.
- `cmd/callgraph` prints the call graph of a Go program.
- `cmd/digraph` is a utility for manipulating directed graphs in textual notation.
- `cmd/stringer` generates declarations (including a String method) for "enum" types.
- `cmd/toolstash` is a utility to simplify working with multiple versions of the Go toolchain.

Selected packages:

- `go/ssa` provides a static single-assignment form (SSA) intermediate representation (IR) for Go programs, similar to a typical compiler, for use by analysis tools.
- `go/packages` provides a simple interface for loading, parsing, and type checking a complete Go program from source code.
- `go/analysis` provides a framework for modular static analysis of Go programs.
- `go/callgraph` provides call graphs of Go programs using a variety of algorithms with different trade-offs.
- `go/ast/inspector` provides an optimized means of traversing a Go parse tree for use in analysis tools.
- `go/cfg` provides a simple control-flow graph (CFG) for a Go function.
- `go/gcexportdata` and `go/gccgoexportdata` read and write the binary files containing type information used by the standard and gccgo compilers.
- `go/types/objectpath` provides a stable naming scheme for named entities ("objects") in the go/types API.

Numerous other packages provide more esoteric functionality.

## benchmark 
                                                         
## blog      

Package blog implements a web server for articles written in present format.

## cmd



| Command/Module   | Description                                                                                                                                                                       |
| :--------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| auth             | module                                                                                                                                                                            |
| benchcmp         | Deprecated: benchcmp is deprecated in favor of benchstat: golang.org/x/perf/cmd/benchstat                                                                                         |
| bisect           | Bisect finds changes responsible for causing a failure.                                                                                                                           |
| bundle           | Bundle creates a single-source-file version of a source package suitable for inclusion in a particular target package.                                                            |
| callgraph        | callgraph: a tool for reporting the call graph of a Go program.                                                                                                                   |
| compilebench     | Compilebench benchmarks the speed of the Go compiler.                                                                                                                             |
| cover            | module                                                                                                                                                                            |
| deadcode         | The deadcode command reports unreachable functions in Go programs.                                                                                                                |
| digraph          |                                                                                                                                                                                   |
| eg               | The eg command performs example-based refactoring.                                                                                                                                |
| file2fuzz        | file2fuzz converts binary files, such as those used by go-fuzz, to the Go fuzzing corpus format.                                                                                  |
| fiximports       | The fiximports command fixes import declarations to use the canonical import path for packages that have an "import comment" as defined by https://golang.org/s/go14customimport. |
| getgo            | module                                                                                                                                                                            |
| go-contrib-init  | The go-contrib-init command helps new Go contributors get their development environment set up for the Go contribution process.                                                   |
| godex            | The godex command prints (dumps) exported information of packages or selected package objects.                                                                                    |
| godoc            | module                                                                                                                                                                            |
| goimports        | Command goimports updates your Go import lines, adding missing ones and removing unreferenced ones.                                                                               |
| gomvpkg          | The gomvpkg command moves go packages, updating import declarations.                                                                                                              |
| gonew            | Gonew starts a new Go module by copying a template module.                                                                                                                        |
| gorename         | module                                                                                                                                                                            |
| gotype           | The gotype command, like the front-end of a Go compiler, parses and type-checks a single Go package.                                                                              |
| goyacc           | Goyacc is a version of yacc for Go.                                                                                                                                               |
| guru             | module                                                                                                                                                                            |
| html2article     | This program takes an HTML file and outputs a corresponding article file in present format.                                                                                       |
| present          | Present displays slide presentations and articles.                                                                                                                                |
| present2md       | Present2md converts legacy-syntax present files to Markdown-syntax present files.                                                                                                 |
| signature-fuzzer |                                                                                                                                                                                   |
| splitdwarf       | Splitdwarf uncompresses and copies the DWARF segment of a Mach-O executable into the "dSYM" file expected by lldb and ports of gdb on OSX.                                        |
| ssadump          | ssadump: a tool for displaying and interpreting the SSA form of Go programs.                                                                                                      |
| stress           | The stress utility is intended for catching sporadic failures.                                                                                                                    |
| stringer         | Stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer interface.                                                                                   |
| toolstash        | Toolstash provides a way to save, run, and restore a known good copy of the Go toolchain and to compare the object files generated by two toolchains.                             |


## container
                                                          
## copyright 

Package copyright checks that files have the correct copyright notices.

## cover     

Package cover provides support for parsing coverage profiles generated by "go test -coverprofile=cover.out".

## go

## godoc

## gopls
* https://github.com/golang/tools/tree/master/gopls

The language server for Go. It provides a wide variety of IDE features to any LSP-compatible editor.

- Setting up your workspace: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md

`gopls` supports both *Go module* and *GOPATH modes*.
```shell
✗ go work init
✗ go work use ./codes/src/tgpl 
```

## imports   

Package imports implements a Go pretty-printer (like package "go/format") that also adds or removes import statements as necessary.

## playground

Package playground registers an HTTP handler at "/compile" that proxies requests to the golang.org playground service.

## present   

Package present implements parsing and rendering of present files, which can be slide presentations as in golang.org/x/tools/cmd/present or articles as in golang.org/x/blog (the Go blog).

## refactor

## txtar     

Package txtar implements a trivial text-based file archive format.

# tour

# vuln: Go Vulnerability Management
* https://pkg.go.dev/golang.org/x/vulns

```shell
go install golang.org/x/vuln/cmd/govulncheck@latest

govulncheck ./...
```

# website
