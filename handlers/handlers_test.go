package handlers

import (
	"encoding/json"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestCase struct {
	name       string
	Response   string
	StatusCode int
}

//func TestGetHomePage(t *testing.T) {
//t.Parallel()
//ctrl := gomock.NewController(t)
//defer ctrl.Finish()
//
//// given
//mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
//columns := []string{"id", "email", "password", "salt"}
//
//salt, _ := uuid.NewV4()
//pwd := "Pass123321"
//password, _ := utils.HashAndSalt(pwd, salt.String())
//
//pgxRows := pgxpoolmock.NewRows(columns).AddRow(int64(1), "Ilias@mail.ru", password, salt.String()).ToPgxRows()
//
//user := models.User{
//	ID:       1,
//	Name:     "Ilias",
//	Email:    "Ilias@mail.ru",
//	Password: pwd,
//	Salt:     salt.String(),
//}
//mockPool.EXPECT().Query(gomock.Any(), `SELECT id, email, password, salt FROM users WHERE email=$1`, user.Email).Return(pgxRows, nil)
//
//// when
////actualOrder := orderDao.GetOrderByID(1)
//
//userPool := &utils.UserPool{
//	Pool: mockPool,
//}
//
//userID, result, err := userPool.IsUserExists(user)
//
//// then
//assert.Equal(t, int64(1), userID)
//assert.Equal(t, true, result)
//assert.Nil(t, err)
//
//cases := []TestCase{
//	TestCase{
//		name:       "test1",
//		Response:   `"Test: homePageHandler"`,
//		StatusCode: http.StatusOK,
//	},
//}
//
//for _, item := range cases {
//
//	server := echo.New()
//	response := httptest.NewRequest(echo.GET, "/", nil)
//	rec := httptest.NewRecorder()
//	ctx := server.NewContext(response, rec)
//
//	resp := GetHomePageHandler()
//
//	if assert.NoError(t, resp(ctx)) {
//		assert.Equal(t, item.StatusCode, rec.Code)
//
//		body, _ := ioutil.ReadAll(rec.Result().Body)
//		bodyStr := string(body)
//		bodyStr = strings.TrimSpace(bodyStr)
//
//		assert.Equal(t, item.Response, bodyStr)
//	}
//
//}
//}

func TestCreateUser(t *testing.T) {
	//cases := []TestLoginCase{
	//	TestLoginCase{
	//		name:       "test1",
	//		Response:   `"OK: User created"`,
	//		StatusCode: http.StatusCreated,
	//		response: models.User{
	//			ID:       115,
	//			Name:     "Lol1",
	//			Password: "Lol1Password",
	//			Email:    "lol@lol.com",
	//		},
	//	},
	//}
	//
	//dbPool, err := db.ConnectDB()
	//if err != nil {
	//	return
	//}
	//
	//for _, item := range cases {
	//
	//	server := echo.New()
	//	d, _ := json.Marshal(item.response)
	//	response := httptest.NewRequest(echo.POST, "/signup", bytes.NewBuffer(d))
	//	rec := httptest.NewRecorder()
	//	response.Header.Set("Content-Type", "application/json")
	//	ctx := server.NewContext(response, rec)
	//
	//	resp := CreateUserHandler(dbPool)
	//
	//	if assert.NoError(t, resp(ctx)) {
	//		assert.Equal(t, item.StatusCode, rec.Code)
	//
	//		body, _ := ioutil.ReadAll(rec.Result().Body)
	//		bodyStr := string(body)
	//		bodyStr = strings.TrimSpace(bodyStr)
	//
	//		assert.Equal(t, item.Response, bodyStr)
	//	}
	//	_ = server.Close()
	//}
	//dbPool.Close()
}

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
		var receive Response
		err := json.Unmarshal(body, &receive)
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
		req, _ := json.Marshal(movieCompilations)
		reqStr := string(req)
		reqStr = strings.TrimSpace(reqStr)

		assert.Equal(t, reqStr, bodyStr)
	}
}

type TestGetHomePage struct {
	name       string
	StatusCode int
	Response   ResponseName
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
		},
	}

	server := echo.New()
	defer func(server *echo.Echo) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	for _, item := range cases {

		req := httptest.NewRequest(echo.DELETE, "/api/v1/logout", nil)
		rec := httptest.NewRecorder()
		ctx := server.NewContext(req, rec)
		ctx.Set("USER_ID", int64(1))

		//cmd := conn.Command("SET", item.cookie.Value).ExpectMap(map[string]string{
		//	"": "",
		//})
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
		columns := []string{"username"}

		pgxRows := pgxpoolmock.NewRows(columns).AddRow(item.Response.Name).ToPgxRows()

		user := models.User{
			ID:    int64(1),
			Name:  item.Response.Name,
			Email: "Ivan@mail.ru",
		}

		expectedResult := ""

		mockPool.EXPECT().QueryRow(gomock.Any(), `SELECT username FROM users WHERE id=$1`, user.ID).Return(pgxRows)

		if pgxRows.Next() {
			err := pgxRows.Scan(&expectedResult)
			if err != nil {
				log.Fatal(err)
			}
		}
		userPool := &utils.UserPool{
			Pool: mockPool,
		}

		resp := GetHomePageHandler(userPool)

		if assert.NoError(t, resp(ctx)) {
			//assert.Equal(t, conn.Stats(cmd), 1)
		}
		assert.Equal(t, rec.Code, item.StatusCode)
		body, _ := ioutil.ReadAll(rec.Result().Body)
		var receive ResponseName
		err := json.Unmarshal(body, &receive)
		assert.NoError(t, err)
		assert.Equal(t, item.Response, receive)
	}

}
