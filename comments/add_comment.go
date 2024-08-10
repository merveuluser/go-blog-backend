package comments

import (
	"blog-backend/models"
	"database/sql"
	"time"
)

func CreateComment(db *sql.DB, content string, post_id int, author_id int) (models.Comment, error) {
	var comment models.Comment

	query := `INSERT INTO comments (content, post_id, author_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, content, post_id, author_id, created_at`

	err := db.QueryRow(query, content, post_id, author_id, time.Now()).Scan(&comment.ID, &comment.Content, &comment.PostID, &comment.AuthorID, &comment.CreatedAt)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
