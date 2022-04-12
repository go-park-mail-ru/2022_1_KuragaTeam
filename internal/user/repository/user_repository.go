package repository

import (
	"context"
	"database/sql"
	"myapp/internal/user"
	"myapp/internal/utils/constants"
	"myapp/internal/utils/hash"
	"myapp/internal/utils/images"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"github.com/minio/minio-go/v7"
)

type userStorage struct {
	db *sql.DB
}

type redisStore struct {
	redis *redis.Pool
}

type imageStorage struct {
	client *minio.Client
}

func NewStorage(db *sql.DB) user.Storage {
	return &userStorage{db: db}
}

func NewRedisStore(redis *redis.Pool) user.RedisStore {
	return &redisStore{redis: redis}
}

func NewImageStorage(client *minio.Client) user.ImageStorage {
	return &imageStorage{client: client}
}

func (us *userStorage) IsUserExists(userModel *user.User) (int64, bool, error) {
	var userID int64
	sqlScript := "SELECT id, password, salt FROM users WHERE email=$1"
	rows, err := us.db.Query(sqlScript, userModel.Email)
	if err != nil {
		return userID, false, err
	}

	// убедимся, что всё закроется при выходе из программы
	defer func() {
		rows.Close()
	}()

	// Из базы пришел пустой запрос, значит пользователя в базе данных нет
	if !rows.Next() {
		return userID, false, nil
	}

	signInUser := user.User{}
	err = rows.Scan(&signInUser.ID, &signInUser.Password, &signInUser.Salt)

	userID = signInUser.ID
	// выход при ошибке
	if err != nil {
		return userID, false, err
	}

	result, err := hash.ComparePasswords(signInUser.Password, signInUser.Salt, userModel.Password)
	if err != nil {
		return userID, false, constants.ErrWrongData
	}

	return userID, result, nil
}

func (us *userStorage) IsUserUnique(userModel *user.User) (bool, error) {
	sqlScript := "SELECT id FROM users WHERE email=$1"
	rows, err := us.db.Query(sqlScript, userModel.Email)

	if err != nil {
		return false, err
	}

	defer func() {
		rows.Close()
	}()

	if rows.Next() { // Пользователь с таким email зарегистрирован
		return false, nil
	}
	return true, nil
}

func (us *userStorage) CreateUser(userModel *user.User) (int64, error) {
	var userID int64

	salt, err := uuid.NewV4()
	if err != nil {
		return userID, err
	}

	hashPassword, err := hash.HashAndSalt(userModel.Password, salt.String())
	if err != nil {
		return userID, err
	}

	sqlScript := "INSERT INTO users(username, email, password, salt, avatar, subscription_expires) VALUES($1, $2, $3, $4, $5, LOCALTIMESTAMP) RETURNING id"

	if err = us.db.QueryRow(sqlScript, userModel.Name, userModel.Email, hashPassword, salt, constants.DefaultImage).Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

func (r *redisStore) StoreSession(userID int64) (string, error) {
	connRedis := r.redis.Get()
	defer connRedis.Close()

	session, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	_, err = connRedis.Do("SET", session, userID, "EX", int64(30*24*time.Hour.Seconds()))

	if err != nil {
		return "", err
	}

	return session.String(), nil
}

func (r *redisStore) GetUserId(session string) (int64, error) {
	connRedis := r.redis.Get()
	defer connRedis.Close()

	userID, err := redis.Int64(connRedis.Do("GET", session))
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (r *redisStore) DeleteSession(session string) error {
	connRedis := r.redis.Get()
	defer connRedis.Close()

	_, err := connRedis.Do("DEL", session)
	if err != nil {
		return err
	}

	return nil
}

func (us *userStorage) GetUserProfile(userID int64) (*user.User, error) {
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

	userData := user.User{
		Name:   name,
		Email:  email,
		Avatar: avatarUrl,
	}

	return &userData, nil
}

func (us *userStorage) EditProfile(user *user.User) error {
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

func (us *userStorage) EditAvatar(user *user.User) (string, error) {
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

func (i imageStorage) UploadFile(input user.UploadInput) (string, error) {
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
