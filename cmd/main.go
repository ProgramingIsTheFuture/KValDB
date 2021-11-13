package main

import (
	"KValDB/external/gRPC/server"
	"KValDB/internal/handlers"
)

func main() {
	kval := handlers.NewKVal()

	server := server.NewServer(kval)

	server.Serve("0.0.0.0:50051")

}
