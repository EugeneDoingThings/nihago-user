package main

import "nihago-user/cmd/grpcserver"

func main() {
	var server grpcserver.Server
	server.Run()
}
