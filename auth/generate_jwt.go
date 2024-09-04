package auth

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("secret")

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	userID, err := helpers.GetUserID(database.DB, username)
	if err != nil {
		return "", err
	}

	claims := &Claims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
