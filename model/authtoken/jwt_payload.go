package authtoken

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTPayload struct {
	UserId string `json:"id,omitempty"`
	jwt.StandardClaims
}
