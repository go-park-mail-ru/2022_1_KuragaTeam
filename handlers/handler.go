package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"myapp/models"
	"net/http"
)

func CreateUserHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return err
		}

		isUnique, err := models.IsUserUnique(db, *user)
		if err != nil {
			return err
		}

		if !isUnique {
			return c.JSON(http.StatusBadRequest, "ERROR: Email is not unique")
		}

		err = models.CreateUser(db, *user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, "OK: User created")
	}
}

func LoginUserHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return err
		}

		userExists, err := models.IsUserExists(db, *user)
		if err != nil {
			return err
		}

		if !userExists {
			return c.JSON(http.StatusNotFound, "ERROR: User not found")
		}

		return c.JSON(http.StatusOK, "OK: User can be registered")
	}
}

func GetHomePageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Test: homePageHandler")
	}
}
