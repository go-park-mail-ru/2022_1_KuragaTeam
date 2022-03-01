package models

import (
	"database/sql"
	"log"
	"myapp/utils"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Используется LoginUserHandler.
//Проверяет, что пользователь есть в базе данных.
func IsUserExists(db *sql.DB, user User) (bool, error) {
	sql := "SELECT username, password FROM USERS"
	rows, err := db.Query(sql)
	// выходим, если SQL не сработал по каким-то причинам
	if err != nil {
		return false, err
	}
	// убедимся, что всё закроется при выходе из программы
	defer rows.Close()

	// Из базы пришел пустой запрос, значит пользователя в базе данных нет
	if !rows.Next() {
		return false, nil
	}

	var u User
	err = rows.Scan(&u.Name, &u.Password)

	// выход при ошибке
	if err != nil {
		return false, err
	}

	result, err := utils.ComparePasswords(u.Password, []byte(user.Password))
	if err != nil {
		return false, err
	}

	return result, nil
}

//Используется CreateUserHandler.
//email должен быть уникален
func IsUserUnique(db *sql.DB, user User) (bool, error) {
	sql := "SELECT * FROM users WHERE email=$1"
	rows, err := db.Query(sql, user.Email)
	if err != nil {
		return false, err
	}
	if rows.Next() { // Пользователь с таким email зарегистрирован
		return false, nil
	}
	return true, nil
}

//Используется CreateUserHandler.
//Создает пользователя
func CreateUser(db *sql.DB, user User) error {
	hashPassword, err := utils.HashAndSalt([]byte(user.Password))
	if err != nil {
		return err
	}

	log.Println(user)

	sql := "INSERT INTO users(username, email, password) VALUES($1, $2, $3)"
	_, err = db.Exec(sql, user.Name, user.Email, hashPassword)
	if err != nil {
		return err
	}

	return nil
}
