package usecase

import (
	"math"
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
				if err != nil {
					return err
				}
				episode.Video, err = images.GenerateFileURL(episode.Video, "series")
				if err != nil {
					return err
				}
			}
		}
		movie.Video, err = images.GenerateFileURL(movie.Video, "series")
		if err != nil {
			return err
		}
	} else {
		movie.Video, err = images.GenerateFileURL(movie.Video, "movie")
		if err != nil {
			return err
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

	ratingValues, err := s.movieStorage.GetMovieRating(int(in.MovieID))
	if err != nil {
		return nil, err
	}
	if ratingValues.RatingSum != 0 {
		selectedMovie.Rating = float32(math.Round(float64(float64(ratingValues.RatingSum)/
			float64(ratingValues.RatingCount))*10) / 10)
	} else {
		selectedMovie.Rating = -1.0
	}

	err = s.concatURLs(selectedMovie)
	if err != nil {
		return nil, err
	}
	return selectedMovie, nil
}

func (s *Service) GetRandom(ctx context.Context, in *proto.GetRandomOptions) (*proto.MoviesArr, error) {
	movies, err := s.movieStorage.GetAllMovies(int(in.Limit), int(in.Offset))
	for movieIndex := 0; movieIndex < len(movies); movieIndex++ {
		err = s.fillGenres(movies[movieIndex])
		if err != nil {
			return nil, err
		}
		movies[movieIndex].Country, err = s.countryStorage.GetByMovieID(int(movies[movieIndex].ID))
		if err != nil {
			return nil, err
		}

		ratingValues, err := s.movieStorage.GetMovieRating(int(movies[movieIndex].ID))
		if err != nil {
			return nil, err
		}
		if ratingValues.RatingSum != 0 {
			movies[movieIndex].Rating = float32(math.Round(float64(float64(ratingValues.RatingSum)/
				float64(ratingValues.RatingCount))*10) / 10)
		} else {
			movies[movieIndex].Rating = -1.0
		}

		err = s.concatURLs(movies[movieIndex])
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

func (s *Service) AddMovieRating(ctx context.Context, options *proto.AddRatingOptions) (*proto.NewMovieRating, error) {
	checkRating, err := s.movieStorage.CheckRatingExists(options)
	if err != nil {
		return nil, err
	}

	if options.Rating == -1 && checkRating.Exists {
		err = s.movieStorage.RemoveMovieRating(options)
		if err != nil {
			return nil, err
		}
	} else {
		if options.Rating > 10 {
			options.Rating = 10
		}
		if options.Rating < 1 {
			options.Rating = 1
		}

		if checkRating.Exists {
			if checkRating.Rating != int(options.Rating) {
				err = s.movieStorage.ChangeMovieRating(options)
				if err != nil {
					return nil, err
				}
			}
		} else {
			err = s.movieStorage.AddMovieRating(options)
			if err != nil {
				return nil, err
			}
		}
	}

	ratingValues, err := s.movieStorage.GetMovieRating(int(options.MovieID))
	if err != nil {
		return nil, err
	}
	if ratingValues.RatingSum != 0 {
		return &proto.NewMovieRating{Rating: float32(math.Round(float64(float64(ratingValues.RatingSum)/
			float64(ratingValues.RatingCount))*10) / 10)}, nil
	}
	return &proto.NewMovieRating{Rating: float32(-1.0)}, nil
}
