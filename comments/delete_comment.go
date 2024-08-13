package comments

import (
	"database/sql"
)

func DeleteComment(db *sql.DB, id int) error {
	query := `DELETE FROM comments WHERE id = $1;`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
