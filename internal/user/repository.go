package user

type Storage interface {
	IsUserExists(user *User) (int64, bool, error)
	IsUserUnique(user *User) (bool, error)
	CreateUser(user *User) (int64, error)
	GetUserProfile(userID int64) (*User, error)
	EditProfile(user *User) error
	EditAvatar(user *User) (string, error)
	GetAvatar(userID int64) (string, error)
}

type RedisStore interface {
	StoreSession(userID int64) (string, error)
	GetUserId(session string) (int64, error)
	DeleteSession(session string) error
}

type ImageStorage interface {
	UploadFile(input UploadInput) (string, error) // Загрузка файлов
	DeleteFile(string) error                      // Удаление файлов
}
