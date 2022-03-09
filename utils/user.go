package utils

import (
	"context"
	"myapp/models"
	"strings"
	"unicode"

	"github.com/driftprogramming/pgxpoolmock"
	"gopkg.in/validator.v2"

	"github.com/gofrs/uuid"
)

type UserPool struct {
	Pool pgxpoolmock.PgxPool
}

// Используется LoginUserHandler.
// Проверяет, что пользователь есть в базе данных.
func (dbPool *UserPool) IsUserExists(user models.User) (int64, bool, error) {
	var userID int64
	sql := "SELECT id, email, password, salt FROM users WHERE email=$1"
	rows, err := dbPool.Pool.Query(context.Background(), sql, user.Email)
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

//// Используется LoginUserHandler.
//// Проверяет, что пользователь есть в базе данных.
//func IsUserExists(dbPool *UserPool, user models.User) (int64, bool, error) {
//	var userID int64
//	sql := "SELECT id, email, password, salt FROM USERS WHERE email=$1"
//	rows, err := dbPool.Pool.Query(context.Background(), sql, user.Email)
//	if err != nil {
//		return userID, false, err
//	}
//
//	// убедимся, что всё закроется при выходе из программы
//	defer func() {
//		rows.Close()
//	}()
//
//	// Из базы пришел пустой запрос, значит пользователя в базе данных нет
//	if !rows.Next() {
//		return userID, false, nil
//	}
//
//	var signInUser models.User
//	err = rows.Scan(&signInUser.ID, &signInUser.Email, &signInUser.Password, &signInUser.Salt)
//
//	userID = signInUser.ID
//	// выход при ошибке
//	if err != nil {
//		return userID, false, err
//	}
//
//	result, err := ComparePasswords(signInUser.Password, signInUser.Salt, user.Password)
//	if err != nil {
//		return userID, false, ErrWrongPassword
//	}
//
//	result = result && signInUser.Email == user.Email
//
//	return userID, result, nil
//}

// Используется CreateUserHandler.
// email должен быть уникален
func (dbPool *UserPool) IsUserUnique(user models.User) (bool, error) {
	sql := "SELECT * FROM users WHERE email=$1;"
	rows, err := dbPool.Pool.Query(context.Background(), sql, user.Email)

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

//func IsUserUnique(dbPool *UserPool, user models.User) (bool, error) {
//	sql := "SELECT * FROM users WHERE email=$1;"
//	rows, err := dbPool.Pool.Query(context.Background(), sql, user.Email)
//
//	if err != nil {
//		return false, err
//	}
//
//	defer func() {
//		rows.Close()
//	}()
//
//	if rows.Next() { // Пользователь с таким email зарегистрирован
//		return false, nil
//	}
//	return true, nil
//}

// Используется CreateUserHandler.
// Создает пользователя
func (dbPool *UserPool) CreateUser(user models.User) (int64, error) {
	var userID int64

	salt, err := uuid.NewV4()
	if err != nil {
		return userID, err
	}

	hashPassword, err := HashAndSalt(user.Password, salt.String())
	if err != nil {
		return userID, err
	}

	sql := "INSERT INTO users(username, email, password, salt) VALUES($1, $2, $3, $4) RETURNING id;"
	if err = dbPool.Pool.QueryRow(context.Background(), sql, user.Name, user.Email, hashPassword, salt).Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

//func CreateUser(dbPool *UserPool, user models.User) (int64, error) {
//	var userID int64
//
//	salt, err := uuid.NewV4()
//	if err != nil {
//		return userID, err
//	}
//
//	hashPassword, err := HashAndSalt(user.Password, salt.String())
//	if err != nil {
//		return userID, err
//	}
//
//	sql := "INSERT INTO users(username, email, password, salt) VALUES($1, $2, $3, $4) RETURNING id;"
//	if err = dbPool.Pool.QueryRow(context.Background(), sql, user.Name, user.Email, hashPassword, salt).Scan(&userID); err != nil {
//		return userID, err
//	}
//
//	return userID, nil
//}

func (dbPool *UserPool) GetUserName(userID int64) (string, error) {
	sql := "SELECT username FROM users WHERE id=$1;"

	var name string
	err := dbPool.Pool.QueryRow(context.Background(), sql, userID).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

//
//func GetUserName(dbPool *UserPool, userID int64) (string, error) {
//	sql := "SELECT username FROM users WHERE id=$1;"
//
//	var name string
//	err := dbPool.Pool.QueryRow(context.Background(), sql, userID).Scan(&name)
//	if err != nil {
//		return "", err
//	}
//
//	return name, nil
//}

func ValidateUser(user *models.User) error {
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
