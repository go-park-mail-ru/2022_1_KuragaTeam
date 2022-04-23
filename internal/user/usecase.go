package user

import (
	"io"
	"myapp/internal/models"
)

type Service interface {
	GetUserProfile(userID int64) (*models.ProfileUserDTO, error)
	EditProfile(dto *models.EditProfileDTO) error
	EditAvatar(dto *models.EditAvatarDTO) error
	UploadAvatar(file io.Reader, size int64, contentType string, userID int64) (string, error)
	GetAvatar(userID int64) (string, error)
}
