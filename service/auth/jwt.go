package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dickyanth/eco-bite-v1/config"
	"github.com/dickyanth/eco-bite-v1/types"
	"github.com/dickyanth/eco-bite-v1/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string
const BuyerKey contextKey = "buyerId"

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

func WithJWTAuth(handleFunc http.HandlerFunc, store types.BuyerStore) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		tokenString := getTokenFromRequest(r)

		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid{
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["buyerId"].(string)

		buyerId, err := strconv.Atoi(str)

		u, err := store.GetBuyerByID(buyerId)
		if err != nil {
			log.Printf("failed to get buyer by id: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, BuyerKey, u.BuyerId)
		r = r.WithContext(ctx)

		handleFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string{
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != ""{
		return tokenAuth
	}

	return ""
}

func validateToken(t string) (*jwt.Token, error){
	return jwt.Parse(t, func (t *jwt.Token) (interface{}, error)  {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter){
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIdFromContext(ctx context.Context) int{
	buyerId, ok := ctx.Value(BuyerKey).(int)
	if !ok{
		return -1
	}

	return buyerId
}