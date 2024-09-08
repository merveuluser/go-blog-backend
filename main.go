package main

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/handlers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/create_post", auth.AuthMiddleware(handlers.CreatePostHandler))
	mux.HandleFunc("/update_post", auth.AuthMiddleware(handlers.UpdatePostHandler))
	mux.HandleFunc("/delete_post", auth.AuthMiddleware(handlers.DeletePostHandler))
	mux.HandleFunc("/add_comment", auth.AuthMiddleware(handlers.AddCommentHandler))
	mux.HandleFunc("/update_comment", auth.AuthMiddleware(handlers.UpdateCommentHandler))
	mux.HandleFunc("/delete_comment", auth.AuthMiddleware(handlers.DeleteCommentHandler))
	mux.HandleFunc("/create_category", auth.AuthMiddleware(handlers.CreateCategoryHandler))
	mux.HandleFunc("/update_category", auth.AuthMiddleware(handlers.UpdateCategoryHandler))
	mux.HandleFunc("/delete_category", auth.AuthMiddleware(handlers.DeleteCategoryHandler))
	mux.HandleFunc("/add_category_to_post", auth.AuthMiddleware(handlers.AddCategoryToPostHandler))
	mux.HandleFunc("/remove_category_from_post", auth.AuthMiddleware(handlers.RemoveCategoryFromPostHandler))
	mux.HandleFunc("/posts", handlers.GetPostsHandler)
	mux.HandleFunc("/authors", handlers.GetAuthorsHandler)

	mux.HandleFunc("/create_tables", handlers.CreateTablesHandler)
	mux.HandleFunc("/delete_tables", handlers.DeleteTablesHandler)

	handler := cors.Default().Handler(mux)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
