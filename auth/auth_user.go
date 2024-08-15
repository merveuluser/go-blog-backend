package auth

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func AuthLogin(db *sql.DB, username, password string) (int, error) {
	var id int
	var hashedPassword string

	query := `SELECT id, password FROM authors WHERE username = $1;`

	err := db.QueryRow(query, username).Scan(&id, &hashedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_id")
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Cookie not found")
			return
		}

		userID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Invalid user ID")
			return
		}

		userExists, err := helpers.CheckUserByID(database.DB, userID)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		if !userExists {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: User not found")
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
