package posts

import (
	"blog-backend/models"
	"database/sql"
	"time"
)

func CreatePost(db *sql.DB, title string, content string, summary string, url string, authorId int) (*models.Post, error) {
	var post models.Post

	query := `INSERT INTO posts (title, content, summary, url, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, content, summary, url, author_id, created_at, updated_at;`

	err := db.QueryRow(query, title, content, summary, url, authorId, time.Now(), time.Now()).Scan(&post.ID, &post.Title, &post.Content, &post.Summary, &post.URL, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
