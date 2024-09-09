package auth

import (
	"strconv"
	"time"

	"github.com/dickyanth/eco-bite-v1/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"userId":strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "",err
	}
	
	return tokenString, nil
}