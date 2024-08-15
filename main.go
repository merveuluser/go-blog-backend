package main

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/create_post", auth.AuthMiddleware(handlers.CreatePostHandler))
	http.HandleFunc("/update_post", auth.AuthMiddleware(handlers.UpdatePostHandler))
	http.HandleFunc("/delete_post", auth.AuthMiddleware(handlers.DeletePostHandler))
	http.HandleFunc("/add_comment", auth.AuthMiddleware(handlers.AddCommentHandler))
	http.HandleFunc("/update_comment", auth.AuthMiddleware(handlers.UpdateCommentHandler))
	http.HandleFunc("/delete_comment", auth.AuthMiddleware(handlers.DeleteCommentHandler))
	http.HandleFunc("/create_category", auth.AuthMiddleware(handlers.CreateCategoryHandler))
	http.HandleFunc("/update_category", auth.AuthMiddleware(handlers.UpdateCategoryHandler))
	http.HandleFunc("/delete_category", auth.AuthMiddleware(handlers.DeleteCategoryHandler))
	http.HandleFunc("/add_category_to_post", auth.AuthMiddleware(handlers.AddCategoryToPostHandler))
	http.HandleFunc("/remove_category_from_post", auth.AuthMiddleware(handlers.RemoveCategoryFromPostHandler))

	http.HandleFunc("/check_cookie", handlers.CheckCookieHandler)
	http.HandleFunc("/create_tables", handlers.CreateTablesHandler)
	http.HandleFunc("/delete_tables", handlers.DeleteTablesHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
