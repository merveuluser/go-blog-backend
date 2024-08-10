package posts

import (
	"blog-backend/helpers"
	"database/sql"
	"fmt"
)

func DeletePost(db *sql.DB, id int, author_id int) error {
	authOK, err := helpers.AuthAction(db, "posts", id, author_id)
	if err != nil {
		return err
	}
	if !authOK {
		return fmt.Errorf("Unauthorized")
	}

	deleteQuery := `DELETE FROM posts WHERE id=$1`

	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
