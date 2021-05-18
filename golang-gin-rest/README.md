# Golang Gin

> Runs a simple Go web server implementing a REST example using Gin framework, with unit tests.

Run locally

```shellsession
$ go run main.go
$ go build
$ ./golang-gin-rest
```

Run with Docker

```shellsession
$ docker build -t golang-gin-rest .
$ docker run -p8080:8080 golang-gin-rest
```

Run with Skaffold

```shellsession
$ skaffold dev --port-forward
```

Open http://localhost:8080
