package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashAndSalt(t *testing.T) {
	tests := []struct {
		name          string
		pwd           string
		salt          string
		areEqualError error
		err           error
	}{
		{
			name:          "Hash",
			pwd:           "123abcABC",
			salt:          "123abc",
			areEqualError: nil,
			err:           nil,
		},
		{
			name:          "Hash",
			pwd:           "qwerty123456QWE",
			salt:          "asdfghmkg,tr463_sf",
			areEqualError: nil,
			err:           nil,
		},
		{
			name:          "Hash",
			pwd:           "mkdlsnbk940389jdsv",
			salt:          "lsdkmvkl-4353sDvv_SDgkl",
			areEqualError: nil,
			err:           nil,
		},
		{
			name:          "Hash",
			pwd:           "98hh4iwnsgGOFDJSRknrkd",
			salt:          "DRGJ:DJB4OE+_Ojrlrsln4GJ",
			areEqualError: nil,
			err:           nil,
		},
		{
			name:          "Hash",
			pwd:           "lksngjdsb$TEMG59e8eg",
			salt:          "sEGSrlkgjdr",
			areEqualError: nil,
			err:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			hashedPwd, err := HashAndSalt(th.pwd, th.salt)
			areEqual := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(th.pwd+th.salt))

			assert.Equal(t, th.err, err)
			assert.Equal(t, th.areEqualError, areEqual)
		})
	}
}

func TestComparePasswords(t *testing.T) {
	testsEqual := []struct {
		name          string
		pwd           string
		salt          string
		hash          []byte
		expectedEqual bool
		err           error
	}{
		{
			name:          "СompareEqual",
			pwd:           "123abcABC",
			salt:          "123abc",
			expectedEqual: true,
		},
		{
			name:          "СompareEqual",
			pwd:           "qwerty123456QWE",
			salt:          "asdfghmkg,tr463_sf",
			expectedEqual: true,
		},
		{
			name:          "СompareEqual",
			pwd:           "mkdlsnbk940389jdsv",
			salt:          "lsdkmvkl-4353sDvv_SDgkl",
			expectedEqual: true,
		},
		{
			name:          "СompareEqual",
			pwd:           "98hh4iwnsgGOFDJSRknrkd",
			salt:          "DRGJ:DJB4OE+_Ojrlrsln4GJ",
			expectedEqual: true,
		},
		{
			name:          "СompareEqual",
			pwd:           "lksngjdsb$TEMG59e8eg",
			salt:          "sEGSrlkgjdr",
			expectedEqual: true,
		},
	}

	for _, test := range testsEqual {
		test.hash, test.err = bcrypt.GenerateFromPassword([]byte(test.pwd+test.salt), bcrypt.MinCost)

		t.Run(test.name, func(t *testing.T) {
			th := test
			areEqual, err := ComparePasswords(string(th.hash), th.salt, th.pwd)

			assert.Equal(t, th.err, err)
			assert.Equal(t, th.expectedEqual, areEqual)
		})
	}

	testsNotEqual := []struct {
		name          string
		pwd           string
		salt          string
		hash          string
		expectedEqual bool
	}{
		{
			name:          "СompareNotEqual",
			pwd:           "123abcABC",
			salt:          "123abc",
			hash:          "wronghashdfgkchkvv6787fghjch",
			expectedEqual: false,
		},
		{
			name:          "СompareNotEqual",
			pwd:           "qwerty123456QWE",
			salt:          "asdfghmkg,tr463_sf",
			hash:          "wronghashhjfchk_78ghvj",
			expectedEqual: false,
		},
	}

	for _, test := range testsNotEqual {
		t.Run(test.name, func(t *testing.T) {
			th := test
			areEqual, _ := ComparePasswords(th.hash, th.salt, th.pwd)

			assert.Equal(t, th.expectedEqual, areEqual)
		})
	}
}
