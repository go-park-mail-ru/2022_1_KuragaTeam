package usecase

import (
	"golang.org/x/net/context"
	"myapp/internal"
	"myapp/internal/country"
	"myapp/internal/genre"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/movie"
	"myapp/internal/persons"
	"myapp/internal/utils/images"

	mapper "github.com/stroiman/go-automapper"
)

type service struct {
	proto.UnimplementedMoviesServer

	movieStorage   movie.Storage
	genreStorage   genre.Storage
	countryStorage country.Storage
	staffStorage   persons.Storage
}

func NewService(movieStorage movie.Storage, genreStorage genre.Storage,
	countryStorage country.Storage, staffStorage persons.Storage) *service {
	return &service{movieStorage: movieStorage, genreStorage: genreStorage,
		countryStorage: countryStorage, staffStorage: staffStorage}
}

func (s *service) concatURLs(movie *internal.Movie) error {
	var err error
	movie.Picture, err = images.GenerateFileURL(movie.Picture, "posters")
	if err != nil {
		return err
	}

	movie.Video, err = images.GenerateFileURL(movie.Video, "movie")
	if err != nil {
		return err
	}

	movie.Trailer, err = images.GenerateFileURL(movie.Trailer, "trailers")
	if err != nil {
		return err
	}

	movie.NamePicture, err = images.GenerateFileURL(movie.NamePicture, "logos")
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, in *proto.GetMovieOptions) (*proto.Movie, error) {
	selectedMovie, err := s.movieStorage.GetOne(int(in.MovieID))
	if err != nil {
		return nil, err
	}
	selectedMovie.Genre, err = s.genreStorage.GetByMovieID(selectedMovie.ID)
	if err != nil {
		return nil, err
	}

	selectedMovie.Country, err = s.countryStorage.GetByMovieID(selectedMovie.ID)
	if err != nil {
		return nil, err
	}

	selectedMovie.Staff, err = s.staffStorage.GetByMovieID(selectedMovie.ID)
	if err != nil {
		return nil, err
	}

	for i, _ := range selectedMovie.Staff {
		selectedMovie.Staff[i].Photo, err = images.GenerateFileURL(selectedMovie.Staff[i].Photo, "persons")
		if err != nil {
			return nil, err
		}
	}

	selectedMovie.Rating = 8.1 // пока что просто замокано

	err = s.concatURLs(selectedMovie)
	if err != nil {
		return nil, err
	}
	newStruct := proto.Movie{
		ID:              int64(selectedMovie.ID),
		Name:            selectedMovie.Name,
		NamePicture:     selectedMovie.NamePicture,
		Year:            int32(selectedMovie.Year),
		Duration:        selectedMovie.Duration,
		AgeLimit:        int32(selectedMovie.AgeLimit),
		Description:     selectedMovie.Description,
		KinopoiskRating: selectedMovie.KinopoiskRating,
		Rating:          selectedMovie.Rating,
		Tagline:         selectedMovie.Tagline,
		Picture:         selectedMovie.Picture,
		Video:           selectedMovie.Video,
		Trailer:         selectedMovie.Trailer,
		Country:         nil,
		Genre:           nil,
		Staff:           nil,
	}

	//mapper.Map(selectedMovie, &newStruct)
	return &newStruct, nil
}

func (s *service) GetRandom(ctx context.Context, in *proto.GetRandomOptions) (*proto.MoviesArr, error) {
	movies, err := s.movieStorage.GetAllMovies(int(in.Limit), int(in.Offset))
	for i := 0; i < len(movies); i++ {
		movies[i].Genre, err = s.genreStorage.GetByMovieID(movies[i].ID)
		if err != nil {
			return nil, err
		}
		movies[i].Country, err = s.countryStorage.GetByMovieID(movies[i].ID)
		if err != nil {
			return nil, err
		}

		movies[i].Rating = 8.1 // пока что просто замокано

		err = s.concatURLs(&movies[i])
		if err != nil {
			return nil, err
		}

	}
	//return movies, err
	return nil, nil
}

func (s *service) GetMainMovie(ctx context.Context, in *proto.GetMainMovieOptions) (*proto.MainMovie, error) {
	selectedMovie, err := s.movieStorage.GetRandomMovie()
	if err != nil {
		return nil, err
	}
	selectedMovie.Picture, err = images.GenerateFileURL(selectedMovie.Picture, "posters")
	if err != nil {
		return nil, err
	}
	selectedMovie.NamePicture, err = images.GenerateFileURL(selectedMovie.NamePicture, "logos")
	if err != nil {
		return nil, err
	}

	newStruct := proto.MainMovie{}

	mapper.Map(selectedMovie, &newStruct)
	return &newStruct, nil
}
