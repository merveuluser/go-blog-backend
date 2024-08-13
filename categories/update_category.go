package categories

import (
	"blog-backend/models"
	"database/sql"
)

func UpdateCategory(db *sql.DB, id int, name string) (models.Category, error) {
	var category models.Category

	query := `UPDATE categories SET category_name = $1 WHERE id = $2;`

	_, err := db.Exec(query, name, id)
	if err != nil {
		return category, err
	}

	selectQuery := `SELECT * FROM categories WHERE id = $1;`

	err = db.QueryRow(selectQuery, id).Scan(&category.ID, &category.Name)
	if err != nil {
		return category, err
	}

	return category, nil
}
