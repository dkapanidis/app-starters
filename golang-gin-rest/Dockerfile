# Build image
FROM golang:1.15.5-alpine AS build
RUN apk add upx gcc musl-dev
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go test -v ./...
RUN go build -ldflags="-s -w" -o ./out/app .
RUN upx ./out/app

# Runtime image
FROM golang:1.15.5-alpine
COPY --from=build /app/out/app /usr/local/bin/app
EXPOSE 8080
CMD ["/usr/local/bin/app"]

