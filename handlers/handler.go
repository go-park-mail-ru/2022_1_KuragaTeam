package handlers

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/gofrs/uuid"
	"log"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/labstack/echo"
)

func CreateUserHandler(dbPool *pgxpool.Pool, connRedis *redis.Conn) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := new(models.User)
		if err := context.Bind(&user); err != nil {
			return err
		}

		isUnique, err := utils.IsUserUnique(dbPool, *user)
		if err != nil {
			return err
		}

		if !isUnique {
			return context.JSON(http.StatusBadRequest, "ERROR: Email is not unique")
		}

		userID, err := utils.CreateUser(dbPool, *user)
		if err != nil {
			return err
		}

		value, err := uuid.NewV4()
		if err != nil {
			return err
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    value.String(),
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
		}

		context.SetCookie(&cookie)

		_, err = (*connRedis).Do("SET", value, userID, "EX", int64(time.Hour.Seconds()))
		if err != nil {
			return err
		}

		return context.JSON(http.StatusCreated, "OK: User created")
	}
}

func LoginUserHandler(dbPool *pgxpool.Pool, connRedis *redis.Conn) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := new(models.User)
		if err := context.Bind(user); err != nil {
			return err
		}

		userID, userExists, err := utils.IsUserExists(dbPool, *user)
		if err != nil {
			if errors.Is(err, utils.ErrWrongPassword) {
				return context.JSON(http.StatusUnauthorized, "ERROR: Wrong password")
			}

			return err
		}

		if !userExists {
			return context.JSON(http.StatusNotFound, "ERROR: User not found")
		}

		value, err := uuid.NewV4()
		if err != nil {
			return err
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    value.String(),
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
		}

		context.SetCookie(&cookie)

		log.Println(connRedis, value, userID)
		_, err = (*connRedis).Do("SET", value, userID, "EX", int64(time.Hour.Seconds()))
		if err != nil {
			return err
		}

		return context.JSON(http.StatusOK, "OK: User can be logged in")
	}
}

func GetHomePageHandler() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Test: homePageHandler")
	}
}
