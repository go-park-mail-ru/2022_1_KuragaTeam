package usecase

import (
	"io"
	"myapp/internal/user"
	"myapp/internal/utils/constants"
	"myapp/internal/utils/validation"
)

type service struct {
	storage      user.Storage
	redisStore   user.RedisStore
	imageStorage user.ImageStorage
}

func NewService(storage user.Storage, redisStore user.RedisStore, imageStorage user.ImageStorage) user.Service {
	return &service{
		storage:      storage,
		redisStore:   redisStore,
		imageStorage: imageStorage,
	}
}

func (s *service) SignUp(dto *user.CreateUserDTO) (string, string, error) {
	userModel := &user.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	if err := validation.ValidateUser(userModel); err != nil {
		return "", "", err
	}

	isUnique, err := s.storage.IsUserUnique(userModel)
	if err != nil {
		return "", "", err
	}

	if !isUnique {
		return "", "ERROR: Email is not unique", nil
	}

	userID, err := s.storage.CreateUser(userModel)
	if err != nil {
		return "", "", err
	}

	session, err := s.redisStore.StoreSession(userID)
	if err != nil {
		return "", "", err
	}

	return session, "", nil
}

func (s *service) LogIn(dto *user.LogInUserDTO) (string, error) {
	userModel := &user.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	userID, userExists, err := s.storage.IsUserExists(userModel)
	if err != nil {
		return "", err
	}

	if !userExists {
		return "", nil
	}

	session, err := s.redisStore.StoreSession(userID)
	if err != nil {
		return "", err
	}

	return session, nil
}

func (s *service) LogOut(session string) error {
	return s.redisStore.DeleteSession(session)
}

func (s *service) CheckAuthorization(session string) (int64, error) {
	userID, err := s.redisStore.GetUserId(session)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (s *service) GetUserProfile(userID int64) (*user.ProfileUserDTO, error) {
	userData, err := s.storage.GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	userDTO := user.ProfileUserDTO{
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar,
	}

	return &userDTO, nil
}

func (s *service) EditProfile(dto *user.EditProfileDTO) error {
	userModel := &user.User{
		ID:       dto.ID,
		Name:     dto.Name,
		Password: dto.Password,
	}

	err := s.storage.EditProfile(userModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) EditAvatar(dto *user.EditAvatarDTO) error {
	userModel := &user.User{
		ID:     dto.ID,
		Avatar: dto.Avatar,
	}

	oldAvatar, err := s.storage.EditAvatar(userModel)
	if err != nil {
		return err
	}

	if oldAvatar != constants.DefaultImage {
		err = s.imageStorage.DeleteFile(oldAvatar)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) UploadAvatar(file io.Reader, size int64, contentType string, userID int64) (string, error) {
	uploadImage := user.UploadInput{
		UserID:      userID,
		File:        file,
		Size:        size,
		ContentType: contentType,
	}

	return s.imageStorage.UploadFile(uploadImage)
}
