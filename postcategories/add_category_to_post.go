package postcategories

import (
	"database/sql"
)

func AddCategoryToPost(db *sql.DB, postID int, categoryName string) error {
	var categoryID int

	selectQuery := `SELECT * FROM categories WHERE name = $1;`

	err := db.QueryRow(selectQuery, categoryName).Scan(&categoryID, &categoryName)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO post_categories (post_id, category_id) VALUES ($1, $2);`
	_, err = db.Exec(insertQuery, postID, categoryID)
	if err != nil {
		return err
	}

	return nil
}
