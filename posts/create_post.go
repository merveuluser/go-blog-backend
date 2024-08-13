package posts

import (
	"blog-backend/models"
	"database/sql"
	"time"
)

func CreatePost(db *sql.DB, title string, content string, authorId int) (models.Post, error) {
	var post models.Post

	query := `INSERT INTO posts (title, content, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, content, author_id, created_at, updated_at;`

	err := db.QueryRow(query, title, content, authorId, time.Now(), time.Now()).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return post, err
	}

	return post, nil
}
