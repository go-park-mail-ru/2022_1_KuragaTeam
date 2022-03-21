package moviesCompilations

type Service interface {
	GetMainCompilations() ([]MovieCompilation, error)
	GetByGenre(genreID int) (MovieCompilation, error)
	GetByMovie(movieID int) (MovieCompilation, error)
	//GetPopularID(context echo.Context, limit int) (MovieCompilation, error)
}
