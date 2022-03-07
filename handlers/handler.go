package handlers

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
	"myapp/models"
	"net/http"

	_ "myapp/docs"
)

// CreateUserHandler godoc
// @Summary Creates new user.
// @Description Create new user in database with validation.
// @Tags Signup
// @Accept */*
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /signup [post]
func CreateUserHandler(dbPool *pgxpool.Pool) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := new(models.User)
		if err := context.Bind(user); err != nil {
			return err
		}

		if err := models.ValidateUser(user); err != nil {
			return context.JSON(http.StatusBadRequest, err)
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

		return context.JSON(http.StatusOK, "OK: User can be registered")
	}
}

// GetHomePageHandler godoc
// @Summary Get Home Page.
// @Description Get your home page.
// @Tags GetHomePage
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func GetHomePageHandler() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Test: homePageHandler")
	}
}
