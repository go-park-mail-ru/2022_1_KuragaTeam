package models

type User struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Salt     string `json:"salt" form:"salt"`
}

type Session struct {
	UserID  int64
	Session string
}
