package domain

type Movie struct {
	ID          int      `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Year        int      `json:"year" form:"year"`
	Country     []string `json:"country"`
	Description string   `json:"description" form:"description"`
	Picture     string   `json:"picture" form:"picture"`
	Video       string   `json:"video" form:"video"`
	Trailer     string   `json:"trailer" form:"trailer"`
	Genre       []string `json:"genre"`
}

type MovieCompilation struct {
	Name   string  `json:"compilationName"`
	Movies []Movie `json:"movies"`
}
