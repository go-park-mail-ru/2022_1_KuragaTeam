package repository

import (
	"context"
	"database/sql"
	"myapp/internal/models"
	"myapp/internal/user"
	"myapp/internal/utils/constants"
	"myapp/internal/utils/hash"
	"myapp/internal/utils/images"

	"github.com/gofrs/uuid"
	"github.com/minio/minio-go/v7"
)

type userStorage struct {
	db *sql.DB
}

type imageStorage struct {
	client *minio.Client
}

func NewStorage(db *sql.DB) user.Storage {
	return &userStorage{db: db}
}

func NewImageStorage(client *minio.Client) user.ImageStorage {
	return &imageStorage{client: client}
}

func (us *userStorage) GetUserProfile(userID int64) (*models.User, error) {
	sqlScript := "SELECT username, email, avatar FROM users WHERE id=$1"

	var name, email, avatar string
	err := us.db.QueryRow(sqlScript, userID).Scan(&name, &email, &avatar)

	if err != nil {
		return nil, err
	}

	avatarUrl, err := images.GenerateFileURL(avatar, constants.UserObjectsBucketName)
	if err != nil {
		return nil, err
	}

	userData := models.User{
		Name:   name,
		Email:  email,
		Avatar: avatarUrl,
	}

	return &userData, nil
}

func (us *userStorage) EditProfile(user *models.User) error {
	sqlScript := "SELECT username, password, salt FROM users WHERE id=$1"

	var oldName, oldPassword, oldSalt string
	err := us.db.QueryRow(sqlScript, user.ID).Scan(&oldName, &oldPassword, &oldSalt)
	if err != nil {
		return err
	}

	notChangedPassword, _ := hash.ComparePasswords(oldPassword, oldSalt, user.Password)

	switch {
	case notChangedPassword == false && len(user.Password) != 0 && user.Name != oldName && len(user.Name) != 0:
		salt, err := uuid.NewV4()
		if err != nil {
			return err
		}

		hashPassword, err := hash.HashAndSalt(user.Password, salt.String())
		if err != nil {
			return err
		}

		sqlScript := "UPDATE users SET username = $2, password = $3, salt = $4 WHERE id = $1"

		_, err = us.db.Exec(sqlScript, user.ID, user.Name, hashPassword, salt)
		if err != nil {
			return err
		}

		return nil

	case notChangedPassword == false && len(user.Password) != 0:
		salt, err := uuid.NewV4()
		if err != nil {
			return err
		}

		hashPassword, err := hash.HashAndSalt(user.Password, salt.String())
		if err != nil {
			return err
		}

		sqlScript := "UPDATE users SET password = $2, salt = $3 WHERE id = $1"

		_, err = us.db.Exec(sqlScript, user.ID, hashPassword, salt)
		if err != nil {
			return err
		}

		return nil

	default:
		sqlScript := "UPDATE users SET username = $2 WHERE id = $1"

		_, err = us.db.Exec(sqlScript, user.ID, user.Name)
		if err != nil {
			return err
		}

		return nil
	}
}

func (us *userStorage) EditAvatar(user *models.User) (string, error) {
	sqlScript := "SELECT avatar FROM users WHERE id=$1"

	var oldAvatar string
	err := us.db.QueryRow(sqlScript, user.ID).Scan(&oldAvatar)
	if err != nil {
		return "", err
	}

	if len(user.Avatar) != 0 {
		sqlScript := "UPDATE users SET avatar = $2 WHERE id = $1"

		_, err = us.db.Exec(sqlScript, user.ID, user.Avatar)
		if err != nil {
			return "", err
		}

		return oldAvatar, nil
	}

	return "", nil
}

func (i imageStorage) UploadFile(input models.UploadInput) (string, error) {
	imageName := images.GenerateObjectName(input)

	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := i.client.PutObject(
		context.Background(),
		constants.UserObjectsBucketName, // Константа с именем бакета
		imageName,
		input.File,
		input.Size,
		opts,
	)
	if err != nil {
		return "", err
	}

	return imageName, nil
}

func (i imageStorage) DeleteFile(name string) error {
	opts := minio.RemoveObjectOptions{}

	err := i.client.RemoveObject(
		context.Background(),
		constants.UserObjectsBucketName,
		name,
		opts,
	)
	if err != nil {
		return err
	}

	return nil
}

func (us userStorage) GetAvatar(userID int64) (string, error) {
	sqlScript := "SELECT avatar FROM users WHERE id=$1"

	var avatar string
	err := us.db.QueryRow(sqlScript, userID).Scan(&avatar)

	if err != nil {
		return "", err
	}

	return avatar, nil
}
