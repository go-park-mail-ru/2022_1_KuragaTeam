package usecase

import (
	"myapp/internal"
	"myapp/internal/country"
	"myapp/internal/domain/staff"
	"myapp/internal/genre"
	"myapp/internal/movie"
)

type service struct {
	movieStorage   movie.Storage
	genreStorage   genre.Storage
	countryStorage country.Storage
	staffStorage   staff.Storage
}

func NewService(movieStorage movie.Storage, genreStorage genre.Storage,
	countryStorage country.Storage, staffStorage staff.Storage) movie.Service {
	return &service{movieStorage: movieStorage, genreStorage: genreStorage,
		countryStorage: countryStorage, staffStorage: staffStorage}
}

func (s *service) GetByID(id int) (*internal.Movie, error) {
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

	selectedMovie.Staff, err = s.staffStorage.GetByMovieID(selectedMovie.ID)
	if err != nil {
		return nil, err
	}

	selectedMovie.Rating = 8.1 // пока что просто замокано
	return selectedMovie, nil
}
func (s *service) GetRandom(limit, offset int) ([]internal.Movie, error) {
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

		movies[i].Rating = 8.1 // пока что просто замокано
	}
	return movies, err
}

func (s *service) GetMainMovie() (*internal.MainMovieInfoDTO, error) {
	return s.movieStorage.GetRandomMovie()
}
