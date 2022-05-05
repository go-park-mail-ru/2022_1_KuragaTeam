package position

type Storage interface {
	GetByPersonID(id int) ([]string, error)
	GetPersonPosByMovieID(personID, movieId int) (string, error)
	//GetAllMovies() (string, error)
}
