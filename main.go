package main

import (
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
	http.HandleFunc("/create_post", handlers.CreatePostHandler)
	http.HandleFunc("/update_post", handlers.UpdatePostHandler)
	http.HandleFunc("/delete_post", handlers.DeletePostHandler)
	http.HandleFunc("/add_comment", handlers.AddCommentHandler)
	http.HandleFunc("/delete_comment", handlers.DeleteCommentHandler)

	http.HandleFunc("/check_cookie", handlers.CheckCookieHandler)
	http.HandleFunc("/create_tables", handlers.CreateTablesHandler)
	http.HandleFunc("/delete_tables", handlers.DeleteTablesHandler)

	http.ListenAndServe(":8080", nil)
}
