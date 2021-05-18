# Golang hello world

Run locally

```shellsession
$ go run main.go
$ go build
$ ./golang-hello-world
```

Run with Docker

```shellsession
$ docker build -t golang-hello-world .
$ docker run -p8080:8080 golang-hello-world
