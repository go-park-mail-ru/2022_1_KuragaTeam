package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
	"myapp/internal/mock"
	"myapp/internal/utils/images"
	"testing"
)

func TestMovieUsecase_GetMainMovie(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := internal.MainMovieInfoDTO{
		ID:          0,
		NamePicture: "name_picture.webp",
		Tagline:     "This is test movie",
		Picture:     "movie_picture.webp",
	}
	movieFromStorage := internal.MainMovieInfoDTO{
		ID:          movie.ID,
		NamePicture: movie.NamePicture,
		Tagline:     movie.Tagline,
		Picture:     movie.Picture,
	}

	movie.Picture, _ = images.GenerateFileURL(movie.Picture, "posters")
	movie.NamePicture, _ = images.GenerateFileURL(movie.NamePicture, "logos")

	tests := []struct {
		name          string
		storageMock   *mock.MockMovieStorage
		expected      internal.MainMovieInfoDTO
		expectedError bool
	}{
		{
			name: "Get main movie",
			storageMock: &mock.MockMovieStorage{
				GetRandomMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
					return &movieFromStorage, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
		{
			name: "Return error",
			storageMock: &mock.MockMovieStorage{
				GetRandomMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.storageMock, nil, nil, nil)
			mainMovie, err := r.GetMainMovie()

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, *mainMovie)
			}
		})
	}
}

func TestMovieUsecase_GetByID(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	person1 := internal.PersonInMovieDTO{
		ID:       1,
		Name:     "Актер 1",
		Photo:    "photo.webp",
		Position: "Актер",
	}
	person2 := internal.PersonInMovieDTO{
		ID:       2,
		Name:     "Актер 2",
		Photo:    "photo.webp",
		Position: "Актер",
	}
	movie := internal.Movie{
		ID:              1,
		Name:            "Movie1",
		NamePicture:     "name_picture.webp",
		Year:            2000,
		Duration:        "1 час 52 минуты",
		AgeLimit:        12,
		Description:     "Test film",
		KinopoiskRating: 9.1,
		Rating:          8.1,
		Tagline:         "Tagline of test film",
		Picture:         "picture.webp",
		Video:           "video.webp",
		Trailer:         "trailer.webp",
		Country:         []string{"Россия", "Франция"},
		Genre:           []string{"Драма", "Комедия"},
		Staff:           []internal.PersonInMovieDTO{person1, person2},
	}

	movieFromStorage := internal.Movie{
		ID:              movie.ID,
		Name:            movie.Name,
		NamePicture:     movie.NamePicture,
		Year:            movie.Year,
		Duration:        movie.Duration,
		AgeLimit:        movie.AgeLimit,
		Description:     movie.Description,
		KinopoiskRating: movie.KinopoiskRating,
		Rating:          movie.Rating,
		Tagline:         movie.Tagline,
		Picture:         movie.Picture,
		Video:           movie.Video,
		Trailer:         movie.Trailer,
		Country:         nil,
		Genre:           nil,
		Staff:           nil,
	}

	movie.Picture, _ = images.GenerateFileURL(movie.Picture, "posters")
	movie.Video, _ = images.GenerateFileURL(movie.Video, "movie")
	movie.Trailer, _ = images.GenerateFileURL(movie.Trailer, "trailers")
	movie.NamePicture, _ = images.GenerateFileURL(movie.NamePicture, "logos")

	tests := []struct {
		name               string
		movieStorageMock   *mock.MockMovieStorage
		countryStorageMock *mock.MockCountryStorage
		genreStorageMock   *mock.MockGenreStorage
		personStorageMock  *mock.MockPersonsStorage
		expected           internal.Movie
		expectedError      bool
	}{
		{
			name: "Get one movie",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*internal.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			personStorageMock: &mock.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]internal.PersonInMovieDTO, error) {
					return movie.Staff, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
		{
			name: "Movie storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*internal.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			personStorageMock: &mock.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]internal.PersonInMovieDTO, error) {
					return movie.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Country storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*internal.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			personStorageMock: &mock.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]internal.PersonInMovieDTO, error) {
					return movie.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*internal.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			personStorageMock: &mock.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]internal.PersonInMovieDTO, error) {
					return movie.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Staff storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*internal.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			personStorageMock: &mock.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]internal.PersonInMovieDTO, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.movieStorageMock, test.genreStorageMock, test.countryStorageMock, test.personStorageMock)
			mainMovie, err := r.GetByID(1)

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, *mainMovie)
			}
		})
	}
}

func TestMovieUsecase_GetRandom(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := internal.Movie{
		ID:              1,
		Name:            "Movie1",
		NamePicture:     "name_picture.webp",
		Year:            2000,
		Duration:        "1 час 52 минуты",
		AgeLimit:        12,
		Description:     "Test film",
		KinopoiskRating: 9.1,
		Rating:          8.1,
		Tagline:         "Tagline of test film",
		Picture:         "picture.webp",
		Video:           "video.webp",
		Trailer:         "trailer.webp",
		Country:         []string{"Россия", "Франция"},
		Genre:           []string{"Драма", "Комедия"},
		Staff:           nil,
	}

	movieFromStorage := internal.Movie{
		ID:              movie.ID,
		Name:            movie.Name,
		NamePicture:     movie.NamePicture,
		Year:            movie.Year,
		Duration:        movie.Duration,
		AgeLimit:        movie.AgeLimit,
		Description:     movie.Description,
		KinopoiskRating: movie.KinopoiskRating,
		Rating:          movie.Rating,
		Tagline:         movie.Tagline,
		Picture:         movie.Picture,
		Video:           movie.Video,
		Trailer:         movie.Trailer,
		Country:         nil,
		Genre:           nil,
		Staff:           nil,
	}

	movie.Picture, _ = images.GenerateFileURL(movie.Picture, "posters")
	movie.Video, _ = images.GenerateFileURL(movie.Video, "movie")
	movie.Trailer, _ = images.GenerateFileURL(movie.Trailer, "trailers")
	movie.NamePicture, _ = images.GenerateFileURL(movie.NamePicture, "logos")

	tests := []struct {
		name               string
		movieStorageMock   *mock.MockMovieStorage
		countryStorageMock *mock.MockCountryStorage
		genreStorageMock   *mock.MockGenreStorage
		expected           []internal.Movie
		expectedError      bool
	}{
		{
			name: "Get one movie",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]internal.Movie, error) {
					return []internal.Movie{movieFromStorage}, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			expected:      []internal.Movie{movie},
			expectedError: false,
		},
		{
			name: "Movie storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]internal.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Country storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]internal.Movie, error) {
					return []internal.Movie{movieFromStorage}, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Genre, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]internal.Movie, error) {
					return []internal.Movie{movieFromStorage}, nil
				},
			},
			countryStorageMock: &mock.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &mock.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.movieStorageMock, test.genreStorageMock, test.countryStorageMock, nil)
			mainMovie, err := r.GetRandom(1, 0)

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, mainMovie)
			}
		})
	}
}
