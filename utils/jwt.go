package utils

import "github.com/dgrijalva/jwt-go"

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func NewJWT(signingKey string) *JWT {
	return &JWT{
		SigningKey: []byte(signingKey),
	}
}

// New creates a new JWT
func (j *JWT) New(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
