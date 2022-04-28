package authorization

import "myapp/internal/microservices/authorization/proto"

type Storage interface {
	IsUserExists(data *proto.LogInData) (int64, error)
	IsUserUnique(email string) (bool, error)
	CreateUser(data *proto.SignUpData) (int64, error)

	StoreSession(userID int64) (string, error)
	GetUserId(session string) (int64, error)
	DeleteSession(session string) error
}
