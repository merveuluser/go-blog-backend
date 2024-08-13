package posts

import (
	"blog-backend/models"
	"database/sql"
	"time"
)

func UpdatePost(db *sql.DB, id int, title string, content string) (models.Post, error) {
	var post models.Post

	query := `UPDATE posts SET title = $1, content = $2, updated_at = $3 WHERE id = $4;`

	_, err := db.Exec(query, title, content, time.Now(), id)
	if err != nil {
		return post, err
	}

	selectQuery := `SELECT * FROM posts WHERE id = $1;`

	err = db.QueryRow(selectQuery, id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return post, err
	}

	return post, nil
}
