package repository

import (
	"bytes"
	"context"
	"database/sql"
	"myapp/internal/constants"
	"myapp/internal/microservices/authorization/utils/hash"
	"myapp/internal/microservices/profile"
	"myapp/internal/microservices/profile/proto"
	"myapp/internal/microservices/profile/utils/images"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	db    *sql.DB
	minio *minio.Client
}

func NewStorage(db *sql.DB, minio *minio.Client) profile.Storage {
	return &Storage{db: db, minio: minio}
}

func (s Storage) GetUserProfile(userID int64) (*proto.ProfileData, error) {
	sqlScript := "SELECT username, email, avatar FROM users WHERE id=$1"

	var name, email, avatar string
	err := s.db.QueryRow(sqlScript, userID).Scan(&name, &email, &avatar)

	if err != nil {
		return nil, err
	}

	avatarUrl, err := images.GenerateFileURL(avatar, constants.UserObjectsBucketName)
	if err != nil {
		return nil, err
	}

	return &proto.ProfileData{
		Name:   name,
		Email:  email,
		Avatar: avatarUrl,
	}, nil
}

func (s Storage) EditProfile(data *proto.EditProfileData) error {
	sqlScript := "SELECT username, password, salt FROM users WHERE id=$1"

	var oldName, oldPassword, oldSalt string
	err := s.db.QueryRow(sqlScript, data.ID).Scan(&oldName, &oldPassword, &oldSalt)
	if err != nil {
		return err
	}

	notChangedPassword, _ := hash.ComparePasswords(oldPassword, oldSalt, data.Password)

	switch {
	case notChangedPassword == false && len(data.Password) != 0 && data.Name != oldName && len(data.Name) != 0:
		salt, err := uuid.NewV4()
		if err != nil {
			return err
		}

		hashPassword, err := hash.HashAndSalt(data.Password, salt.String())
		if err != nil {
			return err
		}

		sqlScript := "UPDATE users SET username = $2, password = $3, salt = $4 WHERE id = $1"

		_, err = s.db.Exec(sqlScript, data.ID, data.Name, hashPassword, salt)
		if err != nil {
			return err
		}

		return nil

	case notChangedPassword == false && len(data.Password) != 0:
		salt, err := uuid.NewV4()
		if err != nil {
			return err
		}

		hashPassword, err := hash.HashAndSalt(data.Password, salt.String())
		if err != nil {
			return err
		}

		sqlScript := "UPDATE users SET password = $2, salt = $3 WHERE id = $1"

		_, err = s.db.Exec(sqlScript, data.ID, hashPassword, salt)
		if err != nil {
			return err
		}

		return nil

	default:
		sqlScript := "UPDATE users SET username = $2 WHERE id = $1"

		_, err = s.db.Exec(sqlScript, data.ID, data.Name)
		if err != nil {
			return err
		}

		return nil
	}
}

func (s Storage) EditAvatar(data *proto.EditAvatarData) (string, error) {
	sqlScript := "SELECT avatar FROM users WHERE id=$1"

	var oldAvatar string
	err := s.db.QueryRow(sqlScript, data.ID).Scan(&oldAvatar)
	if err != nil {
		return "", err
	}

	if len(data.Avatar) != 0 {
		sqlScript := "UPDATE users SET avatar = $2 WHERE id = $1"

		_, err = s.db.Exec(sqlScript, data.ID, data.Avatar)
		if err != nil {
			return "", err
		}

		return oldAvatar, nil
	}

	return "", nil
}

func (s Storage) GetAvatar(userID int64) (string, error) {
	sqlScript := "SELECT avatar FROM users WHERE id=$1"

	var avatar string
	err := s.db.QueryRow(sqlScript, userID).Scan(&avatar)

	if err != nil {
		return "", err
	}

	return avatar, nil
}

func (s Storage) UploadAvatar(data *proto.UploadInputFile) (string, error) {
	imageName := images.GenerateObjectName(data.ID)

	opts := minio.PutObjectOptions{
		ContentType:  data.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := s.minio.PutObject(
		context.Background(),
		constants.UserObjectsBucketName, // Константа с именем бакета
		imageName,
		bytes.NewReader(data.File),
		data.Size,
		opts,
	)
	if err != nil {
		return "", err
	}

	return imageName, nil
}

func (s Storage) DeleteFile(name string) error {
	opts := minio.RemoveObjectOptions{}

	err := s.minio.RemoveObject(
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

func (s Storage) AddLike(data *proto.LikeData) error {
	sqlScript := "SELECT likes FROM users WHERE id=$1"

	likes := make([]int64, 0)
	err := s.db.QueryRow(sqlScript, data.UserID).Scan(pq.Array(&likes))
	if err != nil {
		return err
	}

	for _, a := range likes {
		if a == data.MovieID {
			return nil
		}
	}

	sqlScript = "UPDATE users SET likes = array_append(likes, $2) WHERE id=$1"

	_, err = s.db.Exec(sqlScript, data.UserID, data.MovieID)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) RemoveLike(data *proto.LikeData) error {
	sqlScript := "SELECT likes FROM users WHERE id=$1"

	likes := make([]int64, 0)
	err := s.db.QueryRow(sqlScript, data.UserID).Scan(pq.Array(&likes))
	if err != nil {
		return err
	}

	for _, a := range likes {
		if a == data.MovieID {
			sqlScript = "UPDATE users SET likes = array_remove(likes, $2) WHERE id=$1"

			_, err = s.db.Exec(sqlScript, data.UserID, data.MovieID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s Storage) GetFavorites(userID int64) (*proto.Favorites, error) {
	sqlScript := "SELECT likes FROM users WHERE id=$1"

	likes := make([]int64, 0)
	err := s.db.QueryRow(sqlScript, userID).Scan(pq.Array(&likes))
	if err != nil {
		return &proto.Favorites{}, err
	}
	favorites := &proto.Favorites{Id: likes}

	return favorites, nil
}

func (s Storage) GetRating(data *proto.MovieRating) (*proto.Rating, error) {
	sqlScript := "SELECT rating FROM rating WHERE user_id=$1 AND movie_id = $2"

	returnValue := proto.Rating{}
	err := s.db.QueryRow(sqlScript, data.UserID, data.MovieID).Scan(&returnValue.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			returnValue.Rating = -1
			return &returnValue, nil
		}
		return nil, err
	}
	return &returnValue, nil
}
