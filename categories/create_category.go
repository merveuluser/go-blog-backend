package categories

import (
	"blog-backend/models"
	"database/sql"
)

func CreateCategory(db *sql.DB, categoryName string) (*models.Category, error) {
	var category models.Category

	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id, name;`

	err := db.QueryRow(query, categoryName).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
