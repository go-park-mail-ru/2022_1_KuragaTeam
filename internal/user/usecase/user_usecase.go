package usecase

import (
	"io"
	"myapp/internal/models"
	"myapp/internal/user"
	"myapp/internal/utils/constants"
)

type service struct {
	storage      user.Storage
	imageStorage user.ImageStorage
}

func NewService(storage user.Storage, imageStorage user.ImageStorage) user.Service {
	return &service{
		storage:      storage,
		imageStorage: imageStorage,
	}
}

func (s *service) GetUserProfile(userID int64) (*models.ProfileUserDTO, error) {
	userData, err := s.storage.GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	userDTO := models.ProfileUserDTO{
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar,
	}

	return &userDTO, nil
}

func (s *service) EditProfile(dto *models.EditProfileDTO) error {
	userModel := &models.User{
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

func (s *service) EditAvatar(dto *models.EditAvatarDTO) error {
	userModel := &models.User{
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
	uploadImage := models.UploadInput{
		UserID:      userID,
		File:        file,
		Size:        size,
		ContentType: contentType,
	}

	return s.imageStorage.UploadFile(uploadImage)
}

func (s *service) GetAvatar(userID int64) (string, error) {
	return s.storage.GetAvatar(userID)
}
