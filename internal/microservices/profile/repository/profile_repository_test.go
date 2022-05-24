package repository

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"myapp/internal/constants"
	"myapp/internal/microservices/profile/proto"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gomodule/redigo/redis"
	"github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestProfileRepository_GetUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		id          int64
		expected    *proto.ProfileData
		expectedErr error
	}{
		{
			name: "Get user profile",
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "email", "avatar", "subscription_expires"}).AddRow("Ilias", "Ilias@mail.ru", constants.DefaultImage, "2022-05-18T22:26:17.289395Z")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, email, avatar, subscription_expires FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id: int64(1),
			expected: &proto.ProfileData{
				Name:   "Ilias",
				Email:  "Ilias@mail.ru",
				Avatar: "https://" + os.Getenv("HOST") + "/api/v1/minio/avatars/default_avatar.webp",
				Date:   "2022-05-18T22:26:17.289395Z",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, email, avatar, subscription_expires FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			id:          int64(1),
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			selectedUser, err := storage.GetUserProfile(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, selectedUser)
			}
		})
	}
}

func TestProfileRepository_EditProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		user        *proto.EditProfileData
		expectedErr error
	}{
		{
			name: "Edit user profile, change username",
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "password", "salt"}).AddRow("Ilias", "pass123123", "salt")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, password, salt FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET username = $2 WHERE id = $1`)).
					WithArgs(
						driver.Value(0),
						driver.Value("Ivan"),
					).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user: &proto.EditProfileData{
				Name: "Ivan",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, password, salt FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnError(errors.New("Error occurred during request "))
			},
			user: &proto.EditProfileData{
				Name:     "Ilias",
				Password: "pass123321",
			},
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Error occurred during UPDATE",
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "password", "salt"}).AddRow("Ilias", "pass123123", "salt")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, password, salt FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET username = $2 WHERE id = $1`)).
					WithArgs(
						driver.Value(0),
						driver.Value("Ivan"),
					).WillReturnError(errors.New("Error occurred during request ")).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user: &proto.EditProfileData{
				Name: "Ivan",
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.EditProfile(th.user)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_EditAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name              string
		mock              func()
		user              *proto.EditAvatarData
		expectedOldAvatar string
		expectedErr       error
	}{
		{
			name: "Edit user avatar",
			mock: func() {
				rows := sqlmock.NewRows([]string{"avatar"}).AddRow("avatar")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET avatar = $2 WHERE id = $1`)).
					WithArgs(
						driver.Value(0),
						driver.Value("new_avatar"),
					).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user: &proto.EditAvatarData{
				Avatar: "new_avatar",
			},
			expectedOldAvatar: "avatar",
			expectedErr:       nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnError(errors.New("Error occurred during request "))
			},
			user: &proto.EditAvatarData{
				Avatar: "new_avatar",
			},
			expectedOldAvatar: "",
			expectedErr:       errors.New("Error occurred during request "),
		},
		{
			name: "Error occurred during UPDATE",
			mock: func() {
				rows := sqlmock.NewRows([]string{"avatar"}).AddRow("avatar")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET avatar = $2 WHERE id = $1`)).
					WithArgs(
						driver.Value(0),
						driver.Value("new_avatar"),
					).WillReturnError(errors.New("Error occurred during request ")).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user: &proto.EditAvatarData{
				Avatar: "new_avatar",
			},
			expectedOldAvatar: "",
			expectedErr:       errors.New("Error occurred during request "),
		},
		{
			name: "Empty avatar",
			mock: func() {
				rows := sqlmock.NewRows([]string{"avatar"}).AddRow("avatar")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(0)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET avatar = $2 WHERE id = $1`)).
					WithArgs(
						driver.Value(0),
						driver.Value(""),
					).WillReturnError(nil).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user:              &proto.EditAvatarData{},
			expectedOldAvatar: "",
			expectedErr:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			oldAvatar, err := storage.EditAvatar(th.user)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expectedOldAvatar, oldAvatar)
			}
		})
	}
}

func TestProfileRepository_GetAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		id          int64
		expected    string
		expectedErr error
	}{
		{
			name: "Get avatar",
			mock: func() {
				rows := sqlmock.NewRows([]string{"avatar"}).AddRow(constants.DefaultImage)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id:          int64(1),
			expected:    "default_avatar.webp",
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			id:          int64(1),
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			avatar, err := storage.GetAvatar(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, avatar)
			}
		})
	}
}

func TestProfileRepository_AddLike(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		data        *proto.LikeData
		expectedErr error
	}{
		{
			name: "Add like",
			mock: func() {
				likes := make([]int, 0)
				likes = append(likes, 1, 2, 3)
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET likes = array_append(likes, $2) WHERE id=$1`)).
					WithArgs(
						driver.Value(1),
						driver.Value(4),
					).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 4,
			},
			expectedErr: nil,
		},
		{
			name: "Like exists",
			mock: func() {
				likes := make([]int, 0)
				likes = append(likes, 1, 2, 3)
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 3,
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 4,
			},
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Error occurred during UPDATE request",
			mock: func() {
				likes := make([]int, 0)
				likes = append(likes, 1, 2, 3)
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET likes = array_append(likes, $2) WHERE id=$1`)).
					WithArgs(
						driver.Value(1),
						driver.Value(4),
					).WillReturnError(errors.New("Error occurred during request "))
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 4,
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.AddLike(th.data)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_RemoveLike(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		data        *proto.LikeData
		expectedErr error
	}{
		{
			name: "Remove like",
			mock: func() {
				likes := make([]int, 0)
				likes = append(likes, 1, 2, 3)
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET likes = array_remove(likes, $2) WHERE id=$1`)).
					WithArgs(
						driver.Value(1),
						driver.Value(3),
					).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 3,
			},
			expectedErr: nil,
		},
		{
			name: "Like not exists",
			mock: func() {
				likes := make([]int, 0)
				likes = append(likes, 1, 2, 3)
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 4,
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 4,
			},
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Error occurred during UPDATE request",
			mock: func() {
				likes := make([]int, 0)
				likes = append(likes, 1, 2, 3)
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET likes = array_append(likes, $2) WHERE id=$1`)).
					WithArgs(
						driver.Value(1),
						driver.Value(3),
					).WillReturnError(errors.New("Error occurred during request "))
			},
			data: &proto.LikeData{
				UserID:  1,
				MovieID: 3,
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.RemoveLike(th.data)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_GetFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	likes := make([]int64, 0)
	likes = append(likes, 1, 2, 3)

	tests := []struct {
		name        string
		mock        func()
		id          int64
		expected    *proto.Favorites
		expectedErr error
	}{
		{
			name: "Get favorites",
			mock: func() {
				rows := sqlmock.NewRows([]string{"likes"}).AddRow(pq.Array(likes))
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id:          int64(1),
			expected:    &proto.Favorites{Id: likes},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT likes FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			id:          int64(1),
			expected:    &proto.Favorites{},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			avatar, err := storage.GetFavorites(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, avatar)
			}
		})
	}
}

func TestProfileRepository_GetRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	likes := make([]int64, 0)
	likes = append(likes, 1, 2, 3)

	tests := []struct {
		name        string
		mock        func()
		data        *proto.MovieRating
		expected    *proto.Rating
		expectedErr error
	}{
		{
			name: "Get rating",
			mock: func() {
				rows := sqlmock.NewRows([]string{"rating"}).AddRow(5)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT rating FROM rating WHERE user_id=$1 AND movie_id = $2`)).
					WithArgs(driver.Value(1), driver.Value(1)).WillReturnRows(rows)
			},
			data: &proto.MovieRating{
				UserID:  1,
				MovieID: 1,
			},
			expected:    &proto.Rating{Rating: 5},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT rating FROM rating WHERE user_id=$1 AND movie_id = $2`)).
					WithArgs(driver.Value(1), driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			data: &proto.MovieRating{
				UserID:  1,
				MovieID: 1,
			},
			expected:    &proto.Rating{},
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "No Rows",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT rating FROM rating WHERE user_id=$1 AND movie_id = $2`)).
					WithArgs(driver.Value(1), driver.Value(1)).WillReturnError(sql.ErrNoRows)
			},
			data: &proto.MovieRating{
				UserID:  1,
				MovieID: 1,
			},
			expected:    &proto.Rating{Rating: -1},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			avatar, err := storage.GetRating(th.data)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, avatar)
			}
		})
	}
}

func TestProfileRepository_SetToken(t *testing.T) {
	db, _, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		token       string
		userID      int64
		expireTime  int64
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				token := "token"
				userID := int64(1)
				expireTime := int64(time.Hour.Seconds())
				conn.Command("SET", token, userID, "EX", expireTime).ExpectError(nil)
			},
			token:       "token",
			userID:      1,
			expireTime:  int64(time.Hour.Seconds()),
			expectedErr: nil,
		},
		{
			name: "Error occurred in redis Set method",
			mock: func() {
				token := "token"
				userID := int64(1)
				expireTime := int64(time.Hour.Seconds())
				conn.Command("SET", token, userID, "EX", expireTime).ExpectError(errors.New("error"))
			},
			token:       "token",
			userID:      1,
			expireTime:  int64(time.Hour.Seconds()),
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.SetToken(th.token, th.userID, th.expireTime)

			if test.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_GetIdByToken(t *testing.T) {
	db, _, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		input       string
		expected    int64
		expectedErr error
	}{
		{
			name: "Successfully returned id",
			mock: func() {
				conn.Command("GET", "token").Expect(int64(1)).ExpectError(nil)
			},
			input:       "token",
			expected:    int64(1),
			expectedErr: nil,
		},
		{
			name: "Error occurred in redis Get method",
			mock: func() {
				conn.Command("GET", "token").ExpectError(errors.New("error"))
			},
			input:       "cookie",
			expected:    -1,
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			id, err := storage.GetIdByToken(th.input)

			if test.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, id)
			}
		})
	}
}

func TestProfileRepository_CreatePayment(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		token       string
		userID      int64
		price       float64
		expectedErr error
	}{
		{
			name: "Error occurred during INSERT request",
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO payments(amount, users_id, pay_token) VALUES($1, $2, $3)`)).
					WithArgs(
						driver.Value(float64(1)),
						driver.Value(int64(1)),
						driver.Value("token"),
					).WillReturnError(errors.New("Error occurred during request "))
			},
			token:       "token",
			userID:      int64(1),
			price:       float64(1),
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Successfully",
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO payments(amount, users_id, pay_token) VALUES($1, $2, $3)`)).
					WithArgs(
						driver.Value(float64(1)),
						driver.Value(int64(1)),
						driver.Value("token"),
					).WillReturnResult(driver.ResultNoRows).WillReturnError(nil)
			},
			token:       "token",
			userID:      int64(1),
			price:       float64(1),
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.CreatePayment(th.token, th.userID, th.price)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_CreateSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		userID      int64
		expectedErr error
	}{
		{
			name: "Error occurred during UPDATE request",
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET subscription_expires = LOCALTIMESTAMP + interval '1 month' WHERE id=$1`)).
					WithArgs(
						driver.Value(int64(1)),
					).WillReturnError(errors.New("Error occurred during request "))
			},
			userID:      int64(1),
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Successfully",
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE users SET subscription_expires = LOCALTIMESTAMP + interval '1 month' WHERE id=$1`)).
					WithArgs(
						driver.Value(int64(1)),
					).WillReturnResult(driver.ResultNoRows).WillReturnError(nil)
			},
			userID:      int64(1),
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.CreateSubscribe(th.userID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_UpdatePayment(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	tests := []struct {
		name        string
		mock        func()
		token       string
		userID      int64
		expectedErr error
	}{
		{
			name: "Error occurred during UPDATE request",
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE payments SET status = true WHERE users_id=$1 AND pay_token=$2`)).
					WithArgs(
						driver.Value(int64(1)),
						driver.Value("token"),
					).WillReturnError(errors.New("Error occurred during request "))
			},
			userID:      int64(1),
			token:       "token",
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Successfully",
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE payments SET status = true WHERE users_id=$1 AND pay_token=$2`)).
					WithArgs(
						driver.Value(int64(1)),
						driver.Value("token"),
					).WillReturnResult(driver.ResultNoRows).WillReturnError(nil)
			},
			userID:      int64(1),
			token:       "token",
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.UpdatePayment(th.token, th.userID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_CheckCountPaymentsByToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	likes := make([]int64, 0)
	likes = append(likes, 1, 2, 3)

	tests := []struct {
		name        string
		mock        func()
		token       string
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(id) from payments where pay_token = $1`)).
					WithArgs(driver.Value("token")).WillReturnRows(rows)
			},
			token:       "token",
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(id) from payments where pay_token = $1`)).
					WithArgs(driver.Value("token")).WillReturnError(errors.New("Error occurred during request "))
			},
			token:       "token",
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "Wrong Count Payments For Token",
			mock: func() {
				rows := sqlmock.NewRows([]string{"count"}).AddRow(2)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(id) from payments where pay_token = $1`)).
					WithArgs(driver.Value("token")).WillReturnRows(rows)
			},
			token:       "token",
			expectedErr: constants.WrongCountPaymentsForToken,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.CheckCountPaymentsByToken(th.token)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProfileRepository_GetAmountByToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	likes := make([]int64, 0)
	likes = append(likes, 1, 2, 3)

	tests := []struct {
		name           string
		mock           func()
		token          string
		expectedID     int64
		expectedAmount float32
		expectedErr    error
	}{
		{
			name: "Successfully",
			mock: func() {
				rows := sqlmock.NewRows([]string{"users_id", "amount"}).AddRow(1, 1)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT users_id, amount from payments where pay_token = $1`)).
					WithArgs(driver.Value("token")).WillReturnRows(rows)
			},
			token:          "token",
			expectedID:     int64(1),
			expectedAmount: float32(1),
			expectedErr:    nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT users_id, amount from payments where pay_token = $1`)).
					WithArgs(driver.Value("token")).WillReturnError(errors.New("Error occurred during request "))
			},
			token:       "token",
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			id, amount, err := storage.GetAmountByToken(th.token)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, id, th.expectedID)
				assert.Equal(t, amount, th.expectedAmount)
			}
		})
	}
}

func TestProfileRepository_IsSubscription(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, minioClient, pool)

	likes := make([]int64, 0)
	likes = append(likes, 1, 2, 3)

	tests := []struct {
		name        string
		mock        func()
		userID      int64
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE subscription_expires > LOCALTIMESTAMP AND id = $1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			userID:      int64(1),
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE subscription_expires > LOCALTIMESTAMP AND id = $1`)).
					WithArgs(driver.Value(1)).WillReturnError(errors.New("Error occurred during request "))
			},
			userID:      int64(1),
			expectedErr: errors.New("Error occurred during request "),
		},
		{
			name: "No subscription",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE subscription_expires > LOCALTIMESTAMP AND id = $1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			userID:      int64(1),
			expectedErr: constants.NoSubscription,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.IsSubscription(th.userID)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
