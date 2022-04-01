package main

import (
	"nihago-users/cmd/grpcserver"
)

func main() {
	var server grpcserver.Server
	server.Run()
}
