# The Go Programming Language Specification
* https://go.dev/ref/spec

Language version go1.25 (Feb 25, 2025)

# Introduction

# Notation

# Source code representation
## Characters
## Letters and digits

# Lexical elements
## Comments
## Tokens
## Semicolons
## Identifiers
## Keywords
## Operators and punctuation
## Integer literals
## Floating-point literals
## Imaginary literals
## Rune literals
## String literals

# Constants

# Variables
# Types
## Boolean types
## Numeric types
## String types
## Array types
## Slice types
## Struct types
## Pointer types
## Function types
## Interface types
## Map types
## Channel types



# Properties of types and values
## Representation of values
## Underlying types
## Type identity
## Assignability
## Representability
## Method sets

# Blocks

# Declarations and scope
## Label scopes
## Blank identifier
## Predeclared identifiers
## Exported identifiers
## Uniqueness of identifiers
## Constant declarations
## Iota
## Type declarations
## Type parameter declarations
## Variable declarations
## Short variable declarations
## Function declarations
## Method declarations

# Expressions
## Operands
## Qualified identifiers
## Composite literals
## Function literals
## Primary expressions
## Selectors
## Method expressions
## Method values
## Index expressions
## Slice expressions
## Type assertions
## Calls
## Passing arguments to `...` parameters
## Instantiations
## Type inference
## Operators
## Arithmetic operators
## Comparison operators
## Logical operators
## Address operators
## Receive operator


| Channel State   | Result               |
| :-------------- | :------------------- |
| `nil`           | Block                |
| open, not empty | Value                |
| open, empty     | Block                |
| closed          | default value, false |
| write only      | compilation error    |

## Conversions
## Constant expressions
## Order of evaluation

# Statements
## Terminating statements
## Empty statements
## Labeled statements
## Expression statements
## Send statements

| Channel State  | Result            |
| :------------- | :---------------- |
| `nil`          | Block             |
| open, full     | Block             |
| open, not full | Write value       |
| closed         | panic             |
| receive only   | compilation error |

## IncDec statements
## Assignment statements
## If statements
## Switch statements
## For statements
## Go statements


## Select statements


## Return statements
## Break statements
## Continue statements
## Goto statements
## Fallthrough statements
## Defer statements

# Built-in functions
## Appending to and copying slices
## Clear
## Close


| Channel State   | Result                                                                         |
| :-------------- | :----------------------------------------------------------------------------- |
| `nil`           | panic                                                                          |
| open, not empty | close channel, reads succeed until channel drained, then produce default value |
| open, empty     | close channel, reads produce default value                                     |
| closed          | panic                                                                          |
| receive only    | compilation error                                                              |

## Manipulating complex numbers
## Deletion of map elements
## Length and capacity
## Making slices, maps and channels
## Min and max
## Allocation
## Handling panics
## Bootstrapping

# Packages
## Source file organization
## Package clause
## Import declarations
## An example package

# Program initialization and execution
## The zero value
## Package initialization
## Program initialization
## Program execution

# Errors

# Run-time panics

# System considerations
## Package unsafe
## Size and alignment guarantees

# Appendix
## Language versions
## Type unification rules

# Summary

<!-- Put notes here. -->

data types
- **basic types**
  - booleans: `bool`
  - numeric types
    - intergers
      - signed: `int`, `int8`, `int16`, `int32`, `int64`.
      - unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`.
    - floating-point numbers: `float32`, `float64`. single and double-precision floating-point numbers.
    - complex numbers: `complex64`, `complex128`. complex numbers with `float32` and `float64` components, respectively
    - alias: `byte`(alias for `uint8`, commonly used for raw data), `rune`(alias for `int32`, used to represent a Unicode code point).
  - strings: `string`.
- **aggregate types**
  - arrays: `[n]Type`. fixed-size sequences of elements of the same `Type`.
  - structs: `struct {field1 Type1; field2 Type2; ...}`. collections of named fields, potentially of different types.
- **reference types**
  - pointers: `*Type`. stores the memory address of a value of `Type`.
  - slices: `[]Type`. dynamic-size, flexible views into arrays.
  - maps: `map[KeyType]ValueType`. unordered collections of key-value pairs.
  - channels: `chan Type`. concurrency primitives for communicating data between goroutines.
  - interfaces: `interface{}`, `interface{ Method1(); Method2(); ...}`. define a set of methods that a type must implement, enabling polymorphism.