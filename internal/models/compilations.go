package models

type FindDTO struct {
	Text string `json:"find" form:"find"`
}

type Genre struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type MovieInfo struct {
	ID      int     `json:"id" form:"id"`
	Name    string  `json:"name" form:"name"`
	Genre   []Genre `json:"genre" form:"genre"`
	Picture string  `json:"picture" form:"picture"`
	Rating  float32 `json:"rating" form:"rating"`
}

type PersonInfo struct {
	ID       int      `json:"id" form:"id"`
	Name     string   `json:"name" form:"name"`
	Photo    string   `json:"photo" form:"photo"`
	Position []string `json:"position" form:"position"`
}

type SearchCompilation struct {
	Movies  []MovieInfo  `json:"movies"`
	Series  []MovieInfo  `json:"series"`
	Persons []PersonInfo `json:"persons"`
}
