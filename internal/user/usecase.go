package user

import "io"

type Service interface {
	SignUp(dto *CreateUserDTO) (string, string, error)
	LogIn(dto *LogInUserDTO) (string, error)
	LogOut(session string) error
	CheckAuthorization(session string) (int64, error)
	GetUserProfile(userID int64) (*ProfileUserDTO, error)
	EditProfile(dto *EditProfileDTO) error
	EditAvatar(dto *EditAvatarDTO) error
	UploadAvatar(file io.Reader, size int64, contentType string, userID int64) (string, error)
}
