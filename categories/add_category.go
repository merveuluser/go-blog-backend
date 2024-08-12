package categories

import (
	"blog-backend/models"
	"database/sql"
)

func AddCategory(db *sql.DB, category_name string) (models.Category, error) {
	var category models.Category

	query := `INSERT INTO categories (category_name) VALUES ($1) RETURNING id, category_name`

	err := db.QueryRow(query, category_name).Scan(&category.ID, &category.Name)
	if err != nil {
		return category, err
	}

	return category, nil
}
