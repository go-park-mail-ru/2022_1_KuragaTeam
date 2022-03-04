package handlers

import (
	"errors"
	"fmt"
	"gopkg.in/validator.v2"
	"myapp/models"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/labstack/echo"
)

func CreateUserHandler(dbPool *pgxpool.Pool) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := new(models.User)

		if err := context.Bind(user); err != nil {
			return err
		}

		user.Name = strings.TrimSpace(user.Name)
		if err := validator.Validate(user); err != nil {
			// values not valid, deal with errors here
			fmt.Println(err)
			return context.JSON(http.StatusBadRequest, err)
			//return err
		}

		isUnique, err := models.IsUserUnique(dbPool, *user)
		if err != nil {
			return err
		}

		if !isUnique {
			return context.JSON(http.StatusBadRequest, "ERROR: Email is not unique")
		}

		err = models.CreateUser(dbPool, *user)
		if err != nil {
			return err
		}

		return context.JSON(http.StatusCreated, "OK: User created")
	}
}

func LoginUserHandler(dbPool *pgxpool.Pool) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := new(models.User)
		if err := context.Bind(user); err != nil {
			return err
		}

		userExists, err := models.IsUserExists(dbPool, *user)
		if err != nil {
			if errors.Is(err, models.ErrWrongPassword) {
				return context.JSON(http.StatusUnauthorized, "ERROR: Wrong password")
			}

			return err
		}

		if !userExists {
			return context.JSON(http.StatusNotFound, "ERROR: User not found")
		}

		return context.JSON(http.StatusOK, "OK: User can be logined in")
	}
}

func GetHomePageHandler() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Test: homePageHandler")
	}
}
