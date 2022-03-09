package utils

import (
	"errors"
	"gopkg.in/validator.v2"
	"myapp/models"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name string
		pass models.User
		err  error
	}{
		{
			name: "NoFieldsAtAll",
			pass: models.User{
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

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateUser(&c.pass)

			assert.Equal(t, c.err, err)
		})
	}
}

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

//func TestName(t *testing.T) {
//	t.Parallel()
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// given
//	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
//	columns := []string{"id", "price"}
//	pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
//	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
//	orderDao := testdata.OrderDAO{
//		Pool: mockPool,
//	}
//
//	// when
//	actualOrder := orderDao.GetOrderByID(1)
//
//	// then
//	assert.NotNil(t, actualOrder)
//	assert.Equal(t, 100, actualOrder.ID)
//	assert.Equal(t, 100000.9, actualOrder.Price)
//}

// a successful case
func TestUserExists(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "email", "password", "salt"}

	salt, _ := uuid.NewV4()
	pwd := "Pass123321"
	password, _ := HashAndSalt(pwd, salt.String())

	pgxRows := pgxpoolmock.NewRows(columns).AddRow(int64(1), "Ilias@mail.ru", password, salt.String()).ToPgxRows()

	user := models.User{
		ID:       1,
		Name:     "Ilias",
		Email:    "Ilias@mail.ru",
		Password: pwd,
		Salt:     salt.String(),
	}
	mockPool.EXPECT().Query(gomock.Any(), `SELECT id, email, password, salt FROM users WHERE email=$1`, user.Email).Return(pgxRows, nil)

	// when
	//actualOrder := orderDao.GetOrderByID(1)

	userPool := &UserPool{
		Pool: mockPool,
	}

	userID, result, err := userPool.IsUserExists(user)

	// then
	assert.Equal(t, int64(1), userID)
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}
