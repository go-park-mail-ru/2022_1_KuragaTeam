package handlers

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gofrs/uuid"

	"github.com/labstack/echo/v4"
	_ "myapp/docs"
)

// CreateUserHandler godoc
// @Summary Creates new user.
// @Description Create new user in database with validation.
// @Tags Signup
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param email formData string true "email"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /signup [post]
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseName struct {
	Status int    `json:"status"`
	Name   string `json:"username"`
}

func CreateUserHandler(dbPool *pgxpool.Pool, connRedis *redis.Conn) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := models.User{}

		if err := context.Bind(&user); err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if err := utils.ValidateUser(&user); err != nil {
			return context.JSON(http.StatusBadRequest, &Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		isUnique, err := utils.IsUserUnique(dbPool, user)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})

		}

		if !isUnique {
			return context.JSON(http.StatusBadRequest, &Response{
				Status:  http.StatusBadRequest,
				Message: "ERROR: Email is not unique",
			})
		}

		userID, err := utils.CreateUser(dbPool, user)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		value, err := uuid.NewV4()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		cookie := &http.Cookie{
			Name:     "Session_cookie",
			Value:    value.String(),
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
		}

		context.SetCookie(cookie)

		_, err = (*connRedis).Do("SET", value, userID, "EX", int64(time.Hour.Seconds()))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return context.JSON(http.StatusCreated, &Response{
			Status:  http.StatusCreated,
			Message: "OK: User created",
		})
	}
}

func LoginUserHandler(dbPool *pgxpool.Pool, connRedis *redis.Conn) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := models.User{}
		if err := context.Bind(&user); err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		userID, userExists, err := utils.IsUserExists(dbPool, user)
		if err != nil {
			if errors.Is(err, utils.ErrWrongPassword) {
				return context.JSON(http.StatusUnauthorized, &Response{
					Status:  http.StatusUnauthorized,
					Message: "ERROR: Wrong password",
				})
			}

			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if !userExists {
			return context.JSON(http.StatusNotFound, &Response{
				Status:  http.StatusNotFound,
				Message: "ERROR: User not found",
			})
		}

		value, err := uuid.NewV4()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
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
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return context.JSON(http.StatusOK, &Response{
			Status:  http.StatusOK,
			Message: "OK: User can be logged in",
		})
	}
}

// GetHomePageHandler godoc
// @Summary Get Home Page.
// @Description Get your home page.
// @Tags GetHomePage
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func GetHomePageHandler(dbPool *pgxpool.Pool) echo.HandlerFunc {
	return func(context echo.Context) error {
		userID, ok := context.Get("USER_ID").(int64)
		if !ok {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: "ERROR: Session required",
			})
		}

		if userID == -1 {
			return context.JSON(http.StatusUnauthorized, &Response{
				Status:  http.StatusUnauthorized,
				Message: "ERROR: User is unauthorized",
			})
		}

		name, err := utils.GetUserName(dbPool, userID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return context.JSON(http.StatusOK, &ResponseName{
			Status: http.StatusOK,
			Name:   name,
		})
	}
}

func LogoutHandler(connRedis *redis.Conn) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		_, err = (*connRedis).Do("DEL", cookie.Value)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusOK, &Response{
			Status:  http.StatusOK,
			Message: "OK: User is logged out",
		})
	}
}
