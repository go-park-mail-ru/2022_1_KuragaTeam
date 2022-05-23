package position

type Storage interface {
	GetByPersonID(id int) ([]string, error)
	GetPersonPosByMovieID(personID, movieID int) (string, error)
}
