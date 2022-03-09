package models

type User struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"username" form:"username" validate:"nonzero" example:"name"`
	Email    string `json:"email" form:"email" validate:"regexp=^[0-9a-zA-Z]+@[0-9a-zA-Z]+(\\.[0-9a-zA-Z]+)+$" example:"email@email.com"`
	Password string `json:"password" form:"password" validate:"min=8" example:"password"`
	Salt     string `json:"salt" form:"salt"`
}

type Movie struct {
	Img   string `json:"img" example:"star.png"`
	Href  string `json:"href" example:"/"`
	Name  string `json:"name" example:"StarWars"`
	Genre string `json:"genre" example:"Comedy"`
}

type MovieCompilation struct {
	Name   string  `json:"compilationName"`
	Movies []Movie `json:"movies"`
}

type Session struct {
	UserID  int64
	Session string
}
