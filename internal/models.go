package internal

type MainMovieInfoDTO struct {
	ID          int    `json:"id" form:"id"`
	NamePicture string `json:"name_picture" form:"name_picture"`
	Tagline     string `json:"tagline" form:"tagline"`
	Picture     string `json:"picture" form:"picture"`
}

type Genre struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type Person struct {
	ID          int      `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Photo       string   `json:"photo" form:"photo"`
	AdditPhoto1 string   `json:"addit_photo_1" form:"addit_photo_1"`
	AdditPhoto2 string   `json:"addit_photo_2" form:"addit_photo_2"`
	Description string   `json:"description" form:"description"`
	Position    []string `json:"position" form:"position"`
}
type PersonInMovieDTO struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Photo    string `json:"photo" form:"photo"`
	Position string `json:"position" form:"position"`
}

type Episode struct {
	ID          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Number      int    `json:"number" form:"number"`
	Description string `json:"description" form:"description"`
	Video       string `json:"video" form:"video"`
	Picture     string `json:"picture" form:"picture"`
}

type Season struct {
	ID       int       `json:"id" form:"id"`
	Number   int       `json:"number" form:"number"`
	Episodes []Episode `json:"episodes" form:"episodes"`
}

type Movie struct {
	ID              int     `json:"id" form:"id"`
	Name            string  `json:"name" form:"name"`
	IsMovie         bool    `json:"is_movie" form:"is_movie"`
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

	Season  []Season           `json:"season"`
	Country []string           `json:"country"`
	Genre   []Genre            `json:"genre"`
	Staff   []PersonInMovieDTO `json:"staff"`
}

type MovieInfo struct {
	ID      int     `json:"id" form:"id"`
	Name    string  `json:"name" form:"name"`
	Genre   []Genre `json:"genre" form:"genre"`
	Picture string  `json:"picture" form:"picture"`
}

type MovieCompilation struct {
	Name   string      `json:"compilation_name"`
	Movies []MovieInfo `json:"movies"`
}
