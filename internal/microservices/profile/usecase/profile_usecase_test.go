package usecase

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/constants"
	"myapp/internal/microservices/profile/proto"
	"myapp/internal/microservices/profile/repository"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProfileUseCase_GetUserProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.UserID
		expected    *proto.ProfileData
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &proto.ProfileData{
					Name:   "Danya",
					Email:  "danya@mail.ru",
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().GetUserProfile(int64(1)).Return(userModel, nil),
				)
			},
			input: &proto.UserID{ID: int64(1)},
			expected: &proto.ProfileData{
				Name:   "Danya",
				Email:  "danya@mail.ru",
				Avatar: "avatar.webp",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetUserProfile",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetUserProfile(int64(1)).Return(nil, errors.New("error")),
				)
			},
			input:       &proto.UserID{ID: int64(1)},
			expected:    nil,
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			profile, err := service.GetUserProfile(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expected.Name, profile.Name)
				assert.Equal(t, th.expected.Email, profile.Email)
				assert.Equal(t, th.expected.Avatar, profile.Avatar)
			}
		})
	}
}

func TestProfileUseCase_EditProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.EditProfileData
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &proto.EditProfileData{
					ID:       int64(1),
					Name:     "Danya",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditProfile(userModel).Return(nil),
				)
			},
			input: &proto.EditProfileData{
				ID:       int64(1),
				Name:     "Danya",
				Password: "danya123321",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in EditProfile",
			mock: func() {
				userModel := &proto.EditProfileData{
					ID:       int64(1),
					Name:     "Danya",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditProfile(userModel).Return(errors.New("error")),
				)
			},
			input: &proto.EditProfileData{
				ID:       int64(1),
				Name:     "Danya",
				Password: "danya123321",
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			_, err := service.EditProfile(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestProfileUseCase_EditAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.EditAvatarData
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &proto.EditAvatarData{
					ID:     int64(1),
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditAvatar(userModel).Return("old_avatar", nil),
					mockStorage.EXPECT().DeleteFile("old_avatar").Return(nil),
				)
			},
			input: &proto.EditAvatarData{
				ID:     int64(1),
				Avatar: "avatar.webp",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in EditAvatar",
			mock: func() {
				userModel := &proto.EditAvatarData{
					ID:     int64(1),
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditAvatar(userModel).Return("", errors.New("error")),
				)
			},
			input: &proto.EditAvatarData{
				ID:     int64(1),
				Avatar: "avatar.webp",
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Error occurred in DeleteFile",
			mock: func() {
				userModel := &proto.EditAvatarData{
					ID:     int64(1),
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditAvatar(userModel).Return("old_avatar", nil),
					mockStorage.EXPECT().DeleteFile("old_avatar").Return(errors.New("error")),
				)
			},
			input: &proto.EditAvatarData{
				ID:     int64(1),
				Avatar: "avatar.webp",
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			_, err := service.EditAvatar(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestProfileUseCase_UploadAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name         string
		mock         func()
		input        *proto.UploadInputFile
		expectedName string
		expectedErr  error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &proto.UploadInputFile{
					ID:          int64(1),
					File:        make([]byte, 0),
					Size:        int64(5),
					ContentType: "type",
				}
				gomock.InOrder(
					mockStorage.EXPECT().UploadAvatar(userModel).Return(
						fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(int64(1))),
							fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
								time.Now().Year(), time.Now().Month(), time.Now().Day(),
								time.Now().Hour(), time.Now().Minute(), time.Now().Second()), "webp"), nil),
				)
			},
			input: &proto.UploadInputFile{
				ID:          int64(1),
				File:        make([]byte, 0),
				Size:        int64(5),
				ContentType: "type",
			},
			expectedName: fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(int64(1))),
				fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					time.Now().Year(), time.Now().Month(), time.Now().Day(),
					time.Now().Hour(), time.Now().Minute(), time.Now().Second()), "webp"),
			expectedErr: nil,
		},
		{
			name: "Error occurred in UploadFile",
			mock: func() {
				userModel := &proto.UploadInputFile{
					ID:          int64(1),
					File:        make([]byte, 0),
					Size:        int64(5),
					ContentType: "type",
				}
				gomock.InOrder(
					mockStorage.EXPECT().UploadAvatar(userModel).Return("", errors.New("error")),
				)
			},
			input: &proto.UploadInputFile{
				ID:          int64(1),
				File:        make([]byte, 0),
				Size:        int64(5),
				ContentType: "type",
			},
			expectedName: "",
			expectedErr:  errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			path, err := service.UploadAvatar(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedName, path.Name)
			}
		})
	}
}

func TestProfileUseCase_GetAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.UserID
		expected    *proto.FileName
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetAvatar(int64(1)).Return(constants.DefaultImage, nil),
				)
			},
			input:       &proto.UserID{ID: int64(1)},
			expected:    &proto.FileName{Name: constants.DefaultImage},
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetAvatar",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetAvatar(int64(1)).Return("", errors.New("error")),
				)
			},
			input:       &proto.UserID{ID: int64(1)},
			expected:    &proto.FileName{},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			avatar, err := service.GetAvatar(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expected, avatar)
			}
		})
	}
}

func TestProfileUseCase_AddLike(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.LikeData
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				data := &proto.LikeData{
					UserID:  1,
					MovieID: 2,
				}
				gomock.InOrder(
					mockStorage.EXPECT().AddLike(data).Return(nil),
				)
			},
			input: &proto.LikeData{
				UserID:  1,
				MovieID: 2,
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in AddLike",
			mock: func() {
				data := &proto.LikeData{
					UserID:  1,
					MovieID: 2,
				}

				gomock.InOrder(
					mockStorage.EXPECT().AddLike(data).Return(errors.New("error")),
				)
			},
			input: &proto.LikeData{
				UserID:  1,
				MovieID: 2,
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			_, err := service.AddLike(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestProfileUseCase_RemoveLike(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.LikeData
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				data := &proto.LikeData{
					UserID:  1,
					MovieID: 2,
				}
				gomock.InOrder(
					mockStorage.EXPECT().RemoveLike(data).Return(nil),
				)
			},
			input: &proto.LikeData{
				UserID:  1,
				MovieID: 2,
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in RemoveLike",
			mock: func() {
				data := &proto.LikeData{
					UserID:  1,
					MovieID: 2,
				}

				gomock.InOrder(
					mockStorage.EXPECT().RemoveLike(data).Return(errors.New("error")),
				)
			},
			input: &proto.LikeData{
				UserID:  1,
				MovieID: 2,
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			_, err := service.RemoveLike(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestProfileUseCase_GetFavorites(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	likes := make([]int64, 0)
	likes = append(likes, 1, 2, 3)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.UserID
		expected    *proto.Favorites
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				data := &proto.UserID{
					ID: 1,
				}
				gomock.InOrder(
					mockStorage.EXPECT().GetFavorites(data.ID).Return(&proto.Favorites{Id: likes}, nil),
				)
			},
			input: &proto.UserID{
				ID: 1,
			},
			expected:    &proto.Favorites{Id: likes},
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetFavorites",
			mock: func() {
				data := &proto.UserID{
					ID: 1,
				}

				gomock.InOrder(
					mockStorage.EXPECT().GetFavorites(data.ID).Return(&proto.Favorites{}, errors.New("error")),
				)
			},
			input: &proto.UserID{
				ID: 1,
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			favorites, err := service.GetFavorites(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, favorites.Id, th.expected.Id)
			}
		})
	}
}
