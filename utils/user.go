package utils

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/models"
)

var ErrWrongPassword = errors.New("wrong password")

// Используется LoginUserHandler.
// Проверяет, что пользователь есть в базе данных.
func IsUserExists(dbPool *pgxpool.Pool, user models.User) (uint64, bool, error) {
	var userID uint64
	sql := "SELECT id, email, password, salt FROM USERS WHERE email=$1"
	rows, err := dbPool.Query(context.Background(), sql, user.Email)
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

	var signInUser models.User
	err = rows.Scan(&signInUser.ID, &signInUser.Email, &signInUser.Password, &signInUser.Salt)

	userID = signInUser.ID
	// выход при ошибке
	if err != nil {
		return userID, false, err
	}

	result, err := ComparePasswords(signInUser.Password, signInUser.Salt, user.Password)
	if err != nil {
		return userID, false, ErrWrongPassword
	}

	result = result && signInUser.Email == user.Email

	return userID, result, nil
}

// Используется CreateUserHandler.
// email должен быть уникален
func IsUserUnique(dbPool *pgxpool.Pool, user models.User) (bool, error) {
	sql := "SELECT * FROM users WHERE email=$1;"
	rows, err := dbPool.Query(context.Background(), sql, user.Email)

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

// Используется CreateUserHandler.
// Создает пользователя
func CreateUser(dbPool *pgxpool.Pool, user models.User) (uint64, error) {
	var userID uint64

	salt, err := uuid.NewV4()
	if err != nil {
		return userID, err
	}

	hashPassword, err := HashAndSalt(user.Password, salt.String())
	if err != nil {
		return userID, err
	}

	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return userID, err
	}

	sql := "INSERT INTO users(username, email, password, salt) VALUES($1, $2, $3, $4) RETURNING id;"
	if err = conn.QueryRow(context.Background(), sql, user.Name, user.Email, hashPassword, salt).Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}
