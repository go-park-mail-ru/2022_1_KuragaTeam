package utils

import (
	"github.com/stretchr/testify/assert"
	"myapp/db"
	"myapp/mockDB"
	"myapp/models"
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

func TestAddUsers(t *testing.T) {
	mockDB.MockedDB(mockDB.CREATE)
	defer mockDB.MockedDB(mockDB.DROP)

	//newUser := "fooo"
	//database.AddUser(newUser)

	/*
	   As we know ConnectDB() will connect database named DATABASE_NAME defined in .env,
	   make sure to change that before running tests.
	*/

	database, err := db.ConnectDB()
	if err != nil {
		t.Errorf("Error connecting database in %v\n%v", t.Name(), err)
	}
	defer database.Close()

	newUserInDB := models.User{
		Name:     "user1",
		Email:    "mail@mai.ru",
		Password: "passwordAZAZA123",
	}

	if _, err = CreateUser(database, newUserInDB); err != nil {
		t.Errorf("%v wasn't added.", newUserInDB)
	}

	//if notFound {
	//	t.Errorf("%v was added but was not retrieved.", newUser)
	//}
}
