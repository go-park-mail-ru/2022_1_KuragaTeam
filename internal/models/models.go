package models

import "io"

type User struct {
	ID                  int64  `json:"id" form:"id"`
	Name                string `json:"username" form:"username" validate:"nonzero" example:"name"`
	Email               string `json:"email" form:"email" validate:"regexp=^[0-9a-zA-Z!#$%&'*+/=?^_{|}~-]+@[0-9a-zA-Z+/=?^_{|}~-]+(\\.[0-9a-zA-Z+/=?^_{|}~-]+)+$" example:"email@email.com"`
	Password            string `json:"password" form:"password" validate:"min=8" example:"password"`
	Salt                string `json:"salt" form:"salt"`
	Avatar              string `json:"avatar" form:"avatar"`
	SubscriptionExpires string `json:"subscription_expires" form:"subscription_expires"`
}

type ProfileUserDTO struct {
	Name   string `json:"username" form:"username"`
	Email  string `json:"email" form:"email"`
	Avatar string `json:"avatar" form:"avatar"`
}

type EditProfileDTO struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type EditAvatarDTO struct {
	ID     int64  `json:"id" form:"id"`
	Avatar string `json:"avatar" form:"avatar"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseUserProfile struct {
	Status   int             `json:"status"`
	UserData *ProfileUserDTO `json:"user"`
}

type UploadInput struct {
	UserID      int64
	File        io.Reader
	Size        int64
	ContentType string
}
