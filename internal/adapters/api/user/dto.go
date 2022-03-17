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

type Session struct {
	UserID  int64
	Session string
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
