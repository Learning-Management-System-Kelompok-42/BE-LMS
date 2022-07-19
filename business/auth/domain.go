package auth

import "github.com/dgrijalva/jwt-go"

type Auth struct {
	Token            string
	UserID           string
	CompanyID        string
	SpecializationID string
	LevelAccess      string
}

type Claims struct {
	CompanyID        string
	UserID           string
	LevelAccess      string
	SpecializationID string
	jwt.StandardClaims
}

type Domain struct {
	Email    string
	Password string
}
