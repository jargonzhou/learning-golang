# learning_go

```shell
$ go mod init learning_go
```

# Dependencies

| Dependency         | Description                                                                               |
| :----------------- | :---------------------------------------------------------------------------------------- |
| testify            | test assertions and mocks                                                                 |
| google/uuid        | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. |
| cockroachdb/errors | Go error library with error portability over the network                                  |

```shell
go get github.com/stretchr/testify
go get github.com/google/uuid
go get github.com/cockroachdb/errors
```

# Contents

- [Data Structures](./datastructure/doc.go)
- [Generics](./generics/doc.go)
- [Tests](./tests/doc.go)
- [Concurrency](./concurrency/doc.go)
- [Tools](./tools/doc.go)

# See Also
- [Module, Package, Project](../../tools/tools.md#module-package-project)ls