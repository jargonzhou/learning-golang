# Go Programming Language
# The Go Programming Language Specification
* https://go.dev/ref/spec

# Go User Manual
* https://go.dev/doc

## Getting Started
- **Installing Go**: Instructions for downloading and installing Go.
- **Tutorial**: Getting started: A brief Hello, World tutorial to get started. Learn a bit about Go code, tools, packages, and modules.
- **Tutorial**: Create a module: A tutorial of short topics introducing functions, error handling, arrays, maps, unit testing, and compiling.
- **Tutorial**: Getting started with multi-module workspaces: Introduces the basics of creating and using multi-module workspaces in Go. Multi-module workspaces are useful for making changes across multiple modules.
- **Tutorial**: Developing a RESTful API with Go and Gin: Introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework.
- **Tutorial**: Getting started with generics: With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.
- **Tutorial**: Getting started with fuzzing: Fuzzing can generate inputs to your tests that can catch edge cases and security issues that you may have missed.
- **Writing Web Applications**: Building a simple web application.
- **How to write Go code**: This doc explains how to develop a simple set of Go packages inside a module, and it shows how to use the `go` command to build and test packages.
- **A Tour of Go**: An interactive introduction to Go in four sections. The first section covers basic syntax and data structures; the second discusses methods and interfaces; the third is about Generics; and the fourth introduces Go's concurrency primitives. Each section concludes with a few exercises so you can practice what you've learned.

## Developing modules
- **Developing and publishing modules**: You can collect related packages into modules, then publish the modules for other developers to use. This topic gives an overview of developing and publishing modules.
- **Module release and versioning workflow**: When you develop modules for use by other developers, you can follow a workflow that helps ensure a reliable, consistent experience for developers using the module. This topic describes the high-level steps in that workflow.
- **Managing module source**: When you're developing modules to publish for others to use, you can help ensure that your modules are easier for other developers to use by following the repository conventions described in this topic.
- **Organizing a Go module**: What is the right way to organize the files and directories in a typical Go project? This topic discusses some common layouts depending on the kind of module you have.
- **Developing a major version update**: A major version update can be very disruptive to your module's users because it includes breaking changes and represents a new module. Learn more in this topic.
- **Publishing a module**: When you want to make a module available for other developers, you publish it so that it's visible to Go tools. Once you've published the module, developers importing its packages will be able to resolve a dependency on the module by running commands such as `go get`.
- **Module version numbering**: A module's developer uses each part of a module's version number to signal the version’s stability and backward compatibility. For each new release, a module's release version number specifically reflects the nature of the module's changes since the preceding release.

## Using and understanding Go
- **Effective Go**: A document that gives tips for writing clear, idiomatic Go code. A must read for any new Go programmer. It augments the tour and the language specification, both of which should be read first.
- **Frequently Asked Questions**: Answers to common questions about Go.
- **Editor plugins and IDEs**: A document that summarizes commonly used editor plugins and IDEs with Go support.
- **Diagnostics**: Summarizes tools and methodologies to diagnose problems in Go programs.
- **A Guide to the Go Garbage Collector**: A document that describes how Go manages memory, and how to make the most of it.
- **Managing dependencies**: When your code uses external packages, those packages 
- **Fuzzing**: Main documentation page for Go fuzzing.
- **Coverage for Go applications**: Main documentation page for coverage testing of Go applications.
- **Profile-guided optimization**: Main documentation page for profile-guided optimization 

## Talks
- **A Video Tour of Go**: Three things that make Go fast, fun, and productive: interfaces, reflection, and concurrency. Builds a toy web crawler to demonstrate these.
- **Code that grows with grace**: One of Go's key design goals is code adaptability; that it should be easy to take a simple design and build upon it in a clean and natural way. In this talk Andrew Gerrand describes a simple "chat roulette" server that matches pairs of incoming TCP connections, and then use Go's concurrency mechanisms, interfaces, and standard library to extend it with a web interface and other features. While the function of the program changes dramatically, Go's flexibility preserves the original design as it grows.
- **Go Concurrency Patterns**: Concurrency is the key to designing high performance network services. Go's concurrency primitives 
- **Advanced Go Concurrency Patterns**: This talk expands on the _Go Concurrency Patterns_ talk to dive deeper into Go's concurrency primitives.

## References
* [references.md](./manual/references.md)

- **Package Documentation**: The documentation for the Go standard library.
- **Command Documentation**: The documentation for the Go tools.
- **Language Specification**: The official Go Language specification.
- **Go Modules Reference**: A detailed reference manual for Go's dependency management system.
- **go.mod file reference**: Reference for the directives included in a go.mod file.
- **The Go Memory Model**: A document that specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.
- **Contribution Guide**: Contributing to Go.
- **Release History**: A summary of the changes between Go releases.

## Codewalks
* [codewalks.md](./manual/codewalks.md)

Guided tours of Go programs.
- **First-Class Functions in Go**
- **Generating arbitrary text**: a Markov chain algorithm
- **Share Memory by Communicating**

> Language
- **JSON-RPC**: a tale of interfaces
- **Go's Declaration Syntax**
- **Defer, Panic, and Recover**
- **Go Concurrency Patterns**: Timing out, moving on
- **Go Slices**: usage and internals
- **A GIF decoder**: an exercise in Go interfaces
- **Error Handling and Go**

> Packages
- **JSON and Go**: using the json package.
- **Gobs of data**: the design and use of the gob package..
- **The Laws of Reflection**: the fundamentals of the reflect package.
- **The Go image package**: the fundamentals of the image package..
- **The Go image/draw package**: the fundamentals of the image/draw package.

> Modules
- **Using Go Modules**: an introduction to using modules in a simple project.
- **Migrating to Go Modules**: converting an existing project to use modules.
- **Publishing Go Modules**: how to make new versions of modules available to others.
- **Go Modules**: v2 and Beyond: creating and publishing major versions 2 and higher.
- **Keeping Your Modules Compatible**: how to keep your modules compatible with prior minor/patch versions.

> Tools
- **About the Go command**: why we wrote it, what it is, what it's not, and how to use it.
- **Go Doc Comments**: writing good program documentation
- **Debugging Go Code with GDB**
- **Data Race Detector**: a manual for the data race detector.
- **A Quick Guide to Go's Assembler**: an introduction to the assembler used by Go.
- **C? Go? Cgo!**: linking against C code with cgo..
- **Profiling Go Programs**: tools for measuring your code's CPU and memory usage
- **Introducing the Go Race Detector**: an introduction to the race detector.
- **Gopls: The language server for Go**: getting the most out your editor when working in Go.

## Accessing databases
- **Tutorial**: Accessing a relational database: Introduces the basics of accessing a relational database using Go and the `database/sql` package in the standard library.
- **Accessing relational databases**: An overview of Go's data access features.
- **Opening a database handle**: You use the Go database handle to execute database operations. Once you open a handle with database connection properties, the handle represents a connection pool it manages on your behalf.
- **Executing SQL statements that don't return data**: For SQL operations that might change the database, including SQL `INSERT`, `UPDATE`, and `DELETE`, you use `Exec` methods.
- **Querying for data**: For `SELECT` statements that return data from a query, using the `Query` or `QueryRow` method.
- **Using prepared statements**: Defining a prepared statement for repeated use can help your code run a bit faster by avoiding the overhead of re-creating the statement each time your code performs the database operation.
- **Executing transactions**: `sql.Tx` exports methods representing transaction-specific semantics, including `Commit` and `Rollback`, as well as methods you use to perform common database operations.
- **Canceling in-progress database operations**: Using [context.Context]
- **Managing connections**: For some advanced programs, you might need to tune connection pool parameters or work with connections explicitly.
- **Avoiding SQL injection risk**: You can avoid an SQL injection risk by providing SQL parameter values as `sql` package function arguments

# The Go Blog
* https://go.dev/blog/

# Go Wiki
* https://go.dev/wiki/
maintained by the Go community, includes articles about the Go language, tools, and other resources.

# Internal
* [internal.md](./internal.md)
