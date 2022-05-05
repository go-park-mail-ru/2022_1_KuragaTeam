package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"myapp/internal"
	country "myapp/internal/country/repository"
	genre "myapp/internal/genre/repository"
	"myapp/internal/microservices/movie/proto"
	mock "myapp/internal/microservices/movie/repository"
	"myapp/internal/microservices/movie/utils/images"
	persons "myapp/internal/persons/repository"
	"testing"
)

func TestMovieUsecase_GetMainMovie(t *testing.T) {
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := proto.MainMovie{
		ID:          0,
		NamePicture: "name_picture.webp",
		Tagline:     "This is test movie",
		Picture:     "movie_picture.webp",
	}
	movieFromStorage := proto.MainMovie{
		ID:          movie.ID,
		NamePicture: movie.NamePicture,
		Tagline:     movie.Tagline,
		Picture:     movie.Picture,
	}
	input := proto.GetMainMovieOptions{}

	movie.Picture, _ = images.GenerateFileURL(movie.Picture, "posters")
	movie.NamePicture, _ = images.GenerateFileURL(movie.NamePicture, "logos")

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
			expected:      movie,
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
	movie := proto.Movie{
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
		ID:              movie.ID,
		Name:            movie.Name,
		IsMovie:         true,
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

	movie.Picture, _ = images.GenerateFileURL(movie.Picture, "posters")
	movie.Video, _ = images.GenerateFileURL(movie.Video, "movie")
	movie.Trailer, _ = images.GenerateFileURL(movie.Trailer, "trailers")
	movie.NamePicture, _ = images.GenerateFileURL(movie.NamePicture, "logos")

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
		countryStorageMock *country.MockCountryStorage
		genreStorageMock   *genre.MockGenreStorage
		personStorageMock  *persons.MockPersonsStorage
		expected           proto.Movie
		expectedError      bool
	}{
		{
			name: "Get one movie",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
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
					return movie.Staff, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
		{
			name: "Get one series",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &seriesFromStorage, nil
				},
				GetSeasonsAndEpisodesFunc: func(seriesId int) ([]*proto.Season, error) {
					return []*proto.Season{&season}, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
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
					return movie.Staff, nil
				},
			},
			expected:      series,
			expectedError: false,
		},
		{
			name: "Movie storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
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
					return movie.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Country storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
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
					return movie.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return nil, errors.New(testError)
				},
			},
			personStorageMock: &persons.MockPersonsStorage{
				GetByMovieIDFunc: func(id int) ([]*proto.PersonInMovie, error) {
					return movie.Staff, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Staff storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetOneFunc: func(id int) (*proto.Movie, error) {
					return &movieFromStorage, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
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

	movie := proto.Movie{
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
		ID:              movie.ID,
		Name:            movie.Name,
		NamePicture:     movie.NamePicture,
		IsMovie:         movie.IsMovie,
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
		countryStorageMock *country.MockCountryStorage
		genreStorageMock   *genre.MockGenreStorage
		expected           []*proto.Movie
		expectedError      bool
	}{
		{
			name: "Get one movie",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return []*proto.Movie{&movieFromStorage}, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   int(movie.Genre[0].ID),
							Name: movie.Genre[0].Name,
						},
						{
							ID:   int(movie.Genre[1].ID),
							Name: movie.Genre[1].Name,
						},
					}, nil
				},
			},
			expected:      []*proto.Movie{&movie},
			expectedError: false,
		},
		{
			name: "Movie storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   int(movie.Genre[0].ID),
							Name: movie.Genre[0].Name,
						},
						{
							ID:   int(movie.Genre[1].ID),
							Name: movie.Genre[1].Name,
						},
					}, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Country storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return []*proto.Movie{&movieFromStorage}, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return []internal.Genre{
						{
							ID:   int(movie.Genre[0].ID),
							Name: movie.Genre[0].Name,
						},
						{
							ID:   int(movie.Genre[1].ID),
							Name: movie.Genre[1].Name,
						},
					}, nil
				},
			},
			expectedError: true,
		},
		{
			name: "Genre storage error",
			movieStorageMock: &mock.MockMovieStorage{
				GetAllMoviesFunc: func(limit, offset int) ([]*proto.Movie, error) {
					return []*proto.Movie{&movieFromStorage}, nil
				},
			},
			countryStorageMock: &country.MockCountryStorage{
				GetByMovieIDFunc: func(id int) ([]string, error) {
					return movie.Country, nil
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
