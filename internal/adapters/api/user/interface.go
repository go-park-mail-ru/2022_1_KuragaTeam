package user

type Service interface {
	SignUp(dto *CreateUserDTO) (string, string, error)
	LogIn(dto *LogInUserDTO) (string, error)
	LogOut(session string) error
	CheckAuthorization(session string) (int64, error)
}
