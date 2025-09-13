# Containerize Go Applications
* https://www.infoq.cn/article/cZRp8lJTtdgxaNfYcRTw


```shell
$ go mod init containerization
$ docker build -t go-containerization .
$ docker run --name go-containerization -it go-containerization
Hello Container!

# cleanup
$ docker container rm go-containerization
$ docker image rm go-containerization
```