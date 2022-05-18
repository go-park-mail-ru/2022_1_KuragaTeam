package usecase

import (
	"context"
	"errors"
	"log"
	"myapp/internal/constants"
	"myapp/internal/microservices/profile"
	"myapp/internal/microservices/profile/proto"
	"time"

	"github.com/gofrs/uuid"
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

	return userData, nil
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

func (s *Service) GetPaymentsToken(ctx context.Context, data *proto.UserID) (*proto.Token, error) {
	token, _ := uuid.NewV4()

	err := s.storage.SetToken(token.String(), data.ID, int64(time.Hour.Seconds()))
	if err != nil {
		return &proto.Token{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Token{Token: token.String()}, nil
}

func (s *Service) CheckPaymentsToken(ctx context.Context, data *proto.CheckTokenData) (*proto.Empty, error) {
	id, err := s.storage.GetIdByToken(data.Token)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	if id != data.Id {
		return &proto.Empty{}, status.Error(codes.InvalidArgument, constants.WrongToken.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) CheckToken(ctx context.Context, data *proto.Token) (*proto.Empty, error) {
	_, err := s.storage.GetIdByToken(data.Token)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) CreatePayment(ctx context.Context, data *proto.CheckTokenData) (*proto.Empty, error) {
	err := s.storage.CreatePayment(data.Token, data.Id, float64(constants.Price))
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) CreateSubscribe(ctx context.Context, data *proto.SubscribeData) (*proto.Empty, error) {
	err := s.storage.CheckCountPaymentsByToken(data.Token)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	id, amount, err := s.storage.GetAmountByToken(data.Token)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	log.Println(amount, data.Amount)
	if amount != data.Amount {
		return &proto.Empty{}, status.Error(codes.Internal, constants.WrongAmount.Error())
	}

	err = s.storage.UpdatePayment(data.Token, id)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	err = s.storage.CreateSubscribe(id)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) IsSubscription(ctx context.Context, data *proto.UserID) (*proto.Empty, error) {
	err := s.storage.IsSubscription(data.ID)
	if err != nil {
		if errors.Is(err, constants.NoSubscription) {
			return &proto.Empty{}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (s *Service) GetMovieRating(ctx context.Context, data *proto.MovieRating) (*proto.Rating, error) {
	rating, err := s.storage.GetRating(data)
	if err != nil {
		return &proto.Rating{}, status.Error(codes.Internal, err.Error())
	}

	return rating, nil
}
