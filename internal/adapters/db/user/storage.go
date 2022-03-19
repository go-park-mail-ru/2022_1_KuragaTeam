package user

import (
	"context"
	"myapp/constants"
	"myapp/internal/domain/user"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type userStorage struct {
	db *pgxpool.Pool
}

type redisStore struct {
	redis *redis.Pool
}

func NewStorage(db *pgxpool.Pool) user.Storage {
	return &userStorage{db: db}
}

func NewRedisStore(redis *redis.Pool) user.RedisStore {
	return &redisStore{redis: redis}
}

func HashAndSalt(pwd string, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+salt), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPwd string, salt string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd+salt))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (us *userStorage) IsUserExists(userModel *user.User) (int64, bool, error) {
	var userID int64
	sql := "SELECT id, password, salt FROM users WHERE email=$1"
	rows, err := us.db.Query(context.Background(), sql, userModel.Email)
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

	result, err := ComparePasswords(signInUser.Password, signInUser.Salt, userModel.Password)
	if err != nil {
		return userID, false, constants.ErrWrongData
	}

	return userID, result, nil
}

func (us *userStorage) IsUserUnique(userModel *user.User) (bool, error) {
	sql := "SELECT * FROM users WHERE email=$1"
	rows, err := us.db.Query(context.Background(), sql, userModel.Email)

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

	hashPassword, err := HashAndSalt(userModel.Password, salt.String())
	if err != nil {
		return userID, err
	}

	sql := "INSERT INTO users(username, email, password, salt, avatar, subscription_expires) VALUES($1, $2, $3, $4, '/avatars/default_avatar.png', LOCALTIMESTAMP) RETURNING id"

	if err = us.db.QueryRow(context.Background(), sql, userModel.Name, userModel.Email, hashPassword, salt).Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

func (us *userStorage) GetUserMainPage(userID int64) (*user.User, error) {
	sql := "SELECT username, avatar FROM users WHERE id=$1"

	var name, avatar string
	err := us.db.QueryRow(context.Background(), sql, userID).Scan(&name, &avatar)

	if err != nil {
		return nil, err
	}

	userData := user.User{
		Name:   name,
		Avatar: avatar,
	}

	return &userData, nil
}

func (r *redisStore) StoreSession(userID int64) (string, error) {
	connRedis := r.redis.Get()
	defer connRedis.Close()

	session, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	_, err = connRedis.Do("SET", session, userID, "EX", int64(time.Hour.Seconds()))

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
	sql := "SELECT username, email FROM users WHERE id=$1"

	var name, email string
	err := us.db.QueryRow(context.Background(), sql, userID).Scan(&name, &email)

	if err != nil {
		return nil, err
	}

	userData := user.User{
		Name:  name,
		Email: email,
	}

	return &userData, nil
}

func (us *userStorage) EditProfile(user *user.User) error {
	//sql := "SELECT password, salt FROM users WHERE id=$1"
	//
	//var password, salt string
	//err := us.db.QueryRow(context.Background(), sql, user.ID).Scan(&password, &salt)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = ComparePasswords(user.Password, salt, password)
	//
	//if err != nil {
	//	return err
	//}

	salt, err := uuid.NewV4()
	if err != nil {
		return err
	}

	hashPassword, err := HashAndSalt(user.Password, salt.String())
	if err != nil {
		return err
	}

	sql := "UPDATE users SET username = $2, password = $3, salt = $4 WHERE id = $1"

	_, err = us.db.Exec(context.Background(), sql, user.ID, user.Name, hashPassword, salt)
	if err != nil {
		return err
	}

	return nil
}