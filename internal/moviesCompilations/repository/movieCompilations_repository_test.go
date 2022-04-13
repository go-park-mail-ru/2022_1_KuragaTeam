package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/internal/moviesCompilations"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMovieRepository_GetByGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		genreID     int
		expected    moviesCompilations.MovieCompilation
		expectedErr error
	}{
		{
			name:    "Get top MC",
			genreID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Боевик")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM genre WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				rows = sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
					"m_g.movie_id = m.id WHERE m_g.genre_id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			expected: moviesCompilations.MovieCompilation{
				Name: "Боевик",
				Movies: []moviesCompilations.Movie{
					{
						ID:      1,
						Name:    "Movie1",
						Genre:   nil,
						Picture: "picture.webp",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:    "Error occurred during SELECT request",
			genreID: 1,
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM genre WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
				rows := sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
					"m_g.movie_id = m.id WHERE m_g.genre_id=$1")).WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name:    "Error occurred during SELECT request",
			genreID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Боевик")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM genre WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
					"m_g.movie_id = m.id WHERE m_g.genre_id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			MC, err := storage.GetByGenre(th.genreID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *MC)
			}
		})
	}
}

func TestMovieRepository_GetByCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		countryID   int
		expected    moviesCompilations.MovieCompilation
		expectedErr error
	}{
		{
			name:      "Get by country",
			countryID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Россия")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM country WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				rows = sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_countries m_c ON " +
					"m_c.movie_id = m.id WHERE m_c.country_id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			expected: moviesCompilations.MovieCompilation{
				Name: "Россия",
				Movies: []moviesCompilations.Movie{
					{
						ID:      1,
						Name:    "Movie1",
						Genre:   nil,
						Picture: "picture.webp",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:      "Error occurred during SELECT request",
			countryID: 1,
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM country WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
				rows := sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_countries m_c ON " +
					"m_c.movie_id = m.id WHERE m_c.country_id=$1")).WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name:      "Error occurred during SELECT request",
			countryID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Россия")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM country WHERE id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_countries m_c ON " +
					"m_c.movie_id = m.id WHERE m_c.country_id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			MC, err := storage.GetByCountry(th.countryID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *MC)
			}
		})
	}
}

func TestMovieRepository_GetByMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		movieID     int
		expected    moviesCompilations.MovieCompilation
		expectedErr error
	}{
		{
			name:    "Get by movie",
			movieID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("2", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT DISTINCT m.id, m.name, m.picture FROM movies AS m " +
					"JOIN movies_genre m_g ON m_g.movie_id = m.id " +
					"JOIN movies_genre m_g2 ON m_g2.genre_id = m_g.genre_id " +
					"WHERE m_g2.movie_id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			expected: moviesCompilations.MovieCompilation{
				Name: "Похожие по жанру",
				Movies: []moviesCompilations.Movie{
					{
						ID:      2,
						Name:    "Movie1",
						Genre:   nil,
						Picture: "picture.webp",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:    "Error occurred during SELECT request",
			movieID: 1,
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT DISTINCT m.id, m.name, m.picture FROM movies AS m " +
					"JOIN movies_genre m_g ON m_g.movie_id = m.id " +
					"JOIN movies_genre m_g2 ON m_g2.genre_id = m_g.genre_id " +
					"WHERE m_g2.movie_id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			MC, err := storage.GetByMovie(th.movieID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *MC)
			}
		})
	}
}

func TestMovieRepository_GetByPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		personID    int
		expected    moviesCompilations.MovieCompilation
		expectedErr error
	}{
		{
			name:     "Get by person",
			personID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_staff m_s ON " +
					"m_s.movie_id = m.id WHERE m_s.person_id=$1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			expected: moviesCompilations.MovieCompilation{
				Name: "Фильмография",
				Movies: []moviesCompilations.Movie{
					{
						ID:      1,
						Name:    "Movie1",
						Genre:   nil,
						Picture: "picture.webp",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Error occurred during SELECT request",
			personID: 1,
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_staff m_s ON " +
					"m_s.movie_id = m.id WHERE m_s.person_id=$1")).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			MC, err := storage.GetByPerson(th.personID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *MC)
			}
		})
	}
}

func TestMovieRepository_GetTop(t *testing.T) {
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
		expected    moviesCompilations.MovieCompilation
		expectedErr error
	}{
		{
			name:  "Get top",
			limit: 10,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, picture FROM movies ORDER BY kinopoisk_rating DESC LIMIT $1")).
					WithArgs(driver.Value(10)).WillReturnRows(rows)
			},
			expected: moviesCompilations.MovieCompilation{
				Name: "Топ рейтинга",
				Movies: []moviesCompilations.Movie{
					{
						ID:      1,
						Name:    "Movie1",
						Genre:   nil,
						Picture: "picture.webp",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:  "Error occurred during SELECT request",
			limit: 1,
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, picture FROM movies ORDER BY kinopoisk_rating DESC LIMIT $1")).
					WithArgs(driver.Value(10)).WillReturnError(errors.New("Error occurred during request "))
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			MC, err := storage.GetTop(th.limit)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *MC)
			}
		})
	}
}

func TestMovieRepository_GetTopByYear(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		year        int
		expected    moviesCompilations.MovieCompilation
		expectedErr error
	}{
		{
			name: "Get top by year",
			year: 2011,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "picture"}).
					AddRow("1", "Movie1", "picture.webp")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, picture FROM movies WHERE year=$1 ORDER BY kinopoisk_rating DESC")).
					WithArgs(driver.Value(2011)).WillReturnRows(rows)
			},
			expected: moviesCompilations.MovieCompilation{
				Name: "Лучшее за 2011 год",
				Movies: []moviesCompilations.Movie{
					{
						ID:      1,
						Name:    "Movie1",
						Genre:   nil,
						Picture: "picture.webp",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			year: 2011,
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, picture FROM movies WHERE year=$1 ORDER BY kinopoisk_rating DESC")).
					WithArgs(driver.Value(2011)).WillReturnError(errors.New("Error occurred during request "))
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			MC, err := storage.GetTopByYear(th.year)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *MC)
			}
		})
	}
}
