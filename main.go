package main

import (
	"flag"

	server "github.com/sgace/server"
)

func main() {
	flag.Parse()

	server.Start()
}
