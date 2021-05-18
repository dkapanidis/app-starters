# Golang hello world

> Runs a simple Go Web server without dependencies.

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
```

Run with Pack & Docker

```shellsession
$ pack build --builder=gcr.io/buildpacks/builder golang-hello-world
$ docker run -p8080:8080 golang-hello-world
```

Run with Skaffold

```shellsession
$ skaffold dev --port-forward
```

Open http://localhost:8080

