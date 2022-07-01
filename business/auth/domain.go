package auth

import "github.com/dgrijalva/jwt-go"

type Auth struct {
	Token  string
	UserID string
}

type Claims struct {
	CompanyID   string
	Email       string
	UserID      string
	LevelAccess string
	jwt.StandardClaims
}

type Domain struct {
	Email    string
	Password string
}
