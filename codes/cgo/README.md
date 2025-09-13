# cgo
Examples in 'Learning Go'.

Run with Windows WSL Unbuntu 22.04.
```shell
✗ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 22.04 LTS
Release:        22.04
Codename:       jammy
```

# go_call_c
```shell
$ mkdir -p go_call_c && cd go_call_c && go mod init go_call_c
$ go build
$ ./go_call_c
a: 3, b: 2, sum 5
5
10
200
```

# c_call_go
```shell
$ mkdir -p c_call_go && cd c_call_go && go mod init c_call_go
$ go build
$  ./c_call_go
8
```

# handle: pointers
```shell
$ mkdir -p handle && cd handle && go mod init handle
$ go build
$ ./handle
John 42
```