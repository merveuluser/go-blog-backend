package categories

import (
	"database/sql"
)

func DeleteCategory(db *sql.DB, name string) error {
	deleteQuery := `DELETE FROM categories WHERE name = $1`

	_, err := db.Exec(deleteQuery, name)
	if err != nil {
		return err
	}

	return nil
}
