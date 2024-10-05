package main

import (
	"wallet/internals/server"
)

const (
	port = 8014
)

func main() {
	server.InitServer(port)
}
