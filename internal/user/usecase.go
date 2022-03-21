package user

import "myapp/internal"

type Service interface {
	SignUp(dto *internal.CreateUserDTO) (string, string, error)
	LogIn(dto *internal.LogInUserDTO) (string, error)
	LogOut(session string) error
	CheckAuthorization(session string) (int64, error)
	GetUserMainPage(userID int64) (*internal.MainPageUserDTO, error)
	GetUserProfile(userID int64) (*internal.ProfileUserDTO, error)
	EditProfile(dto *internal.EditProfileDTO) error
}
