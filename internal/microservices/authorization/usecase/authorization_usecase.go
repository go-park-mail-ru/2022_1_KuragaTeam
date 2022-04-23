package usecase

import (
	"context"
	"errors"
	"myapp/internal/microservices/authorization"
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/utils/constants"
	"myapp/internal/utils/validation"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	storage authorization.Storage
}

func NewService(storage authorization.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) SignUp(ctx context.Context, data *proto.SignUpData) (*proto.Cookie, error) {
	if err := validation.ValidateUser(data); err != nil {
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	isUnique, err := s.storage.IsUserUnique(data.Email)
	if err != nil {
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	if !isUnique {
		return &proto.Cookie{}, status.Error(codes.InvalidArgument, constants.EmailIsNotUnique.Error())
	}

	userID, err := s.storage.CreateUser(data)
	if err != nil {
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	session, err := s.storage.StoreSession(userID)
	if err != nil {
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Cookie{
		Cookie: session,
	}, nil
}

func (s *Service) LogIn(ctx context.Context, data *proto.LogInData) (*proto.Cookie, error) {
	userID, err := s.storage.IsUserExists(data)
	if err != nil {
		if errors.Is(err, constants.ErrWrongData) {
			return &proto.Cookie{}, status.Error(codes.NotFound, err.Error())
		}
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	session, err := s.storage.StoreSession(userID)
	if err != nil {
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Cookie{
		Cookie: session,
	}, nil
}

func (s *Service) LogOut(ctx context.Context, cookie *proto.Cookie) (*proto.Empty, error) {
	err := s.storage.DeleteSession(cookie.Cookie)
	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &proto.Empty{}, nil
}

func (s *Service) CheckAuthorization(ctx context.Context, cookie *proto.Cookie) (*proto.UserID, error) {
	userID, err := s.storage.GetUserId(cookie.Cookie)
	if err != nil {
		return &proto.UserID{ID: -1}, status.Error(codes.Internal, err.Error())
	}

	return &proto.UserID{ID: userID}, nil
}
