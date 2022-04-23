package user

import "myapp/internal/models"

type Storage interface {
	GetUserProfile(userID int64) (*models.User, error)
	EditProfile(user *models.User) error
	EditAvatar(user *models.User) (string, error)
	GetAvatar(userID int64) (string, error)
}

type ImageStorage interface {
	UploadFile(input models.UploadInput) (string, error) // Загрузка файлов
	DeleteFile(string) error                             // Удаление файлов
}
