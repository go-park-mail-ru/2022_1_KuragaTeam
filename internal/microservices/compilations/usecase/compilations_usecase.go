package usecase

import (
	"myapp/internal/constants"
	"myapp/internal/genre"
	"myapp/internal/microservices/compilations"
	"myapp/internal/microservices/compilations/proto"
	"myapp/internal/microservices/compilations/utils/contains"
	"myapp/internal/microservices/compilations/utils/images"
	"myapp/internal/persons"
	"strings"

	"golang.org/x/net/context"
)

type Service struct {
	MCStorage    compilations.Storage
	genreStorage genre.Storage
	staffStorage persons.Storage
}

func NewService(MCStorage compilations.Storage, genreStorage genre.Storage, staffStorage persons.Storage) *Service {
	return &Service{MCStorage: MCStorage, genreStorage: genreStorage, staffStorage: staffStorage}
}

func (s *Service) fillGenres(compilation *proto.MovieCompilation) error {
	for mcIndex := 0; mcIndex < len(compilation.Movies); mcIndex++ {
		nextGenres, err := s.genreStorage.GetByMovieID(int(compilation.Movies[mcIndex].ID))
		if err != nil {
			return err
		}
		for _, nextGenre := range nextGenres {
			compilation.Movies[mcIndex].Genre = append(compilation.Movies[mcIndex].Genre, &proto.Genre{
				ID:   int32(nextGenre.ID),
				Name: nextGenre.Name,
			})
		}
	}
	return nil
}

func (s *Service) concatUrls(compilation *proto.MovieCompilation) error {
	var err error
	for i := range compilation.Movies {
		compilation.Movies[i].Picture, err = images.GenerateFileURL(compilation.Movies[i].Picture, "posters")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetAllMovies(ctx context.Context, in *proto.GetCompilationOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetAllMovies(int(in.Limit), int(in.Offset), true)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetAllSeries(ctx context.Context, in *proto.GetCompilationOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetAllMovies(int(in.Limit), int(in.Offset), false)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetMainCompilations(ctx context.Context, in *proto.GetMainCompilationsOptions) (*proto.MovieCompilationsArr, error) {

	compilation := make([]*proto.MovieCompilation, 0)

	nextMC, err := s.GetTop(ctx, &proto.GetCompilationOptions{Limit: 12, Offset: 0})
	if err != nil {
		return nil, err
	}
	compilation = append(compilation, nextMC)

	nextMC, err = s.GetTopByYear(ctx, &proto.GetByIDOptions{
		ID:     2011,
		Limit:  0,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}
	compilation = append(compilation, nextMC)

	nextMC, err = s.GetByGenre(ctx, &proto.GetByIDOptions{
		ID:     2,
		Limit:  0,
		Offset: 0,
	}) // Боевик
	if err != nil {
		return nil, err
	}
	compilation = append(compilation, nextMC)

	nextMC, err = s.GetByCountry(ctx, &proto.GetByIDOptions{
		ID:     3,
		Limit:  0,
		Offset: 0,
	}) // США
	if err != nil {
		return nil, err
	}
	compilation = append(compilation, nextMC)

	return &proto.MovieCompilationsArr{MovieCompilations: compilation}, nil
}

func (s *Service) GetByGenre(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetByGenre(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetByCountry(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetByCountry(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetByMovie(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetByMovie(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetByPerson(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetByPerson(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetTopByYear(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetTopByYear(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetTop(ctx context.Context, in *proto.GetCompilationOptions) (*proto.MovieCompilation, error) {
	if in.Limit > 40 {
		in.Limit = 40
	}

	compilation, err := s.MCStorage.GetTop(int(in.Limit))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}
	return compilation, nil
}

func (s *Service) GetFavorites(ctx context.Context, in *proto.GetFavoritesOptions) (*proto.MovieCompilationsArr, error) {
	compilation, err := s.MCStorage.GetFavorites(in)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation.MovieCompilations[0])
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation.MovieCompilations[0])
	if err != nil {
		return nil, err
	}

	err = s.fillGenres(compilation.MovieCompilations[1])
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation.MovieCompilations[1])
	if err != nil {
		return nil, err
	}
	result := &proto.MovieCompilationsArr{MovieCompilations: compilation.MovieCompilations}
	return result, nil
}

func (s *Service) GetFavoritesFilms(ctx context.Context, in *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetFavoritesFilms(in)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}

	return compilation, nil
}

func (s *Service) GetFavoritesSeries(ctx context.Context, in *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	compilation, err := s.MCStorage.GetFavoritesSeries(in)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(compilation)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(compilation)
	if err != nil {
		return nil, err
	}

	return compilation, nil
}

func (s *Service) Find(ctx context.Context, in *proto.SearchText) (*proto.SearchCompilation, error) {
	data := strings.Join(strings.Fields(in.Text), " & ")
	dataByPartial := strings.TrimSpace(in.Text)
	if len(dataByPartial) == 0 {
		return &proto.SearchCompilation{
			Movies:  nil,
			Series:  nil,
			Persons: nil,
		}, nil
	}

	movieCompilations, err := s.MCStorage.FindMovie(data, true)
	if err != nil {
		return nil, err
	}

	if len(movieCompilations.Movies) < constants.MoviesSearchLimit {
		movieCompilationsByPartial, err := s.MCStorage.FindMovieByPartial(dataByPartial, true)
		if err != nil {
			return nil, err
		}

		for _, movie := range movieCompilationsByPartial.Movies {
			if !contains.MovieContains(movieCompilations.Movies, movie.ID) && len(movieCompilations.Movies) < constants.MoviesSearchLimit {
				movieCompilations.Movies = append(movieCompilations.Movies, movie)
			}
		}
	}

	err = s.fillGenres(movieCompilations)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(movieCompilations)
	if err != nil {
		return nil, err
	}

	seriesCompilations, err := s.MCStorage.FindMovie(data, false)
	if err != nil {
		return nil, err
	}

	if len(seriesCompilations.Movies) < constants.MoviesSearchLimit {
		seriesCompilationsByPartial, err := s.MCStorage.FindMovieByPartial(dataByPartial, false)
		if err != nil {
			return nil, err
		}

		for _, series := range seriesCompilationsByPartial.Movies {
			if !contains.MovieContains(seriesCompilations.Movies, series.ID) && len(seriesCompilations.Movies) < constants.MoviesSearchLimit {
				seriesCompilations.Movies = append(seriesCompilations.Movies, series)
			}
		}
	}

	err = s.fillGenres(seriesCompilations)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(seriesCompilations)
	if err != nil {
		return nil, err
	}

	personsCompilations, err := s.staffStorage.FindPerson(data)
	if err != nil {
		return nil, err
	}

	if len(personsCompilations.Persons) < constants.PersonsSearchLimit {
		personsCompilationsByPartial, err := s.staffStorage.FindPersonByPartial(dataByPartial)
		if err != nil {
			return nil, err
		}

		for _, person := range personsCompilationsByPartial.Persons {
			if !contains.PersonContains(personsCompilations.Persons, person.ID) && len(personsCompilations.Persons) < constants.PersonsSearchLimit {
				personsCompilations.Persons = append(personsCompilations.Persons, person)
			}
		}
	}

	for i := range personsCompilations.Persons {
		personsCompilations.Persons[i].Photo, err = images.GenerateFileURL(personsCompilations.Persons[i].Photo, "persons")
		if err != nil {
			return nil, err
		}
	}

	returnData := &proto.SearchCompilation{
		Movies:  movieCompilations.Movies,
		Series:  seriesCompilations.Movies,
		Persons: personsCompilations.Persons,
	}

	return returnData, nil
}
