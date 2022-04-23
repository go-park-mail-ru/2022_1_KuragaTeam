package validation

import (
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/utils/constants"
	"strings"
	"unicode"

	"gopkg.in/validator.v2"
)

func ValidateUser(user *proto.SignUpData) error {
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
		letter, num  bool
		symbolsCount uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsLetter(char):
			letter = true
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

	if !letter {
		return constants.ErrLetter
	}
	if !num {
		return constants.ErrNum
	}
	if symbolsCount < 8 {
		return constants.ErrCount
	}

	return nil
}
