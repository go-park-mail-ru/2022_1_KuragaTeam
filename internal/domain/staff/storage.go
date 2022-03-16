package staff

import "myapp/internal/domain"

type Storage interface {
	GetByMovieID(id int) ([]domain.Person, error)
	//GetRandomMovies() (string, error)
}
