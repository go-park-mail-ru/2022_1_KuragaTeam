package usecase

import (
	"errors"
	"myapp/internal"
	countryRepository "myapp/internal/country/repository"
	genre "myapp/internal/genre/repository"
	"myapp/internal/microservices/movie"
	"myapp/internal/microservices/movie/proto"
	mock "myapp/internal/microservices/movie/repository"
	"myapp/internal/microservices/movie/utils/images"
	persons "myapp/internal/persons/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestMovieUsecase_GetMainMovie(t *testing.T) {
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mainMovie := proto.MainMovie{
		ID:          0,
		NamePicture: "name_picture.webp",
		Tagline:     "This is test movie",
		Picture:     "movie_picture.webp",
	}
	movieFromStorage := proto.MainMovie{
		ID:          mainMovie.ID,
		NamePicture: mainMovie.NamePicture,
		Tagline:     mainMovie.Tagline,
		Picture:     mainMovie.Picture,
	}
	input := proto.GetMainMovieOptions{}

	mainMovie.Picture, _ = images.GenerateFileURL(mainMovie.Picture, "posters")
	mainMovie.NamePicture, _ = images.GenerateFileURL(mainMovie.NamePicture, "logos")

	tests := []struct {
		name          string
		storageMock   *mock.MockMovieStorage
		input         proto.GetMainMovieOptions
		expected      proto.MainMovie
		expectedError bool
	}{
		{
			name: "Get main movie",
			storageMock: &mock.MockMovieStorage{
				GetRandomMovieFunc: func() (*proto.MainMovie, error) {
					return &movieFromStorage, nil
				},
			},
			input:         input,
			expected:      mainMovie,
			expectedError: false,
		},
		{
			name: "Return error",
			storageMock: &mock.MockMovieStorage{
				GetRandomMovieFunc: func() (*proto.MainMovie, error) {
					return nil, errors.New(testError)
				},
			},
			input:         input,
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.storageMock, nil, nil, nil)
			mainMovie, err := r.GetMainMovie(context.Background(), &test.input)

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
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	person1 := proto.PersonInMovie{
		ID:       1,
		Name:     "Актер 1",
		Photo:    "photo.webp",
		Position: "Актер",
	}
	person2 := proto.PersonInMovie{
		ID:       2,
		Name:     "Актер 2",
		Photo:    "photo.webp",
		Position: "Актер",
	}
	movieStruct := proto.Movie{
		ID:              1,
		Name:            "Movie1",
		IsMovie:         true,
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
		Genre: []*proto.Genres{
			{
				ID:   1,
				Name: "Боевик",
			},
			{
				ID:   2,
				Name: "Детектив",
			},
		},
		Staff: []*proto.PersonInMovie{&person1, &person2},
	}
	season := proto.Season{
		ID:     1,
		Number: 1,
		Episodes: []*proto.Episode{
			{
				ID:          1,
				Name:        "Episode 1",
				Number:      1,
				Description: "Test episode",
				Video:       "episodeVideo.mp4",
				Picture:     "episodePicture.mp4",
			},
		},
	}
	series := proto.Movie{
		ID:              1,
		Name:            "Movie1",
		IsMovie:         false,
		NamePicture:     "name_picture.webp",
		Year:            2000,
		Duration:        "1 час 52 минуты",
		AgeLimit:        12,
		Description:     "Test film",
		KinopoiskRating: 9.1,
		Rating:          8.1,
		Tagline:         "Tagline of test film",
		Picture:         "picture.webp",
		Trailer:         "trailer.webp",
		Video:           "testVideo.mp4",
		Country:         []string{"Россия", "Франция"},
		Seasons: []*proto.Season{
			&season,
		},
		Genre: []*proto.Genres{
			{
				ID:   1,
				Name: "Боевик",
			},
			{
				ID:   2,
				Name: "Детектив",
			},
		},
		Staff: []*proto.PersonInMovie{&person1, &person2},
	}

	movieFromStorage := proto.Movie{
		ID:              movieStruct.ID,
		Name:            movieStruct.Name,
		IsMovie:         true,
		NamePicture:     movieStruct.NamePicture,
		Year:            movieStruct.Year,
		Duration:        movieStruct.Duration,
		AgeLimit:        movieStruct.AgeLimit,
		Description:     movieStruct.Description,
		KinopoiskRating: movieStruct.KinopoiskRating,
		Rating:          movieStruct.Rating,
		Tagline:         movieStruct.Tagline,
		Picture:         movieStruct.Picture,
		Video:           movieStruct.Video,
		Trailer:         movieStruct.Trailer,
		Country:         nil,
		Genre:           nil,
		Staff:           nil,
	}
	seriesFromStorage := proto.Movie{
		ID:              series.ID,
		Name:            series.Name,
		IsMovie:         false,
		NamePicture:     series.NamePicture,
		Year:            series.Year,
		Duration:        series.Duration,
		AgeLimit:        series.AgeLimit,
		Description:     series.Description,
		KinopoiskRating: series.KinopoiskRating,
		Rating:          series.Rating,
		Tagline:         series.Tagline,
		Picture:         series.Picture,
		Trailer:         series.Trailer,
		Video:           series.Video,
		Country:         nil,
		Genre:           nil,
		Staff:           nil,
	}

	movieStruct.Picture, _ = images.GenerateFileURL(movieStruct.Picture, "posters")
	movieStruct.Video, _ = images.GenerateFileURL(movieStruct.Video, "movie")
	movieStruct.Trailer, _ = images.GenerateFileURL(movieStruct.Trailer, "trailers")
	movieStruct.NamePicture, _ = images.GenerateFileURL(movieStruct.NamePicture, "logos")

	for i := 0; i < len(series.Seasons); i++ {
		for j := 0; j < len(series.Seasons[i].Episodes); j++ {
			series.Seasons[i].Episodes[j].Picture, _ = images.GenerateFileURL(series.Seasons[i].Episodes[j].Picture, "posters")
			series.Seasons[i].Episodes[j].Video, _ = images.GenerateFileURL(series.Seasons[i].Episodes[j].Video, "series")
		}
	}
	series.Picture, _ = images.GenerateFileURL(series.Picture, "posters")
	series.Video, _ = images.GenerateFileURL(series.Video, "movie")
	series.Trailer, _ = images.GenerateFileURL(series.Trailer, "trailers")
	series.NamePicture, _ = images.GenerateFileURL(series.NamePicture, "logos")

	tests := []struct {
		name               string
		movieStorageMock   *mock.MockMovieStorage
		countryStorageMock *countryRepository.MockCountryStorage
		genreStorageMock   *genre.MockGenreStorage
		personStorageMock  *persons.MockPersonsStorage
		expected           proto.Movie
		expectedError      bool
	}{
		{
			name: "Get one movie",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   1,
							Name: "Боевик",
						},
						{
							ID:   2,
							Name: "Детектив",
						},
					}, nil
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return movieStruct.Staff, nil
				},
			},
			expected:      movieStruct,
			expectedError: false,
		},
		{
			name: "Get one series",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &seriesFromStorage, nil
				},
				GetSeasonsAndEpisodesFunc: func(seriesId int) ([]*proto.Season, error) {
					return []*proto.Season{&season}, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   1,
							Name: "Боевик",
						},
						{
							ID:   2,
							Name: "Детектив",
						},
					}, nil
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return movieStruct.Staff, nil
				},
			},
			expected:      series,
			expectedError: false,
		},
		{
			name: "Movie storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   1,
							Name: "Боевик",
						},
						{
							ID:   2,
							Name: "Детектив",
						},
					}, nil
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return movieStruct.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Country storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   1,
							Name: "Боевик",
						},
						{
							ID:   2,
							Name: "Детектив",
						},
					}, nil
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return movieStruct.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return nil, errors.New(testError)
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return movieStruct.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Staff storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   1,
							Name: "Боевик",
						},
						{
							ID:   2,
							Name: "Детектив",
						},
					}, nil
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.movieStorageMock, test.genreStorageMock, test.countryStorageMock, test.personStorageMock)
			mainMovie, err := r.GetByID(context.Background(), &proto.GetMovieOptions{MovieID: 1})

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
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieStruct := proto.Movie{
		ID:              1,
		Name:            "Movie1",
		NamePicture:     "name_picture.webp",
		IsMovie:         true,
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
		Genre: []*proto.Genres{
			{
				ID:   1,
				Name: "Драма",
			},
			{
				ID:   2,
				Name: "Комедия",
			},
		},
		Staff: nil,
	}

	movieFromStorage := proto.Movie{
		ID:              movieStruct.ID,
		Name:            movieStruct.Name,
		NamePicture:     movieStruct.NamePicture,
		IsMovie:         movieStruct.IsMovie,
		Year:            movieStruct.Year,
		Duration:        movieStruct.Duration,
		AgeLimit:        movieStruct.AgeLimit,
		Description:     movieStruct.Description,
		KinopoiskRating: movieStruct.KinopoiskRating,
		Rating:          movieStruct.Rating,
		Tagline:         movieStruct.Tagline,
		Picture:         movieStruct.Picture,
		Video:           movieStruct.Video,
		Trailer:         movieStruct.Trailer,
		Country:         nil,
		Genre:           nil,
		Staff:           nil,
	}

	movieStruct.Picture, _ = images.GenerateFileURL(movieStruct.Picture, "posters")
	movieStruct.Video, _ = images.GenerateFileURL(movieStruct.Video, "movie")
	movieStruct.Trailer, _ = images.GenerateFileURL(movieStruct.Trailer, "trailers")
	movieStruct.NamePicture, _ = images.GenerateFileURL(movieStruct.NamePicture, "logos")

	tests := []struct {
		name               string
		movieStorageMock   *mock.MockMovieStorage
		countryStorageMock *countryRepository.MockCountryStorage
		genreStorageMock   *genre.MockGenreStorage
		expected           []*proto.Movie
		expectedError      bool
	}{
		{
			name: "Get one movie",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return []*proto.Movie{&movieFromStorage}, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   int(movieStruct.Genre[0].ID),
							Name: movieStruct.Genre[0].Name,
						},
						{
							ID:   int(movieStruct.Genre[1].ID),
							Name: movieStruct.Genre[1].Name,
						},
					}, nil
				},
			},
			expected:      []*proto.Movie{&movieStruct},
			expectedError: false,
		},
		{
			name: "Movie storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   int(movieStruct.Genre[0].ID),
							Name: movieStruct.Genre[0].Name,
						},
						{
							ID:   int(movieStruct.Genre[1].ID),
							Name: movieStruct.Genre[1].Name,
						},
					}, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Country storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return []*proto.Movie{&movieFromStorage}, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   int(movieStruct.Genre[0].ID),
							Name: movieStruct.Genre[0].Name,
						},
						{
							ID:   int(movieStruct.Genre[1].ID),
							Name: movieStruct.Genre[1].Name,
						},
					}, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetMovieRatingFunc: func(movieID int) (*movie.GetMovieRatingAnswer, error) {
					return &movie.GetMovieRatingAnswer{
						RatingSum:   int(movieStruct.Rating * 10),
						RatingCount: 10,
					}, nil
				},
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return []*proto.Movie{&movieFromStorage}, nil
				},
			},
			countryStorageMock: &countryRepository.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movieStruct.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.movieStorageMock, test.genreStorageMock, test.countryStorageMock, nil)
			mainMovie, err := r.GetRandom(context.Background(), &proto.GetRandomOptions{
				Limit:  1,
				Offset: 0,
			})

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				for i := 0; i < len(test.expected); i++ {
					assert.Equal(t, *test.expected[i], *mainMovie.Movie[i])
				}
			}
		})
	}
}
