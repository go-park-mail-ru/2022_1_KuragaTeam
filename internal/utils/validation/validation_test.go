package validation

import (
	"errors"
	"testing"

	"myapp/internal/user"
	"myapp/internal/utils/constants"

	"github.com/stretchr/testify/assert"
	"gopkg.in/validator.v2"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
		err  error
	}{
		{
			name: "JustEmptyStringAndWhitespace",
			pass: " \n\t\r\v\f ",
			err:  constants.ErrBan,
		},
		{
			name: "MixtureOfEmptyStringAndWhitespace",
			pass: "U u\n1\t?\r1\v2\f34",
			err:  constants.ErrBan,
		},
		{
			name: "MissingNumber",
			pass: "Uua?aaaa",
			err:  constants.ErrNum,
		},
		{
			name: "LessThanRequiredMinimumLength",
			pass: "Uu1?123",
			err:  constants.ErrCount,
		},
		{
			name: "MissingLetter",
			pass: "123456789",
			err:  constants.ErrLetter,
		},
		{
			name: "ValidPassword",
			pass: "Uu1?1234",
			err:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			err := ValidatePassword(th.pass)

			assert.Equal(t, th.err, err)
		})
	}
}

func TestUser(t *testing.T) {
	tests := []struct {
		name string
		pass user.User
		err  error
	}{
		{
			name: "NoFieldsAtAll",
			pass: user.User{
				ID:       0,
				Name:     "Ivan",
				Email:    "ivan@mail.ru",
				Password: "123abc123",
				Salt:     "salt",
			},
			err: nil,
		},
		{
			name: "NoFieldsAtAll",
			pass: user.User{
				ID:       0,
				Name:     "Ivan",
				Email:    "ivan@mail.ru",
				Password: "123456678",
				Salt:     "salt",
			},
			err: constants.ErrLetter,
		},
		{
			name: "NoFieldsAtAll",
			pass: user.User{
				ID:       0,
				Name:     "",
				Email:    "",
				Password: "",
				Salt:     "",
			},
			err: validator.ErrorMap{"Email": validator.ErrorArray{validator.TextErr{Err: errors.New("regular expression mismatch")}},
				"Name":     validator.ErrorArray{validator.TextErr{Err: errors.New("zero value")}},
				"Password": validator.ErrorArray{validator.TextErr{Err: errors.New("less than min")}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			err := ValidateUser(&th.pass)

			assert.Equal(t, th.err, err)
		})
	}
}
