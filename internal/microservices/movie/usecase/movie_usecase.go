package usecase

import (
	"golang.org/x/net/context"
	"myapp/internal/microservices/movie/proto"
)

type server struct {
	proto.UnimplementedMoviesServer
}

func NewService() *server {
	return &server{}
}

func (s *server) GetByID(ctx context.Context, in *proto.GetMovieOptions) (*proto.Movie, error) {
	return &proto.Movie{}, nil
}
