package models

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseUserProfile struct {
	Status   int             `json:"status"`
	UserData *ProfileUserDTO `json:"user"`
}

type ResponseFavorites struct {
	Status        int          `json:"status"`
	FavoritesData *FavoritesID `json:"favorites"`
}

type ResponseMovieRating struct {
	Status int `json:"status"`
	Rating int `json:"rating"`
}
