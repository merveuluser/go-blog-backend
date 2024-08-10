package auth

import (
	"blog-backend/models"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func RegisterUser(db *sql.DB, username, email, password string) (models.Author, error) {
	var author models.Author

	if ifExists(db, "username", username) {
		return author, errors.New("Username already exists")
	}

	if ifExists(db, "email", email) {
		return author, errors.New("Email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return author, err
	}

	query := `INSERT INTO authors (username, email, password, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id, username, email, password, created_at, updated_at`

	err = db.QueryRow(query, username, email, hashedPassword, time.Now(), time.Now()).Scan(&author.ID, &author.Username, &author.Email, &author.Password, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return author, err
	}

	return author, nil
}

func ifExists(db *sql.DB, column string, value string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM authors WHERE %s = $1)", column)

	err := db.QueryRow(query, value).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
