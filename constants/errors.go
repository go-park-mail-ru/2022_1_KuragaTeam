package constants

import "errors"

var (
	ErrUp            = errors.New("at least one upper case letter is required")
	ErrLow           = errors.New("at least one lower case letter is required")
	ErrNum           = errors.New("at least one digit is required")
	ErrCount         = errors.New("at least eight characters long is required")
	ErrBan           = errors.New("password uses unavailable symbols")
	ErrWrongPassword = errors.New("wrong password or email")
)
