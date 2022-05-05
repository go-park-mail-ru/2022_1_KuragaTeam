package repository

import (
	"database/sql/driver"
	"errors"
	"myapp/internal/constants"
	"myapp/internal/microservices/profile/proto"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
)

func TestProfileRepository_GetUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	var minioClient *minio.Client
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db, minioClient)

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
				rows := sqlmock.NewRows([]string{"name", "email", "avatar"}).AddRow("Ilias", "Ilias@mail.ru", constants.DefaultImage)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT username, email, avatar FROM users WHERE id=$1`)).
					WithArgs(driver.Value(1)).WillReturnRows(rows)
			},
			id: int64(1),
			expected: &proto.ProfileData{
				Name:   "Ilias",
				Email:  "Ilias@mail.ru",
				Avatar: "http://" + os.Getenv("HOST") + "/api/v1/minio/avatars/default_avatar.webp",
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

	storage := NewStorage(db, minioClient)

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

	storage := NewStorage(db, minioClient)

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

	storage := NewStorage(db, minioClient)

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

	storage := NewStorage(db, minioClient)

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

	storage := NewStorage(db, minioClient)

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

	storage := NewStorage(db, minioClient)

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
			expected:    &proto.Favorites{MovieId: likes},
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
