package domain

type Movie struct {
	ID          int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Picture     string `json:"picture" form:"picture"`
	Video       string `json:"video" form:"video"`
	Trailer     string `json:"trailer" form:"trailer"`
}
