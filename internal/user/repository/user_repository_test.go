package repository

import (
	"database/sql/driver"
	"errors"
	"myapp/internal/user"
	"myapp/internal/utils/constants"
	"myapp/internal/utils/hash"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_IsUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name           string
		mock           func()
		input          user.User
		expectedID     int64
		expectedResult bool
		expectedErr    error
	}{
		{
			name: "User is in the db",
			mock: func() {
				salt, _ := uuid.NewV4()
				pwd := "Pass123321"
				password, _ := hash.HashAndSalt(pwd, salt.String())

				rows := sqlmock.NewRows([]string{"id", "password", "salt"}).
					AddRow("1", password, salt)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password, salt FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnRows(rows)
			},
			input: user.User{
				Email:    "Ilias@mail.ru",
				Password: "Pass123321",
			},
			expectedID:     1,
			expectedResult: true,
			expectedErr:    nil,
		},
		{
			name: "User is not in db, wrong email",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "password", "salt"})
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password, salt FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnRows(rows)
			},
			input: user.User{
				Email:    "Ilias@mail.ru",
				Password: "Pass123321",
			},
			expectedID:     0,
			expectedResult: false,
			expectedErr:    nil,
		},
		{
			name: "User is in database, wrong password",
			mock: func() {
				salt, _ := uuid.NewV4()
				pwd := "Pass111111"
				password, _ := hash.HashAndSalt(pwd, salt.String())

				rows := sqlmock.NewRows([]string{"id", "password", "salt"}).
					AddRow("1", password, salt)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password, salt FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnRows(rows)
			},
			input: user.User{
				Email:    "Ilias@mail.ru",
				Password: "Pass123321",
			},
			expectedID:     1,
			expectedResult: false,
			expectedErr:    constants.ErrWrongData,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password, salt FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnError(errors.New("Error occurred during request "))
			},
			input: user.User{
				Email:    "Ilias@mail.ru",
				Password: "Pass123321",
			},
			expectedID:     0,
			expectedResult: false,
			expectedErr:    errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			id, result, err := storage.IsUserExists(&th.input)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expectedID, id)
				assert.Equal(t, th.expectedResult, result)
			}
		})
	}
}

func TestUserRepository_IsUserUnique(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		input       user.User
		expected    bool
		expectedErr error
	}{
		{
			name: "Email is unique",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnRows(rows)
			},
			input: user.User{
				Email: "Ilias@mail.ru",
			},
			expected:    true,
			expectedErr: nil,
		},
		{
			name: "Email is not unique",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnRows(rows)
			},
			input: user.User{
				Email: "Ilias@mail.ru",
			},
			expected:    false,
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnError(errors.New("Error occurred during request "))
			},
			input: user.User{
				Email: "Ilias@mail.ru",
			},
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			result, err := storage.IsUserUnique(&th.input)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, result)
			}
		})
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		input       user.User
		expectedID  int64
		expectedErr error
	}{
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users(username, email, password, salt, avatar, subscription_expires) VALUES($1, $2, $3, $4, $5, LOCALTIMESTAMP) RETURNING id`)).
					WithArgs(
						driver.Value("Ilias"),
						driver.Value("Ilias@mail.ru"),
						driver.Value("IliasPassword"),
						driver.Value("salt"),
						driver.Value(constants.DefaultImage),
					).WillReturnError(errors.New("Error occurred during request "))
			},
			input: user.User{
				Name:     "Ilias",
				Email:    "Ilias@mail.ru",
				Password: "IliasPassword",
			},
			expectedID:  1,
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			id, err := storage.CreateUser(&th.input)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expectedID, id)
			}
		})
	}
}

func TestUserRepository_GetUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		id          int64
		expected    *user.User
		expectedErr error
	}{
		{
			name: "Get user profile",
			mock: func() {
				rows := sqlmock.NewRows([]string{"name", "email", "avatar"}).AddRow("Ilias", "Ilias@mail.ru", constants.DefaultImage)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, email, avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id: int64(1),
			expected: &user.User{
				Name:   "Ilias",
				Email:  "Ilias@mail.ru",
				Avatar: "http://localhost:9000/avatars/default_avatar.webp",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, email, avatar FROM users WHERE id=$1`)).
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

			user, err := storage.GetUserProfile(th.id)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, user)
			}
		})
	}
}

func TestUserRepository_EditProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name        string
		mock        func()
		user        *user.User
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
			user: &user.User{
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
			user: &user.User{
				Name:     "Ilias",
				Avatar:   "avatar",
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
			user: &user.User{
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

func TestUserRepository_EditAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	tests := []struct {
		name              string
		mock              func()
		user              *user.User
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
			user: &user.User{
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
			user: &user.User{
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
			user: &user.User{
				Name:   "Ivan",
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
			user: &user.User{
				Name: "Ivan",
			},
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

func TestRedisStore_GetUserId(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

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
				conn.Command("GET", "cookie").Expect(int64(1)).ExpectError(nil)
			},
			input:       "cookie",
			expected:    int64(1),
			expectedErr: nil,
		},
		{
			name: "Error occurred in redis Get method",
			mock: func() {
				conn.Command("GET", "cookie").ExpectError(errors.New("error"))
			},
			input:       "cookie",
			expected:    -1,
			expectedErr: errors.New("error"),
		},
	}

	r := NewRedisStore(pool)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			id, err := r.GetUserId(th.input)

			if test.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, id)
			}
		})
	}
}

func TestRedisStore_DeleteSession(t *testing.T) {
	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	tests := []struct {
		name        string
		mock        func()
		input       string
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				conn.Command("DEL", "cookie").ExpectError(nil)
			},
			input:       "cookie",
			expectedErr: nil,
		},
		{
			name: "Error occurred in redis Del method",
			mock: func() {
				conn.Command("DEL", "cookie").ExpectError(errors.New("error"))
			},
			input:       "cookie",
			expectedErr: errors.New("error"),
		},
	}

	r := NewRedisStore(pool)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := r.DeleteSession(th.input)

			if test.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
