package posts

import (
	"database/sql"
)

func DeletePost(db *sql.DB, id int) error {
	query := `DELETE FROM posts WHERE id = $1;`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
