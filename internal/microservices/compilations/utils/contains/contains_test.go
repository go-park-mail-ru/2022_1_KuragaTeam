package contains

import (
	"myapp/internal/microservices/compilations/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieContains(t *testing.T) {
	movies := make([]*proto.MovieInfo, 0)
	movies = append(movies, &proto.MovieInfo{
		ID:      1,
		Name:    "1",
		Genre:   nil,
		Picture: "1",
	}, &proto.MovieInfo{
		ID:      2,
		Name:    "2",
		Genre:   nil,
		Picture: "2",
	})
	tests := []struct {
		name   string
		movies []*proto.MovieInfo
		id     int64
		result bool
	}{
		{
			name:   "Movie contains",
			movies: movies,
			id:     2,
			result: true,
		},
		{
			name:   "Movie doesn't contains",
			movies: movies,
			id:     3,
			result: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			result := MovieContains(th.movies, th.id)

			assert.Equal(t, th.result, result)
		})
	}
}

func TestPersonContains(t *testing.T) {
	persons := make([]*proto.PersonInfo, 0)
	persons = append(persons, &proto.PersonInfo{
		ID:       1,
		Name:     "1",
		Photo:    "1",
		Position: nil,
	}, &proto.PersonInfo{
		ID:       2,
		Name:     "2",
		Photo:    "2",
		Position: nil,
	})
	tests := []struct {
		name    string
		persons []*proto.PersonInfo
		id      int64
		result  bool
	}{
		{
			name:    "Person contains",
			persons: persons,
			id:      2,
			result:  true,
		},
		{
			name:    "Person doesn't contains",
			persons: persons,
			id:      3,
			result:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			result := PersonContains(th.persons, th.id)

			assert.Equal(t, th.result, result)
		})
	}
}
