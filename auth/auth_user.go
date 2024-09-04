package auth

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(db *sql.DB, username, password string) error {
	var id int
	var hashedPassword string

	query := `SELECT id, password FROM authors WHERE username = $1;`

	err := db.QueryRow(query, username).Scan(&id, &hashedPassword)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
