# Golang Flask

> Runs a simple Python Web server with Flask.

Run locally

```shellsession
$ python3 -m flask run
```

Run with Docker

```shellsession
$ docker build -t python-flask .
$ docker run -p5000:5000 python-flask
```

Run with Skaffold

```shellsession
$ skaffold dev --port-forward
```

Open http://localhost:5000

