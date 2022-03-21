package usecase

import (
	"myapp/internal/genre"
	"myapp/internal/moviesCompilations"
)

type service struct {
	MCStorage    moviesCompilations.Storage
	genreStorage genre.Storage
}

func NewService(MCStorage moviesCompilations.Storage, genreStorage genre.Storage) moviesCompilations.Service {
	return &service{MCStorage: MCStorage, genreStorage: genreStorage}
}

func (s *service) GetMainCompilations() ([]moviesCompilations.MovieCompilation, error) {
	movieCompilations := []moviesCompilations.MovieCompilation{
		{
			Name: "Популярное",
			Movies: []moviesCompilations.Movie{
				{
					ID:      0,
					Name:    "Звездные войны1",
					Picture: "star.png",
					Genre:   []string{"Фантастика1"},
				},
				{
					ID:      0,
					Name:    "Звездные войны2",
					Picture: "",
					Genre:   []string{"Фантастика2"},
				},
				{
					ID:      0,
					Name:    "Звездные войны3",
					Picture: "",
					Genre:   []string{"Фантастика3"},
				},
				{
					ID:      0,
					Name:    "Звездные войны4",
					Picture: "",
					Genre:   []string{"Фантастика4"},
				},
			},
		},
		{
			Name: "Топ",
			Movies: []moviesCompilations.Movie{
				{
					ID:      0,
					Name:    "Звездные войны#1",
					Picture: "",
					Genre:   []string{"Фантастика"},
				},
				{
					ID:      0,
					Name:    "Звездные войны#2",
					Picture: "",
					Genre:   []string{"Фантастика"},
				},
				{
					ID:      0,
					Name:    "Звездные войны#3",
					Picture: "",
					Genre:   []string{"Фантастика"},
				},
			},
		},
		{
			Name: "Семейное",
			Movies: []moviesCompilations.Movie{
				{
					ID:      0,
					Name:    "Звездные войны#1",
					Picture: "",
					Genre:   []string{"Фантастика"},
				},
				{
					ID:      0,
					Name:    "Звездные войны#2",
					Picture: "",
					Genre:   []string{"Фантастика"},
				},
				{
					ID:      0,
					Name:    "Звездные войны#3",
					Picture: "",
					Genre:   []string{"Фантастика"},
				},
				{
					ID:      0,
					Name:    "Звездные войны4",
					Picture: "",
					Genre:   []string{"Фантастика4"},
				},
			},
		},
	}

	return movieCompilations, nil
}

func (s *service) GetByMovie(movieID int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByMovie(movieID)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	for i, _ := range MC.Movies {
		MC.Movies[i].Genre, err = s.genreStorage.GetByMovieID(MC.Movies[i].ID)
		if err != nil {
			return moviesCompilations.MovieCompilation{}, err
		}
	}
	return MC, nil
}

func (s *service) GetByGenre(genreID int) (moviesCompilations.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByGenre(genreID)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}
	for i, _ := range MC.Movies {
		MC.Movies[i].Genre, err = s.genreStorage.GetByMovieID(MC.Movies[i].ID)
		if err != nil {
			return moviesCompilations.MovieCompilation{}, err
		}
	}
	return MC, nil
}
