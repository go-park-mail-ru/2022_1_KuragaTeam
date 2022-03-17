package user

type Storage interface {
	IsUserExists(user *User) (int64, bool, error)
	IsUserUnique(user *User) (bool, error)
	CreateUser(user *User) (int64, error)
}

type RedisStore interface {
	StoreSession(userID int64) (string, error)
	GetUserId(session string) (int64, error)
	DeleteSession(session string) error
}
