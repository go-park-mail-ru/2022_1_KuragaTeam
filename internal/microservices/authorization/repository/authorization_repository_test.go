package repository

import (
	"database/sql/driver"
	"errors"
	"myapp/internal/constants"
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/microservices/authorization/utils/hash"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository_IsUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, pool)

	tests := []struct {
		name           string
		mock           func()
		input          proto.LogInData
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
			input: proto.LogInData{
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
			input: proto.LogInData{
				Email:    "Ilias@mail.ru",
				Password: "Pass123321",
			},
			expectedID:     0,
			expectedResult: false,
			expectedErr:    constants.ErrWrongData,
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
			input: proto.LogInData{
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
			input: proto.LogInData{
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

			id, err := storage.IsUserExists(&th.input)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expectedID, id)
			}
		})
	}
}

func TestAuthRepository_IsUserUnique(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, pool)

	tests := []struct {
		name        string
		mock        func()
		input       string
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
			input:       "Ilias@mail.ru",
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
			input:       "Ilias@mail.ru",
			expected:    false,
			expectedErr: nil,
		},
		{
			name: "Error occurred during SELECT request",
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM users WHERE email=$1`)).
					WithArgs(driver.Value("Ilias@mail.ru")).WillReturnError(errors.New("Error occurred during request "))
			},
			input:       "Ilias@mail.ru",
			expectedErr: errors.New("Error occurred during request "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			result, err := storage.IsUserUnique(th.input)
			if th.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, result)
			}
		})
	}
}

func TestAuthRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, pool)

	tests := []struct {
		name        string
		mock        func()
		input       proto.SignUpData
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
			input: proto.SignUpData{
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

func TestAuthRepository_GetUserId(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, pool)

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			id, err := storage.GetUserID(th.input)

			if test.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expected, id)
			}
		})
	}
}

func TestAuthRepository_DeleteSession(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}

	storage := NewStorage(db, pool)

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			err := storage.DeleteSession(th.input)

			if test.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
