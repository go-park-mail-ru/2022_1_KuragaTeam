package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/microservices/movie/usecase"

	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMoviesServer(s, usecase.NewService())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
