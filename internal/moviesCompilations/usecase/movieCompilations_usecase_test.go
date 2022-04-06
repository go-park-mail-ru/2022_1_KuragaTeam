package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"myapp/internal/moviesCompilations"
	"myapp/internal/utils/images"
	mock2 "myapp/mock"
	"testing"
)

func TestMovieCompilationsUsecase_GetMainCompilations(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie1FromStorage := moviesCompilations.Movie{
		ID:      movie1.ID,
		Name:    movie1.Name,
		Genre:   nil,
		Picture: movie1.Picture,
	}
	movie1.Picture, _ = images.GenerateFileURL(movie1.Picture, "posters")

	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}
	movie2FromStorage := moviesCompilations.Movie{
		ID:      movie2.ID,
		Name:    movie2.Name,
		Genre:   nil,
		Picture: movie2.Picture,
	}
	movie2.Picture, _ = images.GenerateFileURL(movie2.Picture, "posters")

	movie3 := moviesCompilations.Movie{
		ID:      3,
		Name:    "Movie3",
		Genre:   []string{"Детектив", "Криминал"},
		Picture: "picture_name3.webp",
	}
	movie3FromStorage := moviesCompilations.Movie{
		ID:      movie3.ID,
		Name:    movie3.Name,
		Genre:   nil,
		Picture: movie3.Picture,
	}
	movie3.Picture, _ = images.GenerateFileURL(movie3.Picture, "posters")

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1FromStorage, movie2FromStorage},
	}
	MC2 := moviesCompilations.MovieCompilation{
		Name:   "Test MC2",
		Movies: []moviesCompilations.Movie{movie2FromStorage, movie3FromStorage},
	}
	MC3 := moviesCompilations.MovieCompilation{
		Name:   "Test MC3",
		Movies: []moviesCompilations.Movie{movie3FromStorage, movie1FromStorage},
	}
	MC4 := moviesCompilations.MovieCompilation{
		Name:   "Test MC3",
		Movies: []moviesCompilations.Movie{movie1FromStorage, movie2FromStorage, movie3FromStorage},
	}

	tests := []struct {
		name             string
		MCstorageMock    *mock2.MockMovieCompilationStorage
		genreStorageMock *mock2.MockGenreStorage
		expected         []moviesCompilations.MovieCompilation
		expectedError    bool
	}{
		{
			name: "Get main compilations",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC2, nil
				},
				GetByGenreFunc: func(genreID int) (moviesCompilations.MovieCompilation, error) {
					return MC3, nil
				},
				GetByCountryFunc: func(countryID int) (moviesCompilations.MovieCompilation, error) {
					return MC4, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				{
					Name:   MC1.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2},
				},
				{
					Name:   MC2.Name,
					Movies: []moviesCompilations.Movie{movie2, movie3},
				},
				{
					Name:   MC3.Name,
					Movies: []moviesCompilations.Movie{movie3, movie1},
				},
				{
					Name:   MC4.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2, movie3},
				},
			},
			expectedError: false,
		},
		{
			name: "MC storage top func return error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC2, nil
				},
				GetByGenreFunc: func(genreID int) (moviesCompilations.MovieCompilation, error) {
					return MC3, nil
				},
				GetByCountryFunc: func(countryID int) (moviesCompilations.MovieCompilation, error) {
					return MC4, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				{
					Name:   MC1.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2},
				},
				{
					Name:   MC2.Name,
					Movies: []moviesCompilations.Movie{movie2, movie3},
				},
				{
					Name:   MC3.Name,
					Movies: []moviesCompilations.Movie{movie3, movie1},
				},
				{
					Name:   MC4.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2, movie3},
				},
			},
			expectedError: true,
		},
		{
			name: "MC storage top by year func return error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
				GetByGenreFunc: func(genreID int) (moviesCompilations.MovieCompilation, error) {
					return MC3, nil
				},
				GetByCountryFunc: func(countryID int) (moviesCompilations.MovieCompilation, error) {
					return MC4, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				{
					Name:   MC1.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2},
				},
				{
					Name:   MC2.Name,
					Movies: []moviesCompilations.Movie{movie2, movie3},
				},
				{
					Name:   MC3.Name,
					Movies: []moviesCompilations.Movie{movie3, movie1},
				},
				{
					Name:   MC4.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2, movie3},
				},
			},
			expectedError: true,
		},
		{
			name: "MC storage get by genre func return error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC2, nil
				},
				GetByGenreFunc: func(genreID int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
				GetByCountryFunc: func(countryID int) (moviesCompilations.MovieCompilation, error) {
					return MC4, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				{
					Name:   MC1.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2},
				},
				{
					Name:   MC2.Name,
					Movies: []moviesCompilations.Movie{movie2, movie3},
				},
				{
					Name:   MC3.Name,
					Movies: []moviesCompilations.Movie{movie3, movie1},
				},
				{
					Name:   MC4.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2, movie3},
				},
			},
			expectedError: true,
		},
		{
			name: "MC storage get by country func return error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC2, nil
				},
				GetByGenreFunc: func(genreID int) (moviesCompilations.MovieCompilation, error) {
					return MC3, nil
				},
				GetByCountryFunc: func(countryID int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				{
					Name:   MC1.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2},
				},
				{
					Name:   MC2.Name,
					Movies: []moviesCompilations.Movie{movie2, movie3},
				},
				{
					Name:   MC3.Name,
					Movies: []moviesCompilations.Movie{movie3, movie1},
				},
				{
					Name:   MC4.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2, movie3},
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage returns error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC2, nil
				},
				GetByGenreFunc: func(genreID int) (moviesCompilations.MovieCompilation, error) {
					return MC3, nil
				},
				GetByCountryFunc: func(countryID int) (moviesCompilations.MovieCompilation, error) {
					return MC4, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				{
					Name:   MC1.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2},
				},
				{
					Name:   MC2.Name,
					Movies: []moviesCompilations.Movie{movie2, movie3},
				},
				{
					Name:   MC3.Name,
					Movies: []moviesCompilations.Movie{movie3, movie1},
				},
				{
					Name:   MC4.Name,
					Movies: []moviesCompilations.Movie{movie1, movie2, movie3},
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.MCstorageMock, test.genreStorageMock)
			mainMC, err := r.GetMainCompilations()

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, mainMC)
			}
		})
	}
}
func TestMovieCompilationsUsecase_GetByPerson(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie1FromStorage := moviesCompilations.Movie{
		ID:      movie1.ID,
		Name:    movie1.Name,
		Genre:   nil,
		Picture: movie1.Picture,
	}
	movie1.Picture, _ = images.GenerateFileURL(movie1.Picture, "posters")

	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}
	movie2FromStorage := moviesCompilations.Movie{
		ID:      movie2.ID,
		Name:    movie2.Name,
		Genre:   nil,
		Picture: movie2.Picture,
	}
	movie2.Picture, _ = images.GenerateFileURL(movie2.Picture, "posters")

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1FromStorage, movie2FromStorage},
	}

	tests := []struct {
		name             string
		MCstorageMock    *mock2.MockMovieCompilationStorage
		genreStorageMock *mock2.MockGenreStorage
		expected         moviesCompilations.MovieCompilation
		expectedError    bool
	}{
		{
			name: "Get compilation by person",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetByPersonFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expected: moviesCompilations.MovieCompilation{
				Name:   MC1.Name,
				Movies: []moviesCompilations.Movie{movie1, movie2},
			},
			expectedError: false,
		},
		{
			name: "MC storage returns error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetByPersonFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					switch id {
					case 1:
						return []string{"Боевик", "Триллер"}, nil
					case 2:
						return []string{"Фантастика", "Семейный"}, nil
					case 3:
						return []string{"Детектив", "Криминал"}, nil
					}
					return nil, nil
				},
			},
			expectedError: true,
		},
		{
			name: "MC storage returns error",
			MCstorageMock: &mock2.MockMovieCompilationStorage{
				GetByPersonFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			genreStorageMock: &mock2.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.MCstorageMock, test.genreStorageMock)
			mainMC, err := r.GetByPerson(1)

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, mainMC)
			}
		})
	}
}
