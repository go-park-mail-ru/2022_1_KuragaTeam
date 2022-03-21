package usecase

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/moviesCompilations"
)

type service struct {
	MCStorage moviesCompilations.Storage
}

func NewService(MCStorage moviesCompilations.Storage) moviesCompilations.Service {
	return &service{MCStorage: MCStorage}
}

func (s *service) GetMainCompilations(context echo.Context) ([]moviesCompilations.MovieCompilation, error) {
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

func (s *service) GetByMovieID(context echo.Context, limit int) (moviesCompilations.MovieCompilation, error) {
	return moviesCompilations.MovieCompilation{}, nil
}

func (s *service) GetByGenre(context echo.Context, genreID int) (moviesCompilations.MovieCompilation, error) {
	return s.MCStorage.GetByGenre(genreID)
}
