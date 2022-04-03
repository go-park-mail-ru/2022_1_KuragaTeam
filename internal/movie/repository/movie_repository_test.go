package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
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
		expected    *internal.Movie
		expectedErr error
	}{
		{
			name: "Get one movie",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "name_picture", "year", "duration", "age_limit",
					"description", "kinopoisk_rating", "tagline", "picture", "video", "trailer"}).
					AddRow("1", "Movie1", "", "0", "", "0", "", "0", "", "", "", "")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, name_picture, year, duration, age_limit, " +
					"description, kinopoisk_rating, tagline, picture, video, trailer FROM movies WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id: 1,
			expected: &internal.Movie{
				ID:              1,
				Name:            "Movie1",
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
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, name_picture, year, duration, age_limit, " +
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

			user, err := storage.GetOne(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, user)
			}
		})
	}
}
