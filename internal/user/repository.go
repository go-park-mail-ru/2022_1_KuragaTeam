package user

import "myapp/internal"

type Storage interface {
	IsUserExists(user *internal.User) (int64, bool, error)
	IsUserUnique(user *internal.User) (bool, error)
	CreateUser(user *internal.User) (int64, error)
	GetUserMainPage(userID int64) (*internal.User, error)
	GetUserProfile(userID int64) (*internal.User, error)
	EditProfile(user *internal.User) error
}

type RedisStore interface {
	StoreSession(userID int64) (string, error)
	GetUserId(session string) (int64, error)
	DeleteSession(session string) error
}
