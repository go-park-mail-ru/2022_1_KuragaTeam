package handlers

import (
	"errors"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"

	_ "myapp/docs"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseName struct {
	Status int    `json:"status"`
	Name   string `json:"username"`
}

type ResponseMovieCompilations struct {
	Status           int                       `json:"status"`
	MovieCompilation []models.MovieCompilation `json:"moviesCompilation"`
}

// CreateUserHandler godoc
// @Summary Creates new user.
// @Description Create new user in database with validation.
// @Produce json
// @Param data body models.User true "Data for user"
// @Success 	201 {object} Response "OK: User created"
// @Failure		400 {object} Response "Invalid request body"
// @Failure		500 {object} Response "Internal server error"
// @Router /signup [post]
func CreateUserHandler(dbPool *utils.UserPool, redisPool *redis.Pool) echo.HandlerFunc {
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

		isUnique, err := dbPool.IsUserUnique(user)
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

		userID, err := dbPool.CreateUser(user)
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
			SameSite: 0,
		}

		context.SetCookie(cookie)

		connRedis := redisPool.Get()
		defer connRedis.Close()
		_, err = connRedis.Do("SET", value, userID, "EX", int64(time.Hour.Seconds()))
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

// LoginUserHandler godoc
// @Summary Login in account.
// @Description Check login and gives session ID.
// @Produce json
// @Param data body models.User true "Data for user"
// @Success 	200 {object} Response "Successful login"
// @Failure		400 {object} Response "Invalid request body"
// @Failure		401 {object} Response "Wrong password"
// @Failure		404 {object} Response "User not found"
// @Failure		500 {object} Response "Internal server error"
// @Router /login [post]
func LoginUserHandler(dbPool *utils.UserPool, redisPool *redis.Pool) echo.HandlerFunc {
	return func(context echo.Context) error {
		user := models.User{}
		if err := context.Bind(&user); err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		userID, userExists, err := dbPool.IsUserExists(user)
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
			SameSite: 0,
		}

		context.SetCookie(&cookie)

		connRedis := redisPool.Get()
		defer connRedis.Close()
		_, err = connRedis.Do("SET", value, userID, "EX", int64(time.Hour.Seconds()))
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
// @Produce json
// @Success 	200 {object} Response models.User.name
// @Failure		401 {object} Response "ERROR: User is unauthorized"
// @Failure		500 {object} Response "Internal server error"
// @Router / [get]
func GetHomePageHandler(dbPool *utils.UserPool) echo.HandlerFunc {
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

		name, err := dbPool.GetUserName(userID)
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

// LogoutHandler godoc
// @Summary Logout.
// @Description Delete session from DB.
// @Produce json
// @Success 	200 {object} Response "OK: User is logged out"
// @Failure		500 {object} Response "Internal server error"
// @Router /logout [delete]
func LogoutHandler(redisPool *redis.Pool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		connRedis := redisPool.Get()
		defer connRedis.Close()
		_, err = connRedis.Do("DEL", cookie.Value)
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

// GetMovieCompilations godoc
// @Summary Get Movie Compilations.
// @Description Get movie compilations for user.
// @Produce json
// @Success 	200 {object} Response models.MovieCompilation
// @Failure		401 {object} Response "ERROR: User is unauthorized"
// @Failure		500 {object} Response "Internal server error"
// @Router 		/movieCompilations [get]
func GetMovieCompilations() echo.HandlerFunc {
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

		movieCompilations := []models.MovieCompilation{
			{
				Name: "Популярное",
				Movies: []models.Movie{
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны1",
						Genre: "Фантастика1",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны2",
						Genre: "Фантастика2",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны3",
						Genre: "Фантастика3",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны4",
						Genre: "Фантастика4",
					},
				},
			},
			{
				Name: "Топ",
				Movies: []models.Movie{
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны#1",
						Genre: "Фантастика",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны#2",
						Genre: "Фантастика",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны#3",
						Genre: "Фантастика",
					},
				},
			},
			{
				Name: "Семейное",
				Movies: []models.Movie{
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны#1",
						Genre: "Фантастика",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны#2",
						Genre: "Фантастика",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны#3",
						Genre: "Фантастика",
					},
					{
						Img:   "star.png",
						Href:  "/",
						Name:  "Звездные войны4",
						Genre: "Фантастика4",
					},
				},
			},
		}

		return context.JSON(http.StatusOK, &ResponseMovieCompilations{
			Status:           http.StatusOK,
			MovieCompilation: movieCompilations,
		})

	}
}
