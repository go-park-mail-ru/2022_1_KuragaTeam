package domain

type MainMovieInfoDTO struct {
	ID      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Quote   string `json:"quote" form:"quote"`
	Picture string `json:"picture" form:"picture"`
}

type Person struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Photo string `json:"photo" form:"photo"`
}

type Movie struct {
	ID          int      `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Year        int      `json:"year" form:"year"`
	Quote       string   `json:"quote" form:"quote"`
	Country     []string `json:"country"`
	Description string   `json:"description" form:"description"`
	Picture     string   `json:"picture" form:"picture"`
	Video       string   `json:"video" form:"video"`
	Trailer     string   `json:"trailer" form:"trailer"`
	Genre       []string `json:"genre"`
	Staff       []Person `json:"staff"`
}

type MovieCompilation struct {
	Name   string  `json:"compilationName"`
	Movies []Movie `json:"movies"`
}
