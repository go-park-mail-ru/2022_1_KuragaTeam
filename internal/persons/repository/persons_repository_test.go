package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
	"myapp/internal/microservices/movie/proto"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPersonRepository_GetOne(t *testing.T) {
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
		expected    internal.Person
		expectedErr error
	}{
		{
			name: "Get one person",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "photo", "addit_photo1", "addit_photo2", "description"}).
					AddRow("1", "Person1", "photo.webp", "addit_photo1.webp", "addit_photo2.webp", "Description")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT p.id, p.name, p.photo, p.addit_photo1, p.addit_photo2, " +
					"p.description FROM person AS p WHERE p.id = $1")).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id: 1,
			expected: internal.Person{
				ID:          1,
				Name:        "Person1",
				Photo:       "photo.webp",
				AdditPhoto1: "addit_photo1.webp",
				AdditPhoto2: "addit_photo2.webp",
				Description: "Description",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT p.id, p.name, p.photo, p.addit_photo1, p.addit_photo2, " +
					"p.description FROM person AS p WHERE p.id = $1")).
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

			movie, err := storage.GetByPersonID(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, *movie)
			}
		})
	}
}

func TestPersonRepository_GetByMovieID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	person1 := proto.PersonInMovie{
		ID:       1,
		Name:     "Person1",
		Photo:    "photo.webp",
		Position: "Position",
	}

	person2 := proto.PersonInMovie{
		ID:       2,
		Name:     "Person2",
		Photo:    "photo.webp",
		Position: "Position",
	}

	tests := []struct {
		name        string
		mock        func()
		id          int
		expected    []*proto.PersonInMovie
		expectedErr error
	}{
		{
			name: "Get by movie id",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "photo", "picture"}).
					AddRow(person1.ID, person1.Name, person1.Photo, person1.Position)
				rows.AddRow(person2.ID, person2.Name, person2.Photo, person2.Position)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT p.id, p.name, p.photo, pos.name FROM person AS p JOIN movies_staff mv_s ON mv_s.person_id = p.id " +
					"JOIN position pos ON mv_s.position_id = pos.id " +
					"WHERE mv_s.movie_id = $1 ORDER BY pos.id")).WillReturnRows(rows)
			},
			id:          1,
			expected:    []*proto.PersonInMovie{&person1, &person2},
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

			persons, err := storage.GetByMovieID(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, persons)
			}
		})
	}
}
