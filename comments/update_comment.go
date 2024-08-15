package comments

import (
	"blog-backend/models"
	"database/sql"
	"time"
)

func UpdateComment(db *sql.DB, id int, content string) (*models.Comment, error) {
	var comment models.Comment

	query := `UPDATE comments SET content = $1, updated_at = $2 WHERE id = $3;`

	_, err := db.Exec(query, content, time.Now(), id)
	if err != nil {
		return nil, err
	}

	selectQuery := `SELECT * FROM comments WHERE id = $1;`

	err = db.QueryRow(selectQuery, id).Scan(&comment.ID, &comment.Content, &comment.PostID, &comment.AuthorID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
