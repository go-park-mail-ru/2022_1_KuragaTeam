package genre

type Storage interface {
	GetByMovieID(id int) ([]string, error)
	//GetRandom() (string, error)
}
