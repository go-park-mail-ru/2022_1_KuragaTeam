package movie

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api/movie"
	"myapp/internal/domain"
	"myapp/internal/domain/genre"
)

type service struct {
	movieStorage Storage
	genreStorage genre.Storage
}

func NewService(movieStorage Storage, genreStorage genre.Storage) movie.Service {
	return &service{movieStorage: movieStorage, genreStorage: genreStorage}
}

func (s *service) GetByID(context echo.Context, id int) (*domain.Movie, error) {
	return nil, nil
}
func (s *service) GetRandom(context echo.Context, limit int) ([]domain.Movie, error) {
	movies, err := s.movieStorage.GetRandom(limit)
	for i := 0; i < len(movies); i++ {
		movies[i].Genre, err = s.genreStorage.GetByMovieID(movies[i].ID)
		if err != nil {
			movies[i].Genre = append(movies[i].Genre, err.Error())
		}
	}
	return movies, err
}
