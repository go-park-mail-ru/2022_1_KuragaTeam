package models

type CreateUserDTO struct {
	Name     string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LogInUserDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
