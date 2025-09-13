# hello_world

skaffold:
```shell
$ mkdir hello_world
$ cd hello_world 
$ go mod init hello_world
```

dependency:
```shell
# https://github.com/Code-Hex/Neo-cowsay
$ go get github.com/Code-Hex/Neo-cowsay/v2

$ go mod tidy
```

build and run:
```shell
$ make
go fmt ./...
go vet ./...
go build

$ ./hello_world.exe 
Hello, 世界!
 ______________ 
< Hello, 世界! >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

cleanup:
```shell
$ make clean
rm -rf hello_world*
```