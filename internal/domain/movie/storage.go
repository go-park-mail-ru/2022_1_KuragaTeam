package movie

import "myapp/internal/domain"

type Storage interface {
	GetOne(id int) (*domain.Movie, error)
	GetRandomMovies(limit, offset int) ([]domain.Movie, error)
	GetRandomMovie() (*domain.MainMovieInfoDTO, error)
}
