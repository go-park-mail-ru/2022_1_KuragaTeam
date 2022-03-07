package utils

import "errors"

var (
	upErr    = errors.New("at least one upper case letter is required")
	lowErr   = errors.New("at least one lower case letter is required")
	numErr   = errors.New("at least one digit is required")
	countErr = errors.New("at least eight characters long is required")
	banErr   = errors.New("password uses unavailable symbols")
)
