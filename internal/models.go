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
