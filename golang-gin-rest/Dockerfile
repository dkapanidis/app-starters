# Build image
FROM golang:1.14.2-alpine AS build
RUN apk add upx
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go test -v ./...
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/app .
RUN upx ./out/app

# Runtime image
FROM scratch
COPY --from=build /app/out/app /usr/local/bin/app
EXPOSE 8080
CMD ["/usr/local/bin/app"]

