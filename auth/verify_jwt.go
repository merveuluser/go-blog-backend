package auth

import (
	"blog-backend/helpers"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Token not found")
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		userID := claims.UserID

		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
