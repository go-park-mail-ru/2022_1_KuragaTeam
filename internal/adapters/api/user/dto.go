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

type ProfileUserDTO struct {
	Name   string `json:"username" form:"username"`
	Email  string `json:"email" form:"email"`
	Avatar string `json:"avatar" form:"avatar"`
}

type EditProfileDTO struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseUserProfile struct {
	Status   int             `json:"status"`
	UserData *ProfileUserDTO `json:"user"`
}
