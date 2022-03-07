package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
		err  error
	}{
		{
			name: "NoCharacterAtAll",
			pass: "",
			err:  upErr,
		},
		{
			name: "JustEmptyStringAndWhitespace",
			pass: " \n\t\r\v\f ",
			err:  banErr,
		},
		{
			name: "MixtureOfEmptyStringAndWhitespace",
			pass: "U u\n1\t?\r1\v2\f34",
			err:  banErr,
		},
		{
			name: "MissingUpperCaseString",
			pass: "uu1?1234",
			err:  upErr,
		},
		{
			name: "MissingLowerCaseString",
			pass: "UU1?1234",
			err:  lowErr,
		},
		{
			name: "MissingNumber",
			pass: "Uua?aaaa",
			err:  numErr,
		},
		{
			name: "LessThanRequiredMinimumLength",
			pass: "Uu1?123",
			err:  countErr,
		},
		{
			name: "ValidPassword",
			pass: "Uu1?1234",
			err:  nil,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := ValidatePassword(c.pass)

			assert.Equal(t, c.err, err)
		})
	}
}
