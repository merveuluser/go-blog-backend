package authors

import (
	"blog-backend/models"
	"database/sql"
)

func GetAuthors(db *sql.DB) ([]models.Author, error) {
	query := `SELECT * from authors;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author

	for rows.Next() {
		var author models.Author

		if err := rows.Scan(&author.ID, &author.Username, &author.Email, &author.Password, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil

}
