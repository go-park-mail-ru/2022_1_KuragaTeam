package movie

import (
	"myapp/internal/adapters/api/movie"
	"myapp/internal/domain"
	"myapp/internal/domain/country"
	"myapp/internal/domain/genre"
)

type service struct {
	movieStorage   Storage
	genreStorage   genre.Storage
	countryStorage country.Storage
}

func NewService(movieStorage Storage, genreStorage genre.Storage, countryStorage country.Storage) movie.Service {
	return &service{movieStorage: movieStorage, genreStorage: genreStorage, countryStorage: countryStorage}
}

func (s *service) GetByID(id int) (*domain.Movie, error) {
	selectedMovie, err := s.movieStorage.GetOne(id)
	if err != nil {
		return nil, err
	}
	selectedMovie.Genre, err = s.genreStorage.GetByMovieID(selectedMovie.ID)
	if err != nil {
		return nil, err
	}

	selectedMovie.Country, err = s.countryStorage.GetByMovieID(selectedMovie.ID)
	if err != nil {
		return nil, err
	}

	return selectedMovie, nil
}
func (s *service) GetRandom(limit, offset int) ([]domain.Movie, error) {
	movies, err := s.movieStorage.GetRandomMovies(limit, offset)
	for i := 0; i < len(movies); i++ {
		movies[i].Genre, err = s.genreStorage.GetByMovieID(movies[i].ID)
		if err != nil {
			movies[i].Genre = append(movies[i].Genre, err.Error())
		}
		movies[i].Country, err = s.countryStorage.GetByMovieID(movies[i].ID)
		if err != nil {
			movies[i].Country = append(movies[i].Country, err.Error())
		}
	}
	return movies, err
}

func (s *service) GetMainMovie() (*domain.MainMovieInfoDTO, error) {
	return s.movieStorage.GetRandomMovie()
}
