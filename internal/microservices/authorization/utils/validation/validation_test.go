package validation

import (
	"myapp/internal/microservices/authorization/proto"
	"testing"

	"myapp/internal/constants"

	"github.com/stretchr/testify/assert"
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
		pass proto.SignUpData
		err  error
	}{
		{
			name: "NoFieldsAtAll",
			pass: proto.SignUpData{
				Name:     "Ivan",
				Email:    "ivan@mail.ru",
				Password: "123abc123",
			},
			err: nil,
		},
		{
			name: "NoFieldsAtAll",
			pass: proto.SignUpData{
				Name:     "Ivan",
				Email:    "ivan@mail.ru",
				Password: "123456678",
			},
			err: constants.ErrLetter,
		},
		{
			name: "NoFieldsAtAll",
			pass: proto.SignUpData{
				Name:     "",
				Email:    "",
				Password: "",
			},
			err: constants.ErrLetter,
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
