package compilations

import "myapp/internal/microservices/compilations/proto"

type Storage interface {
	GetAllMovies(limit, offset int, isMovie bool) (*proto.MovieCompilation, error)
	GetByGenre(limit, offset int, genreID int, random bool) (*proto.MovieCompilation, error)
	GetByCountry(countryID int, random bool) (*proto.MovieCompilation, error)
	GetByMovie(movieID int) (*proto.MovieCompilation, error)
	GetByPerson(personID int) (*proto.MovieCompilation, error)
	GetTop(limit int) (*proto.MovieCompilation, error)
	GetTopByYear(year int) (*proto.MovieCompilation, error)
	GetFavorites(data *proto.GetFavoritesOptions) (*proto.MovieCompilationsArr, error)
	GetFavoritesFilms(data *proto.GetFavoritesOptions) (*proto.MovieCompilation, error)
	GetFavoritesSeries(data *proto.GetFavoritesOptions) (*proto.MovieCompilation, error)
	FindMovie(text string, isMovie bool) (*proto.MovieCompilation, error)
	FindMovieByPartial(text string, isMovie bool) (*proto.MovieCompilation, error)
}
