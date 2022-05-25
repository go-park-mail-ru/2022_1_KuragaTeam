package models

type ProfileUserDTO struct {
	Name   string `json:"username" form:"username"`
	Email  string `json:"email" form:"email"`
	Avatar string `json:"avatar" form:"avatar"`
	Date   string `json:"date" form:"date"`
}

type EditProfileDTO struct {
	Name     string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LikeDTO struct {
	ID int `json:"id,string" form:"id"`
}

type FavoritesID struct {
	ID []int64 `json:"id" form:"id"`
}

type TokenDTO struct {
	Token string `json:"token" form:"token"`
}
