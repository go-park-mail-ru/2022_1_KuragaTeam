package models

type User struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"username" form:"username" validate:"nonzero"`
	Email    string `json:"email" form:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" form:"password" validate:"min=8"`
	Salt     string `json:"salt" form:"salt"`
}

type Session struct {
	UserID  int64
	Session string
}
