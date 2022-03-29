package genre

type Storage interface {
	GetByMovieID(id int) ([]string, error)
	//GetAllMovies() (string, error)
}
