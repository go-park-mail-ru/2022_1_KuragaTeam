package models

import (
	"context"
	"errors"
	"gopkg.in/validator.v2"
	"myapp/utils"
	"strings"
	"unicode"

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

func ValidateUser(user *User) error {
	user.Name = strings.TrimSpace(user.Name)
	if err := validator.Validate(user); err != nil {
		return err
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}
	return nil
}

func ValidatePassword(pass string) error {
	var (
		up, low, num bool
		symbolsCount uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			up = true
			symbolsCount++
		case unicode.IsLower(char):
			low = true
			symbolsCount++
		case unicode.IsNumber(char):
			num = true
			symbolsCount++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			//sym = true
			symbolsCount++
		default:
			return banErr
		}
	}

	if !up {
		return upErr
	}
	if !low {
		return lowErr
	}
	if !num {
		return numErr
	}
	if symbolsCount < 8 {
		return countErr
	}

	return nil
}
