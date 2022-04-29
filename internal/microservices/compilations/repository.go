package compilations

import "myapp/internal/microservices/compilations/proto"

type Storage interface {
	GetAllMovies(limit, offset int, isMovie bool) (*proto.MovieCompilation, error)
	GetByGenre(genreID int) (*proto.MovieCompilation, error)
	GetByCountry(countryID int) (*proto.MovieCompilation, error)
	GetByMovie(movieID int) (*proto.MovieCompilation, error)
	GetByPerson(personID int) (*proto.MovieCompilation, error)
	GetTop(limit int) (*proto.MovieCompilation, error)
	GetTopByYear(year int) (*proto.MovieCompilation, error)
	GetFavorites(data *proto.GetFavoritesOptions) (*proto.MovieCompilation, error)
	FindMovie(text string, isMovie bool) (*proto.MovieCompilation, error)
}
