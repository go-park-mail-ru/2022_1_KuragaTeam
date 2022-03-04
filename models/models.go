package models

import (
	"context"
	"errors"
	"myapp/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       string `json:"id"` // validate:"nonzero"`
	Name     string `json:"username" validate:"nonzero"`
	Email    string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" validate:"min=8"`
}

var ErrWrongPassword = errors.New("wrong password")

// Используется LoginUserHandler.
// Проверяет, что пользователь есть в базе данных.
func IsUserExists(dbPool *pgxpool.Pool, user User) (bool, error) {
	sql := "SELECT email, password FROM USERS WHERE email=$1"
	rows, err := dbPool.Query(context.Background(), sql, user.Email)
	if err != nil {
		return false, err
	}

	// убедимся, что всё закроется при выходе из программы
	defer func() {
		rows.Close()
	}()

	// Из базы пришел пустой запрос, значит пользователя в базе данных нет
	if !rows.Next() {
		return false, nil
	}

	var signInUser User
	err = rows.Scan(&signInUser.Email, &signInUser.Password)

	// выход при ошибке
	if err != nil {
		return false, err
	}

	result, err := utils.ComparePasswords(signInUser.Password, []byte(user.Password))
	if err != nil {
		return false, ErrWrongPassword
	}

	result = result && signInUser.Email == user.Email

	return result, nil
}

// Используется CreateUserHandler.
// email должен быть уникален
func IsUserUnique(dbPool *pgxpool.Pool, user User) (bool, error) {
	sql := "SELECT * FROM users WHERE email=$1"
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
func CreateUser(dbPool *pgxpool.Pool, user User) error {
	hashPassword, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		return err
	}

	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}

	sql := "INSERT INTO users(username, email, password) VALUES($1, $2, $3)"
	conn.QueryRow(context.Background(), sql, user.Name, user.Email, hashPassword)

	return nil
}
