# Go Modules
* https://go.dev/ref/mod

contents
- Introduction
- Modules, packages, and versions
- go.mod files
- Minimal version selection (MVS)
- Module graph pruning
- Workspaces
- Compatibility with non-module repositories
- Module-aware commands
- Module proxies
- Version control systems
- Module zip files
- Private modules
- Module cache
- Authenticating modules
- Environment variables
- Glossary

# `go.mod`
* https://go.dev/doc/modules/gomod-ref

Each Go module is defined by a go.mod file that describes the module’s properties, including its dependencies on other modules and on versions of Go.

These properties include:
- The current module’s **module path**. This should be a location from which the module can be downloaded by Go tools, such as the module code’s repository location. This serves as a unique identifier, when combined with the module’s version number. It is also the prefix of the package path for all packages in the module.
- The **minimum version of Go** required by the current module.
- A list of minimum versions of other **modules required by the current module**.
- Instructions, optionally, to **replace** a required module with another module version or a local directory, or to **exclude** a specific version of a required module.

directives
- `module`: Declares the module’s module path, which is the module’s unique identifier (when combined with the module version number). 
- `go`: Indicates that the module was written assuming the semantics of the Go version specified by the directive.
- `toolchain`: Declares a suggested Go toolchain to use with this module. Only takes effect when the module is the main module and the default toolchain is older than the suggested toolchain.
- `godebug`: Indicates the default `GODEBUG` settings to be applied to the main packages of this module. These override any toolchain defaults, and are overridden by explicit `//go:debug` lines in main packages.
- `require`: Declares a module as a dependency of the current module, specifying the minimum version of the module required.
- `tool`: Adds a package as a dependency of the current module, and makes it available to run with `go tool` when the current working directory is within this module.
- `replace`: Replaces the content of a module at a specific version (or all versions) with another module version or with a local directory. Go tools will use the replacement path when resolving the dependency.
- `exclude`: Specifies a module or module version to exclude from the current module’s dependency graph.
- `retract`: Indicates that a version or range of versions of the module defined by `go.mod` *should not be depended upon*. A retract directive is useful when a version was published prematurely or a severe problem was discovered after the version was published.

# See Also
* [How to Write Go Code](https://go.dev/doc/code.html)
* [Using Go Modules](https://go.dev/blog/using-go-modules)
* [Internal packages - Go 1.4 Release Notes](https://go.dev/doc/go1.4#internalpackages)
* [Simple Go project layout with modules](Simple Go project layout with modules) - 2019-10-01
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout): Standard Go Project Layout