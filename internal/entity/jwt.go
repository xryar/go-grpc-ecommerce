package entity

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
}
