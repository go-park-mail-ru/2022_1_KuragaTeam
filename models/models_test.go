package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
		errs []error
	}{
		{
			"NoCharacterAtAll",
			"",
			[]error{uppErr, lowErr, numErr, countErr},
		},
		{
			"JustEmptyStringAndWhitespace",
			" \n\t\r\v\f ",
			[]error{banErr},
		},
		{
			"MixtureOfEmptyStringAndWhitespace",
			"U u\n1\t?\r1\v2\f34",
			[]error{banErr},
		},
		{
			"MissingUpperCaseString",
			"uu1?1234",
			[]error{uppErr},
		},
		{
			"MissingLowerCaseString",
			"UU1?1234",
			[]error{lowErr},
		},
		{
			"MissingNumber",
			"Uua?aaaa",
			[]error{numErr},
		},
		//{
		//	"MissingSymbol",
		//	"Uu101234",
		//	[]error{errors.New("at least eight characters long is required")},
		//},
		{
			"LessThanRequiredMinimumLength",
			"Uu1?123",
			[]error{countErr},
		},
		{
			"ValidPassword",
			"Uu1?1234",
			nil,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			errs := ValidatePassword(c.pass)

			assert.Equal(t, c.errs, errs)

			for i, err := range errs {
				assert.Equal(t, err, c.errs[i])
			}
		})
	}
}
