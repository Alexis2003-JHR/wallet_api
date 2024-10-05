package main

import (
	"wallet/internals/core/server"
)

const (
	port = 8014
)

func main() {
	server.InitServer(port)
}
