# Golang Gin

> Runs a simple Go web server using Gin framework, with unit tests.

Run locally

```shellsession
$ go run main.go
$ go build
$ ./golang-gin
```

Run with Docker

```shellsession
$ docker build -t golang-gin .
$ docker run -p8080:8080 golang-gin
```

Run with Skaffold

```shellsession
$ skaffold dev --port-forward
```

Open http://localhost:8080
