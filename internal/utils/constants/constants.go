package constants

import "errors"

var (
	ErrLetter        = errors.New("at least one letter is required")
	ErrNum           = errors.New("at least one digit is required")
	ErrCount         = errors.New("at least eight characters long is required")
	ErrBan           = errors.New("password uses unavailable symbols")
	ErrWrongData     = errors.New("wrong password")
	EmailIsNotUnique = errors.New("email is not unique")
)

const (
	UserObjectsBucketName = "avatars"
	DefaultImage          = "default_avatar.webp"

	UserNotFound           = "User not found"
	UserCanBeLoggedIn      = "User can be logged in"
	UserCreated            = "User created"
	SessionRequired        = "Session required"
	UserIsUnauthorized     = "User is unauthorized"
	UserIsLoggedOut        = "User is logged out"
	FileTypeIsNotSupported = "File type is not supported"
	ProfileIsEdited        = "Profile is edited"
	NoRequestId            = "No RequestID in context"
)

const (
	SignupURL  = "/api/v1/signup"
	LoginURL   = "/api/v1/login"
	LogoutURL  = "/api/v1/logout"
	ProfileURL = "/api/v1/profile"
	EditURL    = "/api/v1/edit"
	AvatarURL  = "/api/v1/avatar"
	CsrfURL    = "/api/v1/csrf"
	AuthURL    = "/api/v1/auth"
)
