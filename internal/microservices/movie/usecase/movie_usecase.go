package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"myapp/internal/microservices/movie/proto"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	proto.UnimplementedMoviesServer
}

func (s *server) GetMovie(ctx context.Context, in *proto.GetMovieOptions) (*proto.GetMovieResponse, error) {
	return &proto.GetMovieResponse{MovieID: 1, Title: "Title"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMoviesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
