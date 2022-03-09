package handlers

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"myapp/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestCase struct {
	name       string
	Response   string
	StatusCode int
}

func TestGetHomePage(t *testing.T) {
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
}

type TestLoginCase struct {
	name       string
	Response   string
	StatusCode int
	req        models.User
}

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
			name:       "Logout",
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
	}

	server := echo.New()
	defer func(server *echo.Echo) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	req := httptest.NewRequest(echo.DELETE, "/api/v1/logout", nil)
	rec := httptest.NewRecorder()
	ctx := server.NewContext(req, rec)
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:        func() (redis.Conn, error) { return conn, nil },
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
	}

	for _, item := range cases {

		ctx.Request().AddCookie(item.cookie)

		cmd := conn.Command("DEL", item.cookie.Value).ExpectMap(map[string]string{
			"": "",
		})

		resp := LogoutHandler(pool)

		if assert.NoError(t, resp(ctx)) {

			assert.Equal(t, conn.Stats(cmd), 1)

			assert.Equal(t, item.StatusCode, rec.Code)

			body, _ := ioutil.ReadAll(rec.Result().Body)

			var receive Response
			err := json.Unmarshal(body, &receive)
			assert.NoError(t, err)
			assert.Equal(t, item.response, receive)
		}
	}

}
