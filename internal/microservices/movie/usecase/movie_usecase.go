package usecase

import (
	"myapp/internal/country"
	"myapp/internal/genre"
	"myapp/internal/microservices/movie"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/microservices/movie/utils/images"
	"myapp/internal/persons"

	"golang.org/x/net/context"
)

type Service struct {
	proto.UnimplementedMoviesServer

	movieStorage   movie.Storage
	genreStorage   genre.Storage
	countryStorage country.Storage
	staffStorage   persons.Storage
}

func NewService(movieStorage movie.Storage, genreStorage genre.Storage,
	countryStorage country.Storage, staffStorage persons.Storage) *Service {
	return &Service{movieStorage: movieStorage, genreStorage: genreStorage,
		countryStorage: countryStorage, staffStorage: staffStorage}
}

func (s *Service) fillGenres(movie *proto.Movie) error {
	nextGenres, err := s.genreStorage.GetByMovieID(int(movie.ID))
	if err != nil {
		return err
	}
	for _, nextGenre := range nextGenres {
		movie.Genre = append(movie.Genre, &proto.Genres{
			ID:   int32(nextGenre.ID),
			Name: nextGenre.Name,
		})
	}
	return nil
}

func (s *Service) concatURLs(movie *proto.Movie) error {
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

	if !movie.IsMovie {
		for _, season := range movie.Seasons {
			for _, episode := range (*season).Episodes {
				episode.Picture, err = images.GenerateFileURL(episode.Picture, "posters")
				episode.Video, err = images.GenerateFileURL(episode.Video, "series")
			}
		}
	}

	return nil
}

func (s *Service) GetByID(ctx context.Context, in *proto.GetMovieOptions) (*proto.Movie, error) {
	selectedMovie, err := s.movieStorage.GetOne(int(in.MovieID))
	if err != nil {
		return nil, err
	}

	if !selectedMovie.IsMovie {
		selectedMovie.Seasons, err = s.movieStorage.GetSeasonsAndEpisodes(int(selectedMovie.ID))
		if err != nil {
			return nil, err
		}
	}

	err = s.fillGenres(selectedMovie)
	if err != nil {
		return nil, err
	}

	selectedMovie.Country, err = s.countryStorage.GetByMovieID(int(selectedMovie.ID))
	if err != nil {
		return nil, err
	}

	selectedMovie.Staff, err = s.staffStorage.GetByMovieID(int(selectedMovie.ID))
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(selectedMovie.Staff); i++ {
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
	return selectedMovie, nil
}

func (s *Service) GetRandom(ctx context.Context, in *proto.GetRandomOptions) (*proto.MoviesArr, error) {
	movies, err := s.movieStorage.GetAllMovies(int(in.Limit), int(in.Offset))
	for i := 0; i < len(movies); i++ {
		err = s.fillGenres(movies[i])
		if err != nil {
			return nil, err
		}
		movies[i].Country, err = s.countryStorage.GetByMovieID(int(movies[i].ID))
		if err != nil {
			return nil, err
		}

		movies[i].Rating = 8.1 // пока что просто замокано

		err = s.concatURLs(movies[i])
		if err != nil {
			return nil, err
		}

	}
	return &proto.MoviesArr{Movie: movies}, err
}

func (s *Service) GetMainMovie(ctx context.Context, in *proto.GetMainMovieOptions) (*proto.MainMovie, error) {
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

	return selectedMovie, nil
}

func (s *Service) AddMovieRating(context.Context, *proto.AddRatingOptions) (*proto.NewMovieRating, error) {
	return nil, nil
}
