package contains

import "myapp/internal/microservices/compilations/proto"

func MovieContains(movies []*proto.MovieInfo, movieID int64) bool {
	for _, currentMovie := range movies {
		if currentMovie.ID == movieID {
			return true
		}
	}
	return false
}

func PersonContains(persons []*proto.PersonInfo, personID int64) bool {
	for _, currentPerson := range persons {
		if currentPerson.ID == personID {
			return true
		}
	}
	return false
}
