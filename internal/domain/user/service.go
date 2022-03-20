package user

import (
	"myapp/constants"
	"myapp/internal/adapters/api/user"
	"strings"
	"unicode"

	"gopkg.in/validator.v2"
)

type service struct {
	storage    Storage
	redisStore RedisStore
}

func NewService(storage Storage, redisStore RedisStore) user.Service {
	return &service{
		storage:    storage,
		redisStore: redisStore,
	}
}

func ValidateUser(user *User) error {
	user.Name = strings.TrimSpace(user.Name)
	if err := validator.Validate(user); err != nil {
		return err
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}
	return nil
}

func ValidatePassword(pass string) error {
	var (
		upp, low, num bool
		symbolsCount  uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			symbolsCount++
		case unicode.IsLower(char):
			low = true
			symbolsCount++
		case unicode.IsNumber(char):
			num = true
			symbolsCount++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			symbolsCount++
		default:
			return constants.ErrBan
		}
	}

	if !upp {
		return constants.ErrUp
	}
	if !low {
		return constants.ErrLow
	}
	if !num {
		return constants.ErrNum
	}
	if symbolsCount < 8 {
		return constants.ErrCount
	}

	return nil
}

func (s *service) SignUp(dto *user.CreateUserDTO) (string, string, error) {
	userModel := &User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	if err := ValidateUser(userModel); err != nil {
		return "", "", err
	}

	isUnique, err := s.storage.IsUserUnique(userModel)
	if err != nil {
		return "", "", err
	}

	if !isUnique {
		return "", "ERROR: Email is not unique", nil
	}

	userID, err := s.storage.CreateUser(userModel)
	session, err := s.redisStore.StoreSession(userID)

	if err != nil {
		return "", "", err
	}

	return session, "", nil
}

func (s *service) LogIn(dto *user.LogInUserDTO) (string, error) {
	userModel := &User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	userID, userExists, err := s.storage.IsUserExists(userModel)
	if err != nil {
		return "", err
	}

	if !userExists {
		return "", nil
	}

	session, err := s.redisStore.StoreSession(userID)
	if err != nil {
		return "", err
	}

	return session, nil
}

func (s *service) LogOut(session string) error {
	return s.redisStore.DeleteSession(session)
}

func (s *service) CheckAuthorization(session string) (int64, error) {
	userID, err := s.redisStore.GetUserId(session)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (s *service) GetUserProfile(userID int64) (*user.ProfileUserDTO, error) {
	userData, err := s.storage.GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	userDTO := user.ProfileUserDTO{
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar,
	}

	return &userDTO, nil
}

func (s *service) EditProfile(dto *user.EditProfileDTO) error {
	userModel := &User{
		ID:       dto.ID,
		Name:     dto.Name,
		Password: dto.Password,
	}

	err := s.storage.EditProfile(userModel)
	if err != nil {
		return err
	}

	return nil
}
