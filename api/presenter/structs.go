package presenter

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// ErrorResponse is struct for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}

func (e *ErrorResponse) Wrap(s string, c int) {
	e.Message = s
	e.Code = c
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

// Claims is  a struct that will be encoded to a JWT.
// jwt.StandardClaims is an embedded type to provide expiry time
type Claims struct {
	Email string `json:"name"`
	jwt.StandardClaims
}



// LoginParams is struct to read the request body
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginParams) Validate() error {
	if l.Email == "" {
		return errors.New("email is required")
	}
	if l.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

// SuccessfulLoginResponse is struct to send the request response
type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
}

type SuccessfulSignupResponse struct {
	Email     string
	AuthToken string
	Key 	  string
}

