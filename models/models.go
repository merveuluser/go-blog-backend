package models

import "time"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
	//Categories []string  `json:"categories"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Author struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PostCategory struct {
	PostID       int    `json:"post_id"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
