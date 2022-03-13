package handlers

import (
	"encoding/json"
	"io/ioutil"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

type TestLogoutCase struct {
	name       string
	StatusCode int
	cookie     *http.Cookie
	response   Response
}

func TestLogoutHandler(t *testing.T) {
	cases := []TestLogoutCase{
		TestLogoutCase{
			name:       "Default logout",
			StatusCode: http.StatusOK,
			cookie: &http.Cookie{
				Name:     "Session_cookie",
				Value:    "okdg0wrijifoi",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
				Secure:   false,
			},
			response: Response{
				Status:  http.StatusOK,
				Message: "OK: User is logged out",
			},
		},
		TestLogoutCase{
			name:       "Logout without cookie",
			StatusCode: http.StatusInternalServerError,
			cookie: &http.Cookie{
				Name:     "Cookie",
				Value:    "okdg0wrijifoi",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
				Secure:   false,
			},
			response: Response{
				Status:  http.StatusInternalServerError,
				Message: "http: named cookie not present",
			},
		},
	}

	server := echo.New()
	defer func(server *echo.Echo) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:        func() (redis.Conn, error) { return conn, nil },
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
	}

	for _, item := range cases {

		req := httptest.NewRequest(echo.DELETE, "/api/v1/logout", nil)
		rec := httptest.NewRecorder()
		ctx := server.NewContext(req, rec)
		ctx.Request().AddCookie(item.cookie)

		cmd := conn.Command("DEL", item.cookie.Value).ExpectMap(map[string]string{
			"": "",
		})

		resp := LogoutHandler(pool)

		if assert.NoError(t, resp(ctx)) {
			assert.Equal(t, conn.Stats(cmd), 1)
		}
		assert.Equal(t, item.StatusCode, rec.Code)

		body, _ := ioutil.ReadAll(rec.Result().Body)
		err := rec.Result().Body.Close()
		assert.NoError(t, err)

		var receive Response
		err = json.Unmarshal(body, &receive)
		assert.NoError(t, err)

		assert.Equal(t, item.response, receive)
	}

}

func TestGetMovieCompilations(t *testing.T) {
	server := echo.New()
	response := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations", nil)
	rec := httptest.NewRecorder()
	ctx := server.NewContext(response, rec)
	ctx.Set("USER_ID", int64(1))

	resp := GetMovieCompilations()

	if assert.NoError(t, resp(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		body, _ := ioutil.ReadAll(rec.Result().Body)
		err := rec.Result().Body.Close()
		assert.NoError(t, err)

		movies := []models.MovieCompilation{
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
		movieCompilations := ResponseMovieCompilations{
			Status:           http.StatusOK,
			MovieCompilation: movies,
		}
		bodyStr := string(body)
		bodyStr = strings.TrimSpace(bodyStr)
		req, err := json.Marshal(movieCompilations)
		assert.NoError(t, err)
		reqStr := string(req)
		reqStr = strings.TrimSpace(reqStr)

		assert.Equal(t, reqStr, bodyStr)
	}
}

type TestGetHomePage struct {
	name       string
	StatusCode int
	Response   ResponseName
	UserID     int64
	UserIDKey  string
}

func TestGetHomePageHandler(t *testing.T) {
	cases := []TestGetHomePage{
		TestGetHomePage{
			name:       "Default login case",
			StatusCode: http.StatusOK,
			Response: ResponseName{
				Status: http.StatusOK,
				Name:   "user1",
			},
			UserID:    1,
			UserIDKey: "USER_ID",
		},
		TestGetHomePage{
			name:       "User unauthorized",
			StatusCode: http.StatusUnauthorized,
			Response: ResponseName{
				Status: http.StatusUnauthorized,
				Name:   "",
			},
			UserID:    -1,
			UserIDKey: "USER_ID",
		},
		TestGetHomePage{
			name:       "No user_id",
			StatusCode: http.StatusInternalServerError,
			Response: ResponseName{
				Status: http.StatusInternalServerError,
				Name:   "",
			},
			UserID:    5,
			UserIDKey: "ID",
		},
		//TestGetHomePage{
		//	name:       "Wrong username",
		//	StatusCode: http.StatusOK,
		//	Response: ResponseName{
		//		Status: http.StatusOK,
		//		Name:   "user2",
		//	},
		//},
	}

	server := echo.New()
	defer func(server *echo.Echo) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	ctrl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"username"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow("user1").ToPgxRows()

	for _, item := range cases {

		req := httptest.NewRequest(echo.GET, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		ctx := server.NewContext(req, rec)
		ctx.Set(item.UserIDKey, item.UserID)

		defer ctrl.Finish()

		user := models.User{
			ID:   int64(1),
			Name: item.Response.Name,
		}

		expectedResult := ""
		if item.StatusCode == http.StatusOK {
			mockPool.EXPECT().QueryRow(gomock.Any(), `SELECT username FROM users WHERE id=$1`, user.ID).Return(pgxRows)
		}

		if pgxRows.Next() {
			err := pgxRows.Scan(&expectedResult)
			assert.NoError(t, err)
		}
		userPool := &utils.UserPool{
			Pool: mockPool,
		}

		resp := GetHomePageHandler(userPool)

		assert.NoError(t, resp(ctx))

		assert.Equal(t, rec.Code, item.StatusCode)

		body, _ := ioutil.ReadAll(rec.Result().Body)
		err := rec.Result().Body.Close()
		assert.NoError(t, err)

		var receive ResponseName
		err = json.Unmarshal(body, &receive)
		assert.NoError(t, err)

		assert.Equal(t, item.Response, receive)
	}

}
