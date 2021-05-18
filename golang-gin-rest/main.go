package main

import (
	"flag"

	"github.com/dkapanidis/app-starters/golang-gin-rest/pkg/routers"
)

func main() {
	r := routers.SetupRouter()

	addr := flag.String("address", "localhost:8080", "server address (default: 'localhost:8080')")
	flag.Parse()

	r.Run(*addr)
}
