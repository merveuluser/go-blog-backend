package posts

import (
	"blog-backend/models"
	"database/sql"
)

func GetPostByID(db *sql.DB, id int) (*models.Post, error) {
	var post models.Post

	query := `
		SELECT p.id, p.title, p.content, p.summary, p.url, p.author_id, a.username, p.created_at, p.updated_at
		FROM posts p
		JOIN authors a ON p.author_id = a.id
		WHERE p.id = $1;
	`

	err := db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.Summary, &post.URL, &post.AuthorID, &post.AuthorUsername, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
