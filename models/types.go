package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User struct for user object.
type User struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Token struct for token object.
type Token struct {
	Login string `json:"login" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// CustomClaims struct for jwt token generation claim.
type CustomClaims struct {
	Uid string
	jwt.StandardClaims
}
