package moviesCompilations

type Storage interface {
	GetByGenre(genreID int) (MovieCompilation, error)
	GetByCountry(countryID int) (MovieCompilation, error)
	GetByMovie(movieID int) (MovieCompilation, error)
	GetByPerson(personID int) (MovieCompilation, error)
	GetTop(limit int) (MovieCompilation, error)
	GetTopByYear(year int) (MovieCompilation, error)
}
