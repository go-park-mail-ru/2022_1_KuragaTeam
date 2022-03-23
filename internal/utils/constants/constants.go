package constants

import "errors"

var (
	ErrLetter    = errors.New("at least one letter is required")
	ErrNum       = errors.New("at least one digit is required")
	ErrCount     = errors.New("at least eight characters long is required")
	ErrBan       = errors.New("password uses unavailable symbols")
	ErrWrongData = errors.New("wrong password or email")
)

const (
	UserObjectsBucketName = "avatars"
	DefaultImage          = "default_avatar.webp"
)
