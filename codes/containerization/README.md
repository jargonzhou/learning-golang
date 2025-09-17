# Containerize Go Applications
* https://www.infoq.cn/article/cZRp8lJTtdgxaNfYcRTw

setup
```shell
$ go mod init containerization
$ go build
$ go version -m .
containerization.exe: go1.25.1
        path    containerization
        mod     containerization        (devel)
        build   -buildmode=exe
        build   -compiler=gc
        build   CGO_ENABLED=1
        build   CGO_CFLAGS=
        build   CGO_CPPFLAGS=
        build   CGO_CXXFLAGS=
        build   CGO_LDFLAGS=
        build   GOARCH=amd64
        build   GOOS=windows
        build   GOAMD64=v1
        build   vcs=git
        build   vcs.revision=6d6df2d86bba744eabf4ec1b1e71f4035d73c882
        build   vcs.time=2025-09-13T17:22:39Z
        build   vcs.modified=true
```

docker
```shell
$ docker build -t go-containerization .
$ docker run --name go-containerization -it go-containerization
Hello Container!
```

cleanup
```shell
$ docker container rm go-containerization
$ docker image rm go-containerization
```