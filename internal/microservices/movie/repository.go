package movie

import (
	"myapp/internal/microservices/movie/proto"
)

type Storage interface {
	GetOne(id int) (*proto.Movie, error)
	GetAllMovies(limit, offset int) ([]*proto.Movie, error)
	GetSeasonsAndEpisodes(seriesID int) ([]*proto.Season, error)
	GetRandomMovie() (*proto.MainMovie, error)
	GetMovieRating(movieID int) (*GetMovieRatingAnswer, error)
	AddMovieRating(options *proto.AddRatingOptions) error
	ChangeMovieRating(options *proto.AddRatingOptions) error
	CheckRatingExists(options *proto.AddRatingOptions) (*CheckRatingExistsAnswer, error)
}

type GetMovieRatingAnswer struct {
	RatingSum   int
	RatingCount int
}

type CheckRatingExistsAnswer struct {
	Exists bool
	Rating int
}
