package helpers

import (
	"database/sql"
	"fmt"
)

func AuthAction(db *sql.DB, table string, id int, author_id int) (bool, error) {
	var authOK bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = %d AND author_id = %d)", table, id, author_id)
	err := db.QueryRow(query).Scan(&authOK)
	if err != nil {
		return false, err
	}

	return authOK, nil
}
