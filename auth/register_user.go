package auth

import (
	"blog-backend/models"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func RegisterUser(db *sql.DB, username, email, password string) (models.Author, error) {
	var author models.Author
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return author, err
	}

	query := `INSERT INTO authors (username, email, password, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id, username, email, password, created_at, updated_at;`

	err = db.QueryRow(query, username, email, hashedPassword, time.Now(), time.Now()).Scan(&author.ID, &author.Username, &author.Email, &author.Password, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return author, err
	}

	return author, nil
}
