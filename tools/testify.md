# testify
* https://github.com/stretchr/testify

> A toolkit with common assertions and mocks that plays nicely with the standard library.
>
> We currently support the most recent major Go versions from 1.19 onward.

Features
- Easy assertions
- Mocking
- Testing suite interfaces and functions

```shell
go get github.com/stretchr/testify

# packages
github.com/stretchr/testify/assert
github.com/stretchr/testify/require
github.com/stretchr/testify/mock
github.com/stretchr/testify/suite
github.com/stretchr/testify/http (deprecated)
```

# `assert` package

Every assert func returns a bool indicating whether the assertion was successful or not, this is useful for if you want to go on making further assertions under certain conditions.

```go
assert.Equal(t, 123, 123, "they should be equal")
assert.NotEqual(t, 123, 456, "they should not be equal")
assert.Nil(t, object)

// assert many times
assert := assert.New(t)
assert.Equal(123, 123, "they should be equal")
```

# `require` package

The `require` package provides same global functions as the `assert` package, but instead of returning a boolean result they terminate current test. 

These functions must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Otherwise race conditions may occur.

# `mock` package

The `mock` package provides a mechanism for easily writing mock objects that can be used in place of real objects when writing test code.

You can use the [mockery tool](https://vektra.github.io/mockery/latest/) to autogenerate the mock code against an interface as well, making using mocks much quicker.

# `suite` package

The `suite` package provides functionality that you might be used to from more common object-oriented languages. With it, you can build a testing suite as a struct, build setup/teardown methods and testing methods on your struct, and run them with 'go test' as per normal.

The `suite` package does not support parallel tests.

# See Also
* [mockery](https://github.com/vektra/mockery): A mock code autogenerator for Go. mockery provides the ability to easily generate mocks for Golang interfaces using the `stretchr/testify/mock` package. It removes the boilerplate coding required to use mocks.
