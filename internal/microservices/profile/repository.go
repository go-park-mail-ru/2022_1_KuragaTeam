package profile

import (
	"myapp/internal/microservices/profile/proto"
)

type Storage interface {
	GetUserProfile(userID int64) (*proto.ProfileData, error)
	EditProfile(data *proto.EditProfileData) error
	EditAvatar(data *proto.EditAvatarData) (string, error)
	GetAvatar(userID int64) (string, error)
	UploadAvatar(data *proto.UploadInputFile) (string, error)
	DeleteFile(string) error
	AddLike(data *proto.LikeData) error
	RemoveLike(data *proto.LikeData) error
	GetFavorites(userID int64) (*proto.Favorites, error)
	GetRating(data *proto.MovieRating) (*proto.Rating, error)
	SetToken(token string, userID int64, expireTime int64) error
	GetIdByToken(token string) (int64, error)
	CreatePayment(token string, userID int64, price float64) error
	CreateSubscribe(userID int64) error
	UpdatePayment(token string, userID int64) error
	CheckCountPaymentsByToken(token string) error
	GetAmountByToken(token string) (int64, float32, error)
	IsSubscription(userID int64) error
}
