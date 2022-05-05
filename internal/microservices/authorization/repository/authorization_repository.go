package repository

import (
	"database/sql"
	"myapp/internal/constants"
	"myapp/internal/microservices/authorization"
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/microservices/authorization/utils/hash"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
)

type Storage struct {
	db    *sql.DB
	redis *redis.Pool
}

func NewStorage(db *sql.DB, redis *redis.Pool) authorization.Storage {
	return &Storage{db: db, redis: redis}
}

func (s Storage) IsUserExists(data *proto.LogInData) (int64, error) {
	var userID int64
	sqlScript := "SELECT id, password, salt FROM users WHERE email=$1"
	rows, err := s.db.Query(sqlScript, data.Email)
	if err != nil {
		return userID, err
	}

	// убедимся, что всё закроется при выходе из программы
	defer func() {
		rows.Close()
	}()

	// Из базы пришел пустой запрос, значит пользователя в базе данных нет
	if !rows.Next() {
		return userID, constants.ErrWrongData
	}

	var (
		id             int64
		password, salt string
	)
	err = rows.Scan(&id, &password, &salt)

	userID = id
	// выход при ошибке
	if err != nil {
		return userID, err
	}

	_, err = hash.ComparePasswords(password, salt, data.Password)
	if err != nil {
		return userID, constants.ErrWrongData
	}

	return userID, nil
}

func (s Storage) IsUserUnique(email string) (bool, error) {
	sqlScript := "SELECT id FROM users WHERE email=$1"
	rows, err := s.db.Query(sqlScript, email)

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

func (s Storage) CreateUser(data *proto.SignUpData) (int64, error) {
	var userID int64

	salt, err := uuid.NewV4()
	if err != nil {
		return userID, err
	}

	hashPassword, err := hash.HashAndSalt(data.Password, salt.String())
	if err != nil {
		return userID, err
	}

	sqlScript := "INSERT INTO users(username, email, password, salt, avatar, subscription_expires) VALUES($1, $2, $3, $4, $5, LOCALTIMESTAMP) RETURNING id"

	if err = s.db.QueryRow(sqlScript, data.Name, data.Email, hashPassword, salt, constants.DefaultImage).Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

func (s Storage) StoreSession(userID int64) (string, error) {
	connRedis := s.redis.Get()
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

func (s Storage) GetUserId(session string) (int64, error) {
	connRedis := s.redis.Get()
	defer connRedis.Close()

	userID, err := redis.Int64(connRedis.Do("GET", session))
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (s Storage) DeleteSession(session string) error {
	connRedis := s.redis.Get()
	defer connRedis.Close()

	_, err := connRedis.Do("DEL", session)
	if err != nil {
		return err
	}

	return nil
}
