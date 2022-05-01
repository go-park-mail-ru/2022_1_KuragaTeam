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
	proto.UnimplementedMovieCompilationsServer

	MCStorage    compilations.Storage
	genreStorage genre.Storage
	staffStorage persons.Storage
}

func NewService(MCStorage compilations.Storage, genreStorage genre.Storage, staffStorage persons.Storage) *Service {
	return &Service{MCStorage: MCStorage, genreStorage: genreStorage, staffStorage: staffStorage}
}

func (s *Service) fillGenres(MC *proto.MovieCompilation) error {
	for i := 0; i < len(MC.Movies); i++ {
		nextGenres, err := s.genreStorage.GetByMovieID(int(MC.Movies[i].ID))
		if err != nil {
			return err
		}
		for _, nextGenre := range nextGenres {
			MC.Movies[i].Genre = append(MC.Movies[i].Genre, &proto.Genre{
				ID:   int32(nextGenre.ID),
				Name: nextGenre.Name,
			})
		}
	}
	return nil
}

func (s *Service) concatUrls(MC *proto.MovieCompilation) error {
	var err error
	for i, _ := range MC.Movies {
		MC.Movies[i].Picture, err = images.GenerateFileURL(MC.Movies[i].Picture, "posters")
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetAllMovies(ctx context.Context, in *proto.GetCompilationOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetAllMovies(int(in.Limit), int(in.Offset), true)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetAllSeries(ctx context.Context, in *proto.GetCompilationOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetAllMovies(int(in.Limit), int(in.Offset), false)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetMainCompilations(ctx context.Context, in *proto.GetMainCompilationsOptions) (*proto.MovieCompilationsArr, error) {

	MC := make([]*proto.MovieCompilation, 0)

	nextMC, err := s.GetTop(ctx, &proto.GetCompilationOptions{Limit: 12})
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	nextMC, err = s.GetTopByYear(ctx, &proto.GetByIDOptions{ID: 2011})
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	nextMC, err = s.GetByGenre(ctx, &proto.GetByIDOptions{ID: 2}) // Боевик
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	nextMC, err = s.GetByCountry(ctx, &proto.GetByIDOptions{ID: 3}) // США
	if err != nil {
		return nil, err
	}
	MC = append(MC, nextMC)

	return &proto.MovieCompilationsArr{MovieCompilations: MC}, nil
}

func (s *Service) GetByGenre(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByGenre(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetByCountry(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByCountry(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetByMovie(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByMovie(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetByPerson(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetByPerson(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetTopByYear(ctx context.Context, in *proto.GetByIDOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetTopByYear(int(in.ID))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetTop(ctx context.Context, in *proto.GetCompilationOptions) (*proto.MovieCompilation, error) {
	if in.Limit > 40 {
		in.Limit = 40
	}

	MC, err := s.MCStorage.GetTop(int(in.Limit))
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}
	return MC, nil
}

func (s *Service) GetFavorites(ctx context.Context, in *proto.GetFavoritesOptions) (*proto.MovieCompilationsArr, error) {
	MC, err := s.MCStorage.GetFavorites(in)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC.MovieCompilations[0])
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC.MovieCompilations[0])
	if err != nil {
		return nil, err
	}

	err = s.fillGenres(MC.MovieCompilations[1])
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC.MovieCompilations[1])
	if err != nil {
		return nil, err
	}
	result := &proto.MovieCompilationsArr{MovieCompilations: MC.MovieCompilations}
	return result, nil
}

func (s *Service) GetFavoritesFilms(ctx context.Context, in *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetFavoritesFilms(in)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}

	return MC, nil
}

func (s *Service) GetFavoritesSeries(ctx context.Context, in *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	MC, err := s.MCStorage.GetFavoritesSeries(in)
	if err != nil {
		return nil, err
	}
	err = s.fillGenres(MC)
	if err != nil {
		return nil, err
	}
	err = s.concatUrls(MC)
	if err != nil {
		return nil, err
	}

	return MC, nil
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
			if !contains.ContainsMovie(movieCompilations.Movies, movie.ID) && len(movieCompilations.Movies) < constants.MoviesSearchLimit {
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
			if !contains.ContainsMovie(seriesCompilations.Movies, series.ID) && len(seriesCompilations.Movies) < constants.MoviesSearchLimit {
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
			if !contains.ContainsPerson(personsCompilations.Persons, person.ID) && len(personsCompilations.Persons) < constants.PersonsSearchLimit {
				personsCompilations.Persons = append(personsCompilations.Persons, person)
			}
		}
	}

	for i, _ := range personsCompilations.Persons {
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
