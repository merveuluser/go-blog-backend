package helpers

import (
	"database/sql"
)

func AuthOnPost(db *sql.DB, id int, authorId int) (bool, error) {
	var authOK bool

	query := `SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1 AND author_id = $2);`

	err := db.QueryRow(query, id, authorId).Scan(&authOK)
	if err != nil {
		return false, err
	}

	return authOK, nil
}

func AuthOnComment(db *sql.DB, id int, authorId int) (bool, error) {
	var authOK bool

	query := `SELECT EXISTS(SELECT 1 FROM comments WHERE id = $1 AND author_id = $2);`

	err := db.QueryRow(query, id, authorId).Scan(&authOK)
	if err != nil {
		return false, err
	}

	return authOK, nil
}
