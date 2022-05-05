package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/internal/microservices/movie/proto"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMovieRepository_GetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		id          int
		expected    proto.Movie
		expectedErr error
	}{
		{
			name: "Get one movie",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "is_movie", "name_picture", "year", "duration", "age_limit",
					"description", "kinopoisk_rating", "tagline", "picture", "video", "trailer"}).
					AddRow("1", "Movie1", "true", "", "0", "", "0", "", "0", "", "", "", "")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_movie, name_picture, year, duration, " +
					"age_limit, description, kinopoisk_rating, tagline, picture, video, trailer FROM movies WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id: 1,
			expected: proto.Movie{
				ID:              1,
				Name:            "Movie1",
				IsMovie:         true,
				NamePicture:     "",
				Year:            0,
				Duration:        "",
				AgeLimit:        0,
				Description:     "",
				KinopoiskRating: 0,
				Rating:          0,
				Tagline:         "",
				Picture:         "",
				Video:           "",
				Trailer:         "",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_movie, name_picture, year, duration, age_limit, " +
					"description, kinopoisk_rating, tagline, picture, video, trailer FROM movies WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			id:          1,
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			movie, err := storage.GetOne(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *movie)
			}
		})
	}
}

func TestMovieRepository_GetRandomMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	movie := proto.MainMovie{
		ID:          1,
		NamePicture: "name_picture.webp",
		Tagline:     "Tagline",
		Picture:     "picture",
	}

	tests := []struct {
		name        string
		mock        func()
		id          int
		expected    proto.MainMovie
		expectedErr error
	}{
		{
			name: "Get one movie",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "tagline", "picture"}).
					AddRow("1", "name_picture.webp", "Tagline", "picture")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name_picture, tagline, picture FROM movies ORDER BY " +
					"RANDOM() LIMIT 1")).WillReturnRows(rows)
			},
			id:          1,
			expected:    movie,
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, tagline, picture FROM movies ORDER " +
					"BY RANDOM() LIMIT 1")).WillReturnError(errors.New("Error occurred during request "))
			},
			id:          1,
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			movie, err := storage.GetRandomMovie()
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *movie)
			}
		})
	}
}

func TestMovieRepository_GetAllMovies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		limit       int
		offset      int
		expected    []*proto.Movie
		expectedErr error
	}{
		{
			name: "Get all movies",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "is_movie", "name_picture", "year", "duration", "age_limit",
					"description", "kinopoisk_rating", "tagline", "picture", "video", "trailer"}).
					AddRow("1", "Movie1", "true", "", "0", "", "0", "", "0", "", "", "", "")
				rows.AddRow("2", "Movie2", "true", "", "0", "", "0", "", "0", "", "", "", "")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_movie, name_picture, year, duration, age_limit, "+
					"description, kinopoisk_rating, tagline, picture, video, trailer FROM movies LIMIT $1 OFFSET $2")).
					WithArgs(driver.Value(10), driver.Value(0)).WillReturnRows(rows)
			},
			limit:  10,
			offset: 0,
			expected: []*proto.Movie{
				{
					ID:              1,
					Name:            "Movie1",
					IsMovie:         true,
					NamePicture:     "",
					Year:            0,
					Duration:        "",
					AgeLimit:        0,
					Description:     "",
					KinopoiskRating: 0,
					Rating:          0,
					Tagline:         "",
					Picture:         "",
					Video:           "",
					Trailer:         "",
				},
				{
					ID:              2,
					Name:            "Movie2",
					NamePicture:     "",
					IsMovie:         true,
					Year:            0,
					Duration:        "",
					AgeLimit:        0,
					Description:     "",
					KinopoiskRating: 0,
					Rating:          0,
					Tagline:         "",
					Picture:         "",
					Video:           "",
					Trailer:         "",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_movie, name_picture, year, duration, age_limit, "+
					"description, kinopoisk_rating, tagline, picture, video, trailer FROM movies LIMIT $1 OFFSET $2")).
					WithArgs(driver.Value(10), driver.Value(0)).WillReturnError(errors.New("Error occurred during request "))
			},
			limit:       10,
			offset:      0,
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			movies, err := storage.GetAllMovies(th.limit, th.offset)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				for i := 0; i < len(th.expected); i++ {
					assert.Equal(t, th.expected[i], movies[i])
				}
			}
		})
	}
}

func TestMovieRepository_GetSeasonsAndEpisodes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		seriesID    int
		expected    []*proto.Season
		expectedErr error
	}{
		{
			name: "Get all seasons",
			mock: func() {
				seasonRows := sqlmock.NewRows([]string{"id", "number"}).
					AddRow("1", "1")
				seasonRows.AddRow("2", "2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, number FROM seasons WHERE movie_id = $1 ORDER BY number;")).
					WithArgs(driver.Value(10)).WillReturnRows(seasonRows)
				episodeRows := sqlmock.NewRows([]string{"id", "name", "number", "description", "video", "photo", "season_id", "season_number"}).
					AddRow("1", "Ep1", "1", "Test EP1", "ep1video.mp4", "ep1photo.webp", "1", "1")
				episodeRows.AddRow("2", "Ep2", "2", "Test EP2", "ep2video.mp4", "ep2photo.webp", "1", "1")
				episodeRows.AddRow("3", "Ep3", "3", "Test EP3", "ep3video.mp4", "ep3photo.webp", "2", "2")
				episodeRows.AddRow("4", "Ep4", "4", "Test EP4", "ep4video.mp4", "ep4photo.webp", "2", "2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT e.id, e.name, e.number, e.description, e.video, e.photo, " +
					"s.id, s.number FROM episode AS e JOIN seasons s on e.season_id = s.id WHERE s.movie_id = $1 " +
					"ORDER BY s.number, e.number;")).
					WithArgs(driver.Value(10)).WillReturnRows(episodeRows)
			},
			seriesID: 10,
			expected: []*proto.Season{
				{
					ID:     1,
					Number: 1,
					Episodes: []*proto.Episode{
						{
							ID:          1,
							Name:        "Ep1",
							Number:      1,
							Description: "Test EP1",
							Video:       "ep1video.mp4",
							Picture:     "ep1photo.webp",
						},
						{
							ID:          2,
							Name:        "Ep2",
							Number:      2,
							Description: "Test EP2",
							Video:       "ep2video.mp4",
							Picture:     "ep2photo.webp",
						},
					},
				},
				{
					ID:     2,
					Number: 2,
					Episodes: []*proto.Episode{
						{
							ID:          3,
							Name:        "Ep3",
							Number:      3,
							Description: "Test EP3",
							Video:       "ep3video.mp4",
							Picture:     "ep3photo.webp",
						},
						{
							ID:          4,
							Name:        "Ep4",
							Number:      4,
							Description: "Test EP4",
							Video:       "ep4video.mp4",
							Picture:     "ep4photo.webp",
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, number FROM seasons WHERE movie_id = $1 ORDER BY number;")).
					WithArgs(driver.Value(10)).WillReturnError(errors.New("Error occurred during request "))
			},
			seriesID:    10,
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				seasonRows := sqlmock.NewRows([]string{"id", "number", "movie_id"}).
					AddRow("1", "1", "10")
				seasonRows.AddRow("2", "2", "10")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, number FROM seasons WHERE movie_id = $1 ORDER BY number;")).
					WithArgs(driver.Value(10)).WillReturnRows(seasonRows)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT e.id, e.name, e.number, e.description, e.video, e.photo, " +
					"s.id, s.number FROM episode AS e JOIN seasons s on e.season_id = s.id WHERE s.movie_id = $1 ORDER " +
					"BY s.number, e.number;")).
					WithArgs(driver.Value(10)).WillReturnError(errors.New("Error occurred during request "))
			},
			seriesID:    10,
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			seasons, err := storage.GetSeasonsAndEpisodes(th.seriesID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				for i := 0; i < len(th.expected); i++ {
					for j := 0; j < len(th.expected); j++ {
						assert.Equal(t, th.expected[i].ID, seasons[i].ID)
						assert.Equal(t, th.expected[i].Number, seasons[i].Number)
						assert.Equal(t, th.expected[i].Episodes[j], seasons[i].Episodes[j])
					}
				}
			}
		})
	}
}
