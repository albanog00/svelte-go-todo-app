package api

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"todoapp.com/server/internal/models"
)

func generateJWT(user *models.User) (string, error) {
	jwtExpiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "golang-auth-server",
		"sub":      user.Username,
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(jwtExpiresIn) * time.Minute).Unix(),
		"userId":   user.Id,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return token, err
}

func RetrieveClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := validateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token provided when retrieving claims")
	}

	return claims, nil
}
