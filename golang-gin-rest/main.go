package main

import (
	"flag"

	"github.com/dkapanidis/app-starters/golang-gin-rest/pkg/routers"
)

func main() {
	r := routers.SetupRouter()

	addr := flag.String("address", "0.0.0.0:8080", "server address (default: '0.0.0.0:8080')")
	flag.Parse()

	r.Run(*addr)
}
