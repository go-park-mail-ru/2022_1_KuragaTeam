package validation

import (
	"myapp/constants"
	"myapp/internal/user"
	"strings"
	"unicode"

	"gopkg.in/validator.v2"
)

func ValidateUser(user *user.User) error {
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
