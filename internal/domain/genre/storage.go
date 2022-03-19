package genre

type Storage interface {
	GetByMovieID(id int) ([]string, error)
	//GetRandomMovies() (string, error)
}
