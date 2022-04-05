package usecase

import (
	"myapp/internal/genre"
	"myapp/internal/moviesCompilations"
	"myapp/internal/utils/images"
)

type service struct {
	MCStorage    moviesCompilations.Storage
	genreStorage genre.Storage
}

func NewService(MCStorage moviesCompilations.Storage, genreStorage genre.Storage) moviesCompilations.Service {
	return &service{MCStorage: MCStorage, genreStorage: genreStorage}
}

func (s *service) fillGenres(MC *moviesCompilations.MovieCompilation) error {
	for i := 0; i < len(MC.Movies); i++ {
		var err error
		MC.Movies[i].Genre, err = s.genreStorage.GetByMovieID(MC.Movies[i].ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) concatUrls(MC *moviesCompilations.MovieCompilation) error {
	var err error
	for i, _ := range MC.Movies {
		MC.Movies[i].Picture, err = images.GenerateFileURL(MC.Movies[i].Picture, "posters")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) GetMainCompilations() ([]moviesCompilations.MovieCompilation, error) {

	MC := make([]moviesCompilations.MovieCompilation, 0)

	nextMC, err := s.GetTop(10)
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	nextMC, err = s.GetTopByYear(2011)
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	nextMC, err = s.GetByGenre(2) // Боевик
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	nextMC, err = s.GetByCountry(3) // США
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	return MC, nil
}

func (s *service) GetByMovie(movieID int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByMovie(movieID)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.fillGenres(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.concatUrls(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	return MC, nil
}

func (s *service) GetByGenre(genreID int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByGenre(genreID)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.fillGenres(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.concatUrls(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	return MC, nil
}

func (s *service) GetByCountry(countryID int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByCountry(countryID)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.fillGenres(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.concatUrls(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	return MC, nil
}

func (s *service) GetByPerson(personID int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByPerson(personID)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.fillGenres(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.concatUrls(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	return MC, nil
}
func (s *service) GetTopByYear(year int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetTopByYear(year)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.fillGenres(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.concatUrls(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	return MC, nil
}
func (s *service) GetTop(limit int) (moviesCompilations.MovieCompilation, error) {
	if limit > 10 {
		limit = 10
	}

	MC, err := s.MCStorage.GetTop(limit)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.fillGenres(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	err = s.concatUrls(&MC)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	return MC, nil
}
