package internal

type MainMovieInfoDTO struct {
	ID      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Tagline string `json:"tagline" form:"tagline"`
	Picture string `json:"picture" form:"picture"`
}

type Person struct {
	ID          int      `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Photo       string   `json:"photo" form:"photo"`
	Description string   `json:"description" form:"description"`
	Position    []string `json:"position" form:"position"`
}
type PersonInMovieDTO struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Photo    string `json:"photo" form:"photo"`
	Position string `json:"position" form:"position"`
}

type Movie struct {
	ID              int     `json:"id" form:"id"`
	Name            string  `json:"name" form:"name"`
	NamePicture     string  `json:"name_picture" form:"name_picture"`
	Year            int     `json:"year" form:"year"`
	Duration        string  `json:"duration" form:"duration"`
	AgeLimit        int     `json:"age_limit" form:"age_limit"`
	Description     string  `json:"description" form:"description"`
	KinopoiskRating float32 `json:"kinopoisk_rating" form:"kinopisk_rating"`
	Rating          float32 `json:"rating" form:"rating"`
	Tagline         string  `json:"tagline" form:"tagline"`
	Picture         string  `json:"picture" form:"picture"`
	Video           string  `json:"video" form:"video"`
	Trailer         string  `json:"trailer" form:"trailer"`

	Country []string           `json:"country"`
	Genre   []string           `json:"genre"`
	Staff   []PersonInMovieDTO `json:"staff"`
}

type MovieCompilation struct {
	Name   string  `json:"compilationName"`
	Movies []Movie `json:"movies"`
}

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
