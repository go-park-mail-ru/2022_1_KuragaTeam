package usecase

import (
	"myapp/constants"
	"myapp/internal"
	"myapp/internal/user"
	"strings"
	"unicode"

	"gopkg.in/validator.v2"
)

type service struct {
	storage    user.Storage
	redisStore user.RedisStore
}

func NewService(storage user.Storage, redisStore user.RedisStore) user.Service {
	return &service{
		storage:    storage,
		redisStore: redisStore,
	}
}

func ValidateUser(user *internal.User) error {
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

func (s *service) SignUp(dto *internal.CreateUserDTO) (string, string, error) {
	userModel := &internal.User{
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

func (s *service) LogIn(dto *internal.LogInUserDTO) (string, error) {
	userModel := &internal.User{
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

func (s *service) GetUserMainPage(userID int64) (*internal.MainPageUserDTO, error) {
	userData, err := s.storage.GetUserMainPage(userID)
	if err != nil {
		return nil, err
	}

	userDTO := internal.MainPageUserDTO{
		Name:   userData.Name,
		Avatar: userData.Avatar,
	}

	return &userDTO, nil
}

func (s *service) GetUserProfile(userID int64) (*internal.ProfileUserDTO, error) {
	userData, err := s.storage.GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	userDTO := internal.ProfileUserDTO{
		Name:  userData.Name,
		Email: userData.Email,
	}

	return &userDTO, nil
}

func (s *service) EditProfile(dto *internal.EditProfileDTO) error {
	userModel := &internal.User{
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
