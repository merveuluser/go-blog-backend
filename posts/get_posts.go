package posts

import (
	"blog-backend/models"
	"database/sql"
)

func GetPosts(db *sql.DB) ([]models.Post, error) {
	query := `
		SELECT p.id, p.title, p.content, p.summary, p.url, p.author_id, a.username, p.created_at, p.updated_at
		FROM posts p
		JOIN authors a ON p.author_id = a.id;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Summary, &post.URL, &post.AuthorID, &post.AuthorUsername, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil

}
