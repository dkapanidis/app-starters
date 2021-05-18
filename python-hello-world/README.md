# Golang hello world

> Runs a simple Python Web server without dependencies.

Run locally

```shellsession
$ python3 src/app.py
```

Run with Docker

```shellsession
$ docker build -t python-hello-world .
$ docker run -p5000:5000 python-hello-world
```

Run with Pack & Docker

```shellsession
$ pack build --builder=gcr.io/buildpacks/builder python-hello-world
$ docker run -p5000:5000 python-hello-world
```

Run with Skaffold

```shellsession
$ skaffold dev --port-forward
```

Open http://localhost:5000

