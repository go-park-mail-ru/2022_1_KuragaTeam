package moviesCompilations

type Service interface {
	GetMainCompilations() ([]MovieCompilation, error)
	GetByGenre(genreID int) (MovieCompilation, error)
	GetByMovie(movieID int) (MovieCompilation, error)
	GetByPerson(personID int) (MovieCompilation, error)
	GetTopByYear(year int) (MovieCompilation, error)
	GetTop(limit int) (MovieCompilation, error)
	//GetPopularID(context echo.Context, limit int) (MovieCompilation, error)
}
