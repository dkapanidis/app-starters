# React hello world

> Runs a simple React app with typescript.

Bootstrap

```shellsession
npx create-react-app <APP_NAME> --template typescript
```

Run locally

```shellsession
$ yarn start
```

Run with Docker

```shellsession
# production build (multi-stage without source code, using nginx)
$ docker build -t react-hello-world .
$ docker run -p3000:80 react-hello-world

# dev build
$ docker build -f Dockerfile.dev -t react-hello-world .
$ docker run -v $(pwd):/app -p3000:80 react-hello-world
```

Run with Pack & Docker

```shellsession
$ pack build --builder=gcr.io/buildpacks/builder react-hello-world
$ docker run -p3000:3000 react-hello-world
```

Run with Skaffold

```shellsession
$ skaffold dev --port-forward
```

Open http://localhost:3000

