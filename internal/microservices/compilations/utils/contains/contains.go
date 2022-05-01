package contains

import "myapp/internal/microservices/compilations/proto"

func ContainsMovie(movies []*proto.MovieInfo, movieID int64) bool {
	for _, currentMovie := range movies {
		if currentMovie.ID == movieID {
			return true
		}
	}
	return false
}

func ContainsPerson(persons []*proto.PersonInfo, personID int64) bool {
	for _, currentPerson := range persons {
		if currentPerson.ID == personID {
			return true
		}
	}
	return false
}
