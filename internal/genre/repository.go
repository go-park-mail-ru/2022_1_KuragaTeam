package genre

import "myapp/internal"

type Storage interface {
	GetByMovieID(id int) ([]internal.Genre, error)
	//GetAllMovies() (string, error)
}
