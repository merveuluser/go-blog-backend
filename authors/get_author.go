package authors

import (
	"blog-backend/models"
	"database/sql"
)

func GetAuthor(db *sql.DB, id int) (*models.Author, error) {
	var author models.Author

	query := `SELECT * from authors WHERE id = $1;`

	err := db.QueryRow(query, id).Scan(&author.ID, &author.Username, &author.Email, &author.Password, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &author, nil
}
