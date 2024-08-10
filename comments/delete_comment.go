package comments

import (
	"blog-backend/helpers"
	"database/sql"
	"fmt"
)

func DeleteComment(db *sql.DB, id int, author_id int) error {
	authOK, err := helpers.AuthAction(db, "comments", id, author_id)
	if err != nil {
		return err
	}
	if !authOK {
		return fmt.Errorf("Unauthorizedasd")
	}

	deleteQuery := `DELETE FROM comments WHERE id=$1`

	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
