package server

import (
	"KValDB/external/gRPC/proto"
	"KValDB/external/gRPC/service"
	"KValDB/internal/handlers"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	KValDB handlers.KValDB
}

func NewServer(kv handlers.KValDB) *Server {
	return &Server{KValDB: kv}
}

func (s *Server) Serve(addr string) {
	lst, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	service := service.NewService(s.KValDB)

	grpcServer := grpc.NewServer()
	proto.RegisterKValDBServer(grpcServer, service)

	fmt.Println("Running the server - ", addr)
	if err = grpcServer.Serve(lst); err != nil {
		panic(err)
	}

}
