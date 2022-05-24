package usecase

import (
	"context"
	"errors"
	"myapp/internal/constants"
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/microservices/authorization/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthUseCase_SignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name            string
		mock            func()
		input           *proto.SignUpData
		expectedSession *proto.Cookie
		expectedErr     error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &proto.SignUpData{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel.Email).Return(true, nil),
					mockStorage.EXPECT().CreateUser(userModel).Return(int64(1), nil),
					mockStorage.EXPECT().StoreSession(int64(1)).Return("session", nil),
				)
			},
			input: &proto.SignUpData{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{Cookie: "session"},
			expectedErr:     nil,
		},
		{
			name: "Wrong validation",
			mock: func() {
			},
			input: &proto.SignUpData{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya",
			},
			expectedSession: &proto.Cookie{Cookie: ""},
			expectedErr:     constants.ErrNum,
		},
		{
			name: "User is not unique",
			mock: func() {
				userModel := &proto.SignUpData{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel.Email).Return(false, nil),
				)
			},
			input: &proto.SignUpData{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{Cookie: ""},
			expectedErr:     constants.ErrEmailIsNotUnique,
		},
		{
			name: "Error occurred in IsUserUnique",
			mock: func() {
				userModel := &proto.SignUpData{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel.Email).Return(false, errors.New("error")),
				)
			},
			input: &proto.SignUpData{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{Cookie: ""},
			expectedErr:     errors.New("error"),
		},
		{
			name: "Error occurred in CreateUser",
			mock: func() {
				userModel := &proto.SignUpData{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel.Email).Return(true, nil),
					mockStorage.EXPECT().CreateUser(userModel).Return(int64(1), errors.New("error")),
				)
			},
			input: &proto.SignUpData{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{Cookie: ""},
			expectedErr:     errors.New("error"),
		},
		{
			name: "Error occurred in StoreSession",
			mock: func() {
				userModel := &proto.SignUpData{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel.Email).Return(true, nil),
					mockStorage.EXPECT().CreateUser(userModel).Return(int64(1), nil),
					mockStorage.EXPECT().StoreSession(int64(1)).Return("", errors.New("error")),
				)
			},
			input: &proto.SignUpData{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{Cookie: ""},
			expectedErr:     errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			session, err := service.SignUp(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedSession.Cookie, session.Cookie)
			}
		})
	}
}

func TestAuthUseCase_LogIn(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name            string
		mock            func()
		input           *proto.LogInData
		expectedSession *proto.Cookie
		expectedErr     error
	}{
		{
			name: "Successfully, user exists",
			mock: func() {
				userModel := &proto.LogInData{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(1), nil),
					mockStorage.EXPECT().StoreSession(int64(1)).Return("session", nil),
				)
			},
			input: &proto.LogInData{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{Cookie: "session"},
			expectedErr:     nil,
		},
		{
			name: "Successfully, user not exists",
			mock: func() {
				userModel := &proto.LogInData{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(0), constants.ErrWrongData),
				)
			},
			input: &proto.LogInData{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{},
			expectedErr:     constants.ErrWrongData,
		},
		{
			name: "Error occurred in IsUserExists",
			mock: func() {
				userModel := &proto.LogInData{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(0), errors.New("eroor")),
				)
			},
			input: &proto.LogInData{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{},
			expectedErr:     errors.New("error"),
		},
		{
			name: "Error occurred in StoreSession",
			mock: func() {
				userModel := &proto.LogInData{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(1), nil),
					mockStorage.EXPECT().StoreSession(int64(1)).Return("", errors.New("error")),
				)
			},
			input: &proto.LogInData{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: &proto.Cookie{},
			expectedErr:     errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			session, err := service.LogIn(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedSession.Cookie, session.Cookie)
			}
		})
	}
}

func TestAuthUseCase_LogOut(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.Cookie
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().DeleteSession("session").Return(nil),
				)
			},
			input:       &proto.Cookie{Cookie: "session"},
			expectedErr: nil,
		},
		{
			name: "Error occurred in DeleteSession",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().DeleteSession("session").Return(errors.New("error")),
				)
			},
			input:       &proto.Cookie{Cookie: "session"},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			_, err := service.LogOut(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestAuthUseCase_CheckAuthorization(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := repository.NewMockStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *proto.Cookie
		expectedId  *proto.UserID
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetUserId("session").Return(int64(1), nil),
				)
			},
			input:       &proto.Cookie{Cookie: "session"},
			expectedId:  &proto.UserID{ID: int64(1)},
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetUserID",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetUserId("session").Return(int64(-1), errors.New("error")),
				)
			},
			input:       &proto.Cookie{Cookie: "session"},
			expectedId:  &proto.UserID{ID: int64(-1)},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage)

			id, err := service.CheckAuthorization(context.Background(), th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, th.expectedId.ID, id.ID)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedId.ID, id.ID)
			}
		})
	}
}
