package user

type CreateUserDTO struct {
	Name     string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LogInUserDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type MainPageUserDTO struct {
	Name   string `json:"username" form:"username"`
	Avatar string `json:"avatar" form:"avatar"`
}

type ProfileUserDTO struct {
	Name  string `json:"username" form:"username"`
	Email string `json:"email" form:"email"`
}

type EditProfileDTO struct {
	ID             int64  `json:"id" form:"id"`
	Name           string `json:"username" form:"username"`
	Password       string `json:"password" form:"password"`
	RepeatPassword string `json:"repeat_password" form:"repeat_password"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseUserMainPage struct {
	Status   int              `json:"status"`
	UserData *MainPageUserDTO `json:"user"`
}

type ResponseUserProfile struct {
	Status   int             `json:"status"`
	UserData *ProfileUserDTO `json:"user"`
}
