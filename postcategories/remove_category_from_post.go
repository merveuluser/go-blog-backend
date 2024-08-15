package postcategories

import (
	"database/sql"
)

func RemoveCategoryFromPost(db *sql.DB, postID int, categoryName string) error {
	var categoryID int

	selectQuery := `SELECT * FROM categories WHERE name = $1;`

	err := db.QueryRow(selectQuery, categoryName).Scan(&categoryID, &categoryName)
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM post_categories WHERE post_id = $1 AND category_id = $2;`

	_, err = db.Exec(deleteQuery, postID, categoryID)
	if err != nil {
		return err
	}

	return nil
}
