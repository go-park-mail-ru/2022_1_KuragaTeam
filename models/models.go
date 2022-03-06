package models

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type Session struct {
	UserID  uint64
	Session string
}
