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

Package context defines the Context type, which carries **deadlines**, **cancellation signals**, and other **request-scoped values** *across API boundaries and between processes*.

types
- `CancelCauseFunc func(cause error)`
- `CancelFunc func()`
- `Context interface`
  - `Deadline() (deadline time.Time, ok bool)`
  - `Done() <-chan struct{}`
  - `Err() error`: `nil`, `Canceled`, `DeadlineExceeded`
  - `Value(key any) any`

functions
- `Background()`
- `TODO()`
- `WithValue()`
- `WithCancel()`, `WithCancelCause()`
- `WithoutCancel()`
- `WithDeadline()`
- `WithTimeout()`
- `Cause()`

See Also:
* [learning_go/datastructure/context_test.go](../../codes/learning_go/datastructure/context_test.go)
* `http.Request.Context()`
* `http.Request.WithContext()`
* `http.NewRequestWithContext()`

# crypto
# database
# debug
* https://pkg.go.dev/debug

- `buildinfo`: Package buildinfo provides access to information embedded in a Go binary about how it was built.
- `dwarf`: Package dwarf provides access to DWARF debugging information loaded from executable files, as defined in the DWARF 2.0 Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf.
- `elf`: Package elf implements access to ELF object files.
- `gosym`: Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers.
- `macho`: Package macho implements access to Mach-O object files.
- `pe`: Package pe implements access to PE (Microsoft Windows Portable Executable) files.
- `plan9obj`: Package plan9obj implements access to Plan 9 a.out object files.

# embed
* https://pkg.go.dev/embed

Package embed provides access to files embedded in the running Go program.

directive
```go
//go:embed
```

# encoding
* https://pkg.go.dev/encoding

- `ascii85`: Package ascii85 implements the ascii85 data encoding as used in the btoa tool and Adobe's PostScript and PDF document formats.
- `asn1`: Package asn1 implements parsing of DER-encoded ASN.1 data structures, as defined in ITU-T Rec X.690.
- `base32`: Package base32 implements base32 encoding as specified by RFC 4648.
- `base64`: Package base64 implements base64 encoding as specified by RFC 4648.
- `binary`: Package binary implements simple translation between numbers and byte sequences and encoding and decoding of varints.
- `csv`: Package csv reads and writes comma-separated values (CSV) files.
- `gob`: Package gob manages streams of gobs - binary values exchanged between an Encoder (transmitter) and a Decoder (receiver).
- `hex`: Package hex implements hexadecimal encoding and decoding.
- `json`: Package json implements encoding and decoding of JSON as defined in RFC 7159.
- `pem`: Package pem implements the PEM data encoding, which originated in Privacy Enhanced Mail.
- `xml`: Package xml implements a simple XML 1.0 parser that understands XML name spaces.

## encoding/ascii85
## encoding/asn1
## encoding/base32
## encoding/base64
## encoding/binary
## encoding/csv
## encoding/gob

Package gob manages streams of gobs - binary values exchanged between an Encoder (transmitter) and a Decoder (receiver). A typical use is transporting arguments and results of remote procedure calls (RPCs) such as those provided by [net/rpc](#netrpc).

## encoding/hex

## encoding/json
* https://pkg.go.dev/encoding/json

Package json implements encoding and decoding of JSON as defined in RFC 7159. The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

- `jsontext`: Package jsontext implements syntactic processing of JSON as specified in RFC 4627, RFC 7159, RFC 7493, RFC 8259, and RFC 8785.
- `v2`: Package json implements semantic processing of JSON as specified in RFC 8259.


types
- `Encoder struct`
- `Decoder struct`
- `Marshaler interface`
- `Unmarshaler interface`
- ...

functions
- `NewEncoder()`
- `NewDecoder()`
- `Marshal()`
- `Unmarshal()`
- ...

See Alos
- [encoding_json_test.go](../../codes/snippets/encoding_json_test.go)
- [JSON and Go](https://go.dev/blog/json) - 2011-01-25
- [tags in 'Struct types'](https://go.dev/ref/spec#Struct_types)

## encoding/pem
## encoding/xml

# errors
* https://pkg.go.dev/errors
* [`panic`](https://pkg.go.dev/builtin#panic)
* [`recover`](https://pkg.go.dev/builtin#recover)

# expvar
# flag
# fmt
* https://pkg.go.dev/fmt

Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.


# hash
# html
# image
# index
# io
* https://pkg.go.dev/io

Package io provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package `os`, into shared public interfaces that abstract the functionality, plus some other related primitives.

Because these interfaces and primitives wrap lower-level operations with various implementations, unless otherwise informed clients should **not** assume they are safe for parallel execution.

types
- `Reader interface`
- `Writer interface`
- `Closer interface`
- `Seeker interface`
- ...

functions
- `Copy()`
- `LimitReader()`
- `MultiReader()`
- `TeeReader()`
- `MultiWriter()`
- `ReadAll()`
- ...

See Also
- [os](#os)
- [strings](#strings)

# iter

# log
* https://pkg.go.dev/log

Package log implements a simple logging package

- `slog`: Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs.
- `syslog`: Package syslog provides a simple interface to the system log service.

types
- `Logger`

## log/slog
* https://pkg.go.dev/log/slog

types
- `Handler interface`
- `JSONHandler struct`
- `TextHandler struct`
- `HandlerOptions struct`
- `Logger struct`: A Logger records structured information about each call to its Log, Debug, Info, Warn, and Error methods. For each call, it creates a `Record` and passes it to a `Handler`.
- `Record struct`: A Record holds information about a log event. Copies of a Record share state. Do not modify a Record after handing out a copy to it.
- `Source struct`: Source describes the location of a line of source code.
- `Value struct`: A Value can represent any Go value, but unlike type `any`, it can represent most small values without an allocation. The zero Value corresponds to `nil`.
- `Kind int`: Kind is the kind of a `Value`.
- `Level int`: A Level is the importance or severity of a log event. The higher the level, the more important or severe the event.
- `Attr struct`: An Attr is a key-value pair.
- ...

functions
- `NewJSONHandler()`
- `New()`
- `NewLogLogger()`
- `Info()`
- `Log()`
- `LogAttrs()`: LogAttrs is a more efficient version of `Log` that accepts only Attrs.
- ...

## log.syslog

# maps

# math
* https://pkg.go.dev/math

Package math provides basic constants and mathematical functions.
This package does not guarantee bit-identical results across architectures.

- `big`: Package big implements arbitrary-precision arithmetic (big numbers).
- `bits`: Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.
- `cmplx`: Package cmplx provides basic constants and mathematical functions for complex numbers.
- `rand`: Package rand implements pseudo-random number generators suitable for tasks such as simulation, but it should not be used for security-sensitive work.

## math/big
## math/bits
## math/cmplx
## math/rand

# mine
# net
* https://pkg.go.dev/net

Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

- `http`: Package http provides HTTP client and server implementations.
- `mail`: Package mail implements parsing of mail messages.
- `netip`: Package netip defines an IP address type that's a small value type.
- `rpc`: Package rpc provides access to the exported methods of an object across a network or other I/O connection.
- `smtp`: Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321.
- `textproto`: Package textproto implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP.
- `url`: Package url parses URLs and implements query escaping.

See Also: 
- [golang.org/x/net](../tools/x/golang.org_x.md#net)

## net/http
* https://pkg.go.dev/net/http

Package http provides HTTP client and server implementations.
Starting with Go 1.6, the http package has transparent support for the HTTP/2 protocol when using HTTPS

- `cgi`: Package cgi implements CGI (Common Gateway Interface) as specified in RFC 3875.
- `cookiejar`: Package cookiejar implements an in-memory RFC 6265-compliant http.CookieJar.
- `fcgi`: Package fcgi implements the FastCGI protocol.
- `httptest`: Package httptest provides utilities for HTTP testing.
- `httptrace`: Package httptrace provides mechanisms to trace the events within HTTP client requests.
- `httputil`: Package httputil provides HTTP utility functions, complementing the more common ones in the net/http package.
- `pprof`: Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

types
- `Client struct`
- `Request struct`
- `Response struct`
- `Server struct`
- `Handler interface`
- `HandlerFunc func(ResponseWriter, *Request)`
- `ResponseWriter interface`
- `ServeMux struct`: ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
- `ResponseController struct`: A ResponseController is used by an HTTP handler to control the response.
- ...

functions
- `NewRequestWithContext()`
- `HandleFunc()`
- `NewServeMux()`
- `StripPrefix()`
- `NewResponseController()`
- ...

See Also:
* [net_http_test.go](../../codes/snippets/net_http_test.go)
* [github_test.go](../../codes/snippets/github_test.go)

### net/http/cgi
### net/http/cookiejar
### net/http/fcgi

### net/http/httptest
* https://pkg.go.dev/net/http/httptest

Package httptest provides utilities for HTTP testing.

types
- `Server struct`
- `ResponseRecorder struct`

functions
- `NewServer()`
- `NewRecorder()`

### net/http/httptrace
### net/http/httputil
### net/http/pprof

## net/mail
## net/netip
## net/rpc
## net/smtp
## net/textproto
## net/url

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

Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types. The typical use is to take a value with static type `interface{}` and extract its dynamic type information by calling `TypeOf`, which returns a `Type`.

types
- `ChanDir int`: ChanDir represents a channel type's direction.
- `Kind uint`: A Kind represents the specific kind of type that a `Type` represents. The zero Kind is not a valid kind.
- `MapIter struct`: A MapIter is an iterator for ranging over a map. See `Value.MapRange`.
- `Method struct`: Method represents a single method.
- `SelectCase struct`: A SelectCase describes a single case in a select operation. The kind of case depends on Dir, the communication direction.
- `SelectDir int`: A SelectDir describes the communication direction of a select case.
- `SliceHeader struct`: Deprecated: Use `unsafe.Slice` or `unsafe.SliceData` instead.
- `StringHeader struct`: Deprecated: Use `unsafe.String` or `unsafe.StringData` instead.
- `StructField struct`: A StructField describes a single field in a struct.
- `StructTag string`: A StructTag is the tag string in a struct field.
- `Type interface`: Type is the representation of a Go type.
- `Value struct`: Value is the reflection interface to a Go value.
- `ValueError struct`: A ValueError occurs when a Value method is invoked on a `Value` that does not support it. Such cases are documented in the description of each method.
- ...

```go
// Kind
const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)
```

functions
- `TypeOf()`
- `ValueOf()`
- `New()`
- `MakeChain()`, `MakeMap()`, `MakeSlice()`
- `Append()`
- `MakeFunc()`
- `StructOf()`
- ...

See Also
* [The Laws of Reflection](https://go.dev/blog/laws-of-reflection) - 2011-08-06

# regexp
* https://pkg.go.dev/regexp

# runtime
* https://pkg.go.dev/runtime

Package `runtime` contains operations that interact with Go's runtime system, such as functions to control goroutines. It also includes the low-level type information used by the `reflect` package; see reflect's documentation for the programmable interface to the run-time type system.

Environment Variables
- `GOGC`: sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage.
- `GOMEMLIMIT`: sets a soft memory limit for the runtime.
- `GODEBUG`: controls debugging variables within the runtime.
- `GOMAXPROCS`: limits the number of operating system threads that can execute user-level Go code simultaneously.
- `GORACE`: configures the race detector, for programs built using `-race`. 
- `GOTRACEBACK`: controls the amount of output generated when a Go program fails due to an unrecovered panic or an unexpected runtime condition.

## runtime/cgo
* https://pkg.go.dev/runtime/cgo

Package cgo contains runtime support for code generated by the cgo tool. See the documentation for the cgo command for details on using cgo.

## runtime/coverage
## runtime/debug
## runtime/metrics

## runtime/pprof
* https://pkg.go.dev/runtime/pprof

Package pprof writes runtime profiling data in the format expected by the pprof visualization tool. 

See Also
* [Profiling Go programs with pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/) - 2017

## runtime/race
## runtime/trace


# slices
# sort
# strconv
# strings
* https://pkg.go.dev/strings

Package strings implements simple functions to manipulate UTF-8 encoded strings.

# structs
# sync
* https://pkg.go.dev/sync

- `Locker interface`: A Locker represents an object that can be locked and unlocked.
- `Mutex struct`: A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
  - `RWMutex`: A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer. The zero value for a RWMutex is an unlocked mutex.
- `Cond struct`: Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event. Each Cond has an associated `Locker` L (often a `*Mutex` or `*RWMutex`), which must be held when changing the condition and when calling the `Cond.Wait` method.
- `Map struct`: Map is like a Go `map[any]any` but is safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.
- `Once struct`: Once is an object that will perform exactly one action.
  - `OnceFunc()`, `OnceValue()`, `OnceValues()`
- `WaitGroup struct`: A WaitGroup is a counting semaphore typically used to wait for a group of goroutines or tasks to finish.
- `Pool struct`: A Pool is a set of temporary objects that may be individually saved and retrieved.

## sync/atomic
* https://pkg.go.dev/sync/atomic

Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

types
- `Bool struct`
- `Int32 struct`
- `Int64 struct`
- `Pointer struct`
- `Uint32 struct`
- `Uint64 struct`
- `Uintptr struct`
- `Value struct`

functions: `T` for the types
- `AddT()`
- `CompareAndSwapT()`
- `LoadT()`
- `OrT()`
- `StoreT()`
- `SwapT()`

# syscall
* https://pkg.go.dev/syscall

Package syscall contains an interface to the low-level operating system primitives. The details vary depending on the underlying system, and by default, godoc will display the syscall documentation for the current system.

# testing
* https://pkg.go.dev/testing

Package testing provides support for automated testing of Go packages.

- `fstest`: Package fstest implements support for testing implementations and users of file systems.
- `iotest`: Package iotest implements Readers and Writers useful mainly for testing.
- `quick`: Package quick implements utility functions to help with black box testing.
- `slogtest`: Package slogtest implements support for testing implementations of log/slog.Handler.
- `synctest`: Package synctest provides support for testing concurrent code.

types
- `B struct`: B is a type passed to Benchmark functions to manage benchmark timing and control the number of iterations.
- `F struct`: F is a type passed to fuzz tests.
- `M struct`: M is a type passed to a `TestMain` function to run the actual tests.
- `PB struct`: A PB is used by RunParallel for running parallel benchmarks.
- `T struct`: T is a type passed to Test functions to manage test state and support formatted test logs.
- `TB interface`: TB is the interface common to T, B, and F.
- ...

```go
func TestXxx(*testing.T)

// Benchmarks 
func BenchmarkXxx(*testing.B)
// for b.Loop()
// for range b.N

// Fuzzing 
func FuzzXxx(*testing.F)

// Main
func TestMain(m *testing.M)
```

See Also
* [learning_go/tests/doc.go](../../codes/learning_go/tests/doc.go)
* [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html) - 2007-01-02
* [Don't use build tags for integration tests](https://peter.bourgon.org/blog/2021/04/02/dont-use-build-tags-for-integration-tests.html) - 2021-04-02

## testing/fstest
## testing/iotest
## testing/quick
## testing/slogtest
## testing/synctest

# text
## text/template
* https://pkg.go.dev/text/template
* [text_template_test.go](../codes/snippets/text_template_test.go)

# time
* https://pkg.go.dev/time

Package time provides functionality for measuring and displaying time. The calendrical calculations always assume a Gregorian calendar, with no leap seconds.

types
- `Duration int64`
- `Time struct`

functions
- `ParseDuration()`
- `Truncate()`
- `Round()`
- `Parse()`
- `Format()`
- `Now()`
- `After()`
- `AfterFunc()`
- `NewTicker()`
- ...

# unicode
# unique

# unsafe
* https://pkg.go.dev/unsafe

Package unsafe contains operations that step around the type safety of Go programs.
Packages that import unsafe may be non-portable and are not protected by the Go 1 compatibility guidelines.

types
- `ArbitraryType int`: ArbitraryType is here for the purposes of documentation only and is not actually part of the unsafe package. It represents the type of an arbitrary Go expression.
- `IntegerType int`: IntegerType is here for the purposes of documentation only and is not actually part of the unsafe package. It represents any arbitrary integer type.
- `Pointer *ArbitraryType`: Pointer represents a pointer to an arbitrary type.

functions
- `Alignof(x ArbitraryType) uintptr`: Alignof takes an expression x of any type and returns the required alignment of a hypothetical variable v as if v was declared via `var v = x`.
- `Offsetof(x ArbitraryType) uintptr`: Offsetof returns the offset within the struct of the field represented by x, which must be of the form `structValue.field`.
- `Sizeof(x ArbitraryType) uintptr`: Sizeof takes an expression x of any type and returns the size in bytes of a hypothetical variable v as if v was declared via `var v = x`.
- `String(ptr *byte, len IntegerType) string`: String returns a string value whose underlying bytes start at `ptr` and whose length is `len`.
- `StringData(str string) *byte`: StringData returns a pointer to the underlying bytes of `str`. For an empty string the return value is unspecified, and may be `nil`.
- `Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType`: The function Slice returns a slice whose underlying array starts at `ptr` and whose length and capacity are `len`.
- `SliceData(slice []ArbitraryType) *ArbitraryType`: SliceData returns a pointer to the underlying array of the argument slice.
- ...

See Also
* [Breaking Type Safety in Go: An Empirical Study on the Usage of the unsafe Package](https://arxiv.org/abs/2006.09973) - 2021-07-22
* [unsafe.Pointer and system calls](https://blog.gopheracademy.com/advent-2017/unsafe-pointer-and-system-calls/) - 2017-12-05
* [encoding/binary](#encodingbinary)
* [math/bits](#mathbits)
* [CGO Performance In Go 1.21](https://shane.ai/posts/cgo-performance-in-go1.21/) - 2023-09-01
* [The cost and complexity of Cgo](https://www.cockroachlabs.com/blog/the-cost-and-complexity-of-cgo/) - 2015-12-10
