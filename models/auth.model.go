package models

import "github.com/golang-jwt/jwt"

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type ResponseAuth struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
