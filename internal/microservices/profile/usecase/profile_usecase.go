package usecase

import (
	"context"
	"myapp/internal/constants"
	"myapp/internal/microservices/profile"
	"myapp/internal/microservices/profile/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	storage profile.Storage
}

func NewService(storage profile.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) GetUserProfile(ctx context.Context, userID *proto.UserID) (*proto.ProfileData, error) {
	userData, err := s.storage.GetUserProfile(userID.ID)
	if err != nil {
		return &proto.ProfileData{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.ProfileData{
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar,
	}, nil
}

func (s *Service) EditProfile(ctx context.Context, data *proto.EditProfileData) (*proto.Empty, error) {
	err := s.storage.EditProfile(data)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) EditAvatar(ctx context.Context, data *proto.EditAvatarData) (*proto.Empty, error) {
	oldAvatar, err := s.storage.EditAvatar(data)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	if oldAvatar != constants.DefaultImage {
		err = s.storage.DeleteFile(oldAvatar)
		if err != nil {
			return &proto.Empty{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &proto.Empty{}, nil
}

func (s *Service) UploadAvatar(ctx context.Context, data *proto.UploadInputFile) (*proto.FileName, error) {
	name, err := s.storage.UploadAvatar(data)
	if err != nil {
		return &proto.FileName{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.FileName{Name: name}, nil
}

func (s *Service) GetAvatar(ctx context.Context, userID *proto.UserID) (*proto.FileName, error) {
	name, err := s.storage.GetAvatar(userID.ID)
	if err != nil {
		return &proto.FileName{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.FileName{Name: name}, nil
}

func (s *Service) AddLike(ctx context.Context, data *proto.LikeData) (*proto.Empty, error) {
	err := s.storage.AddLike(data)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) RemoveLike(ctx context.Context, data *proto.LikeData) (*proto.Empty, error) {
	err := s.storage.RemoveLike(data)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) GetFavorites(ctx context.Context, data *proto.UserID) (*proto.Favorites, error) {
	favorites, err := s.storage.GetFavorites(data.ID)
	if err != nil {
		return &proto.Favorites{}, status.Error(codes.Internal, err.Error())
	}

	return favorites, nil
}

func (s *Service) GetMovieRating(ctx context.Context, data *proto.MovieRating) (*proto.Rating, error) {
	rating, err := s.storage.GetRating(data)
	if err != nil {
		return &proto.Rating{}, status.Error(codes.Internal, err.Error())
	}

	return rating, nil
}
