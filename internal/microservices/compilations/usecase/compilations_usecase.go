package usecase

import (
	"golang.org/x/net/context"
	"myapp/internal/genre"
	"myapp/internal/microservices/compilations"
	"myapp/internal/microservices/compilations/proto"
	"myapp/internal/utils/images"
)

type Service struct {
	proto.UnimplementedMovieCompilationsServer

	MCStorage    compilations.Storage
	genreStorage genre.Storage
}

func NewService(MCStorage compilations.Storage, genreStorage genre.Storage) *Service {
	return &Service{MCStorage: MCStorage, genreStorage: genreStorage}
}

func (s *Service) fillGenres(MC *proto.MovieCompilation) error {
	for i := 0; i < len(MC.Movies); i++ {
		var err error
		MC.Movies[i].Genre, err = s.genreStorage.GetByMovieID(int(MC.Movies[i].ID))
		if err != nil {
			return err
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
	MC, err := s.MCStorage.GetAllMovies(int(in.Limit), int(in.Offset))
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
