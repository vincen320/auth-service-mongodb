package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vincen320/auth-service-mongodb/model/appservice/userservice"
	"github.com/vincen320/auth-service-mongodb/model/authtoken"
)

var JWT_SECRET_KEY = []byte("super-secret-key")

func GenerateJWT(user userservice.UserDataGetResponseAppService) string {
	var claims authtoken.JWTPayload

	timeNow := time.Now()
	claims.UserId = user.Id
	claims.ExpiresAt = timeNow.Add(time.Minute * 5).UTC().Unix() //KALAU JWT HARUS UNIX() BUKAN UNIXMILI()
	claims.IssuedAt = timeNow.UTC().Unix()                       // KALAU JWT HARUS UNIX() BUKAN UNIXMILI()
	claims.Issuer = "Vincen App"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(JWT_SECRET_KEY)

	if err != nil {
		panic("error generate token")
	}

	return tokenString
}

func RefreshToken(tokenStr string) string {
	var claims authtoken.JWTPayload

	_, err := jwt.ParseWithClaims(tokenStr, &claims,
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid method")
			}
			return JWT_SECRET_KEY, nil
		})

	if errors, ok := err.(*jwt.ValidationError); ok && errors.Errors == jwt.ValidationErrorExpired {
		claims.ExpiresAt = time.Now().Add(time.Minute * 5).UTC().Unix() //KALAU JWT HARUS UNIX() BUKAN UNIXMILI()
		JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		JWTtokenString, err := JWTtoken.SignedString(JWT_SECRET_KEY)

		if err != nil {
			panic("error generate token")
		}

		return JWTtokenString
	}
	return "token hasn't expired"
}
