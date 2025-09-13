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


| Name       | Description                                                                                                                                                                                 |
| :--------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| benchmark  |                                                                                                                                                                                             |
| blog       | Package blog implements a web server for articles written in present format.                                                                                                                |
| cmd        |                                                                                                                                                                                             |
| container  |                                                                                                                                                                                             |
| copyright  | Package copyright checks that files have the correct copyright notices.                                                                                                                     |
| cover      | Package cover provides support for parsing coverage profiles generated by "go test -coverprofile=cover.out".                                                                                |
| go         |                                                                                                                                                                                             |
| godoc      |                                                                                                                                                                                             |
| gopls      |                                                                                                                                                                                             |
| imports    | Package imports implements a Go pretty-printer (like package "go/format") that also adds or removes import statements as necessary.                                                         |
| playground | Package playground registers an HTTP handler at "/compile" that proxies requests to the golang.org playground service.                                                                      |
| present    | Package present implements parsing and rendering of present files, which can be slide presentations as in golang.org/x/tools/cmd/present or articles as in golang.org/x/blog (the Go blog). |
| refactor   |                                                                                                                                                                                             |
| txtar      | Package txtar implements a trivial text-based file archive format.                                                                                                                          |


## gopls
* https://github.com/golang/tools/tree/master/gopls

The language server for Go. It provides a wide variety of IDE features to any LSP-compatible editor.

- Setting up your workspace: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md

`gopls` supports both *Go module* and *GOPATH modes*.
```shell
✗ go work init
✗ go work use ./codes/src/tgpl 
```