package models

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"myapp/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

var ErrWrongPassword = errors.New("wrong password")

// Используется LoginUserHandler.
// Проверяет, что пользователь есть в базе данных.
func IsUserExists(dbPool *pgxpool.Pool, user User) (bool, error) {
	sql := "SELECT email, password, salt FROM USERS WHERE email=$1"
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
	err = rows.Scan(&signInUser.Email, &signInUser.Password, &signInUser.Salt)

	// выход при ошибке
	if err != nil {
		return false, err
	}

	result, err := utils.ComparePasswords(signInUser.Password, signInUser.Salt, user.Password)
	if err != nil {
		return false, ErrWrongPassword
	}

	result = result && signInUser.Email == user.Email

	return result, nil
}

// Используется CreateUserHandler.
// email должен быть уникален
func IsUserUnique(dbPool *pgxpool.Pool, user User) (bool, error) {
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
func CreateUser(dbPool *pgxpool.Pool, user User) error {
	salt, err := uuid.NewV4()
	if err != nil {
		return err
	}

	hashPassword, err := utils.HashAndSalt(user.Password, salt.String())
	if err != nil {
		return err
	}

	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}

	sql := "INSERT INTO users(username, email, password, salt) VALUES($1, $2, $3, $4);"
	conn.QueryRow(context.Background(), sql, user.Name, user.Email, hashPassword, salt)

	return nil
}
