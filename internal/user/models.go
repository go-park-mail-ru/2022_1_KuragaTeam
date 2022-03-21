package user

type User struct {
	ID                  int64  `json:"id" form:"id"`
	Name                string `json:"username" form:"username" validate:"nonzero" example:"name"`
	Email               string `json:"email" form:"email" validate:"regexp=^[0-9a-zA-Z!#$%&'*+/=?^_{|}~-]+@[0-9a-zA-Z+/=?^_{|}~-]+(\\.[0-9a-zA-Z+/=?^_{|}~-]+)+$" example:"email@email.com"`
	Password            string `json:"password" form:"password" validate:"min=8" example:"password"`
	Salt                string `json:"salt" form:"salt"`
	Avatar              string `json:"avatar" form:"avatar"`
	SubscriptionExpires string `json:"subscription_expires" form:"subscription_expires"`
}

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
