package moviesCompilations

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api/moviesCompilations"
	"myapp/internal/domain"
)

type service struct {
}

func NewService() moviesCompilations.Service {
	return &service{}
}

func (s *service) GetMainCompilations(context echo.Context) ([]domain.MovieCompilation, error) {
	movieCompilations := []domain.MovieCompilation{
		{
			Name: "Популярное",
			Movies: []domain.Movie{
				{
					ID:          0,
					Name:        "Звездные войны1",
					Description: "",
					Picture:     "star.png",
					Video:       "/",
					Trailer:     "",
					Genre:       []string{"Фантастика1"},
				},
				{
					ID:          0,
					Name:        "Звездные войны2",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика2"},
				},
				{
					ID:          0,
					Name:        "Звездные войны3",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика3"},
				},
				{
					ID:          0,
					Name:        "Звездные войны4",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика4"},
				},
			},
		},
		{
			Name: "Топ",
			Movies: []domain.Movie{
				{
					ID:          0,
					Name:        "Звездные войны#1",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика"},
				},
				{
					ID:          0,
					Name:        "Звездные войны#2",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика"},
				},
				{
					ID:          0,
					Name:        "Звездные войны#3",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика"},
				},
			},
		},
		{
			Name: "Семейное",
			Movies: []domain.Movie{
				{
					ID:          0,
					Name:        "Звездные войны#1",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика"},
				},
				{
					ID:          0,
					Name:        "Звездные войны#2",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика"},
				},
				{
					ID:          0,
					Name:        "Звездные войны#3",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика"},
				},
				{
					ID:          0,
					Name:        "Звездные войны4",
					Description: "",
					Picture:     "",
					Video:       "",
					Trailer:     "",
					Genre:       []string{"Фантастика4"},
				},
			},
		},
	}

	return movieCompilations, nil
}
func (s *service) GetByMovieID(context echo.Context, limit int) (domain.MovieCompilation, error) {
	return domain.MovieCompilation{}, nil
}
