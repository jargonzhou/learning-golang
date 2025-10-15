# Golang Tools
* [_containerd.ipynb](./_containerd.ipynb)
* [_docker.ipynb](./_docker.ipynb)
* [_github.ipynb](./_github.ipynb)
* [_weixin.ipynb](./_weixin.ipynb)

# Playground
* [Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go)
* [GoNB](https://github.com/janpfeifer/gonb): a Go Notebook Kernel for Jupyter
* [gophernotes](https://github.com/gopherdata/gophernotes): a Go kernel for Jupyter notebooks and nteract.
* [Go Wiki: Editors and IDEs for Go](https://go.dev/wiki/IDEsAndTextEditorPlugins)

# Skafold
* [Scaffold](https://github.com/hay-kot/scaffold): a project generation tool similar to cookiecutter written in Go that leverages the Go template engine to generate projects from a template.

# Environment
* [andrewkroh/gvm](https://github.com/andrewkroh/gvm): Go Version Manager (written in Go for cross-platform usability)
* [g](https://github.com/voidint/g): Golang Version Manager
* [goenv](https://github.com/go-nv/goenv): Go Version Management
* [Managing Go installations](https://go.dev/doc/manage-install): how to install multiple versions of Go on the same machine, as well as how to uninstall Go.
* [moovweb/gvm](https://github.com/moovweb/gvm): Go Version Manager

# golang.org/x
* [golang.org_x.md](./x/golang.org_x.md)

# Module, Package, Project
* [delve](https://github.com/go-delve/delve): a debugger for the Go programming language.
* [gomods/athens](https://github.com/gomods/athens): A Go module datastore and proxy.
* [Go Toolchains](https://go.dev/doc/toolchain): Starting in Go 1.21, the Go distribution consists of a `go` command and a bundled Go **toolchain**, which is the standard library as well as the compiler, assembler, and other tools. The `go` command can use its bundled Go toolchain as well as other versions that it finds in the local `PATH` or downloads as needed.
* [marwan-at-work/mod](https://github.com/marwan-at-work/mod): Command line tool to upgrade/downgrade Semantic Import Versioning in Go Modules.
* [Wire](https://github.com/google/wire): a code generation tool that automates connecting components using dependency injection. - This repository was archived by the owner on Aug 25, 2025. It is now read-only.

See Also
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout): Standard Go Project Layout

# Test
* [go-cmp](https://github.com/google/go-cmp): Package for comparing Go values in tests.
* [GoMock](https://github.com/golang/mock): GoMock is a mocking framework for the Go programming language. - Update, June 2023: This repo and tool are no longer maintained. Please see go.uber.org/mock for a maintained fork instead.
* [hey](https://github.com/rakyll/hey): a tiny program that sends some load to a web application.
* [MailHog](https://github.com/mailhog/MailHog): an email testing tool for developers.
* [testify](./testify.md): A toolkit with common assertions and mocks that plays nicely with the standard library.
* [uber-go/mock](https://github.com/uber-go/mock): This project originates from Google's golang/mock repo.

# Linter, Formatter, Type Checker
* [golangci-lint](https://github.com/golangci/golangci-lint): a fast Go linters runner. It runs linters in parallel, uses caching, supports YAML configuration, integrates with all major IDEs, and includes over a hundred linters.
* [golint](https://github.com/golang/lint): a linter for Go source code. Golint is deprecated and frozen. There's no drop-in replacement for it, but tools such as Staticcheck and `go vet` should be used instead.
* [revive](https://github.com/mgechev/revive): Fast, configurable, extensible, flexible, and beautiful linter for Go. Drop-in replacement of *golint*. Revive provides a framework for development of custom rules, and lets you define a strict preset for enhancing your development & code review processes.
* [Staticcheck](https://github.com/dominikh/go-tools): a state of the art linter for the Go programming language. Using static analysis, it finds bugs and performance issues, offers simplifications, and enforces style rules.

# Database
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql): Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package. 
* [go-sqlite3](https://github.com/mattn/go-sqlite3): sqlite3 driver for go using database/sql.
* [gorm](https://github.com/go-gorm/gorm): The fantastic ORM library for Golang, aims to be developer friendly.
* [jmoiron/sqlx](https://github.com/jmoiron/sqlx): general purpose extensions to golang's database/sql
* [pg](https://github.com/lib/pq): A pure Go postgres driver for Go's database/sql package

# Server
* [alice](https://github.com/justinas/alice): Painless middleware chaining for Go.
* [chi](https://github.com/go-chi/chi): lightweight, idiomatic and composable router for building Go HTTP services.
* [cockroachdb/errors](https://github.com/cockroachdb/errors): Go errors with network portability. This library aims to be used as a drop-in replacement to `github.com/pkg/errors` and Go's standard `errors` package. It also provides network portability of error objects, in ways suitable for distributed systems with mixed-version software compatibility.
* [containerd](https://containerd.io/): An industry-standard container runtime with an emphasis on simplicity, robustness and portability.
* [Corba](./corba/corba.md): a library for creating powerful modern CLI applications.
* [Echo](https://github.com/labstack/echo): High performance, extensible, minimalist Go web framework.
* [errors](https://github.com/pkg/errors): Simple error handling primitives. - This repository was archived by the owner on Dec 1, 2021. It is now read-only.
* [Gin](./gin/gin.md): a web framework written in Go.
* [go-homedir](https://github.com/mitchellh/go-homedir): Go library for detecting and expanding the user's home directory without cgo.
* [go-kit/kit](https://github.com/go-kit/kit): A standard library for microservices.
* [golang.org/x](./x/golang.org_x.md): These repositories are part of the Go Project but outside the main Go tree.
* [Gorilla](./gorilla/gorilla.md): a web toolkit for the Go programming language that provides useful, composable packages for writing HTTP-based applications.
* [gRPC](./grpc/grpc.md): A high performance, open source universal RPC framework.
* [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin): a Go (golang) plugin system over RPC.
* [httprouter](https://github.com/julienschmidt/httprouter): A high performance HTTP request router that scales well
* [moby](https://github.com/moby/moby/tree/master/client): Go client for the Docker Engine API.
* [Sanitizers](https://github.com/google/sanitizers): This project is the home for Sanitizers: AddressSanitizer, MemorySanitizer, ThreadSanitizer, LeakSanitizer, and more. The actual code resides in the LLVM repository. Here we keep extended documentation, bugfixes and some helper code.
* [shopspring/decimal](https://github.com/shopspring/decimal): Arbitrary-precision fixed-point decimal numbers in go.
* [viper](https://github.com/spf13/viper): Go configuration with fangs.

See Also
* [Top Go Web Frameworks](https://github.com/mingrammer/go-web-framework-stars)

## Logging
* [go-kit/log](https://github.com/go-kit/log): A minimal and extensible structured logger. - last release v0.2.1 2022-05-15
* [Logrus](https://github.com/sirupsen/logrus): a structured logger for Go (golang), completely API compatible with the standard library logger.
* [uber-go/zap](https://github.com/uber-go/zap): Blazing fast, structured, leveled logging in Go.

# Misc
* [Go Imagick](https://github.com/gographics/imagick): Go Imagick is a Go bind to ImageMagick's MagickWand C API. ([ImageMagick](https://imagemagick.org/) is a powerful open-source software suite for creating, editing, converting, and manipulating images in over 200 formats.)
* [gomacro](https://github.com/cosmos72/gomacro): interactive Go interpreter and debugger with generics and macros
* [goproxy.cn](https://goproxy.cn/): The most trusted Go module proxy in China.
* [go-yaml/yaml](https://github.com/go-yaml/yaml): YAML support for the Go language. - This repository was archived by the owner on Apr 2, 2025. It is now read-only.
* [goccy/go-yaml](https://github.com/goccy/go-yaml): YAML support for the Go language.
* [minica](https://github.com/jsha/minica): Minica is a simple CA intended for use in situations where the CA operator also operates each host where a certificate will be used. It automatically generates both a key and a certificate when asked to produce a certificate. It does not offer OCSP or CRL services. Minica is appropriate, for instance, for generating certificates for RPC systems or microservices.

# See Also
* [Awesome Go](https://awesome-go.com/)
* [Go Style](https://google.github.io/styleguide/go/index): The Go Style Guide and accompanying documents codify the current best approaches for writing readable and idiomatic Go.
* [pkg.go.dev](https://pkg.go.dev/): a site that automatically indexes open source Go projects.
* [Release History](https://go.dev/doc/devel/release): This page summarizes the changes between official stable releases of Go. The change log has the full details.