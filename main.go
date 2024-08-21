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
	mux.HandleFunc("/create_post", auth.AuthenticateUserMiddleware(handlers.CreatePostHandler))
	mux.HandleFunc("/update_post", auth.AuthenticateUserMiddleware(handlers.UpdatePostHandler))
	mux.HandleFunc("/delete_post", auth.AuthenticateUserMiddleware(handlers.DeletePostHandler))
	mux.HandleFunc("/add_comment", auth.AuthenticateUserMiddleware(handlers.AddCommentHandler))
	mux.HandleFunc("/update_comment", auth.AuthenticateUserMiddleware(handlers.UpdateCommentHandler))
	mux.HandleFunc("/delete_comment", auth.AuthenticateUserMiddleware(handlers.DeleteCommentHandler))
	mux.HandleFunc("/create_category", auth.AuthenticateUserMiddleware(handlers.CreateCategoryHandler))
	mux.HandleFunc("/update_category", auth.AuthenticateUserMiddleware(handlers.UpdateCategoryHandler))
	mux.HandleFunc("/delete_category", auth.AuthenticateUserMiddleware(handlers.DeleteCategoryHandler))
	mux.HandleFunc("/add_category_to_post", auth.AuthenticateUserMiddleware(handlers.AddCategoryToPostHandler))
	mux.HandleFunc("/remove_category_from_post", auth.AuthenticateUserMiddleware(handlers.RemoveCategoryFromPostHandler))
	mux.HandleFunc("/get_posts", handlers.GetPostsHandler)

	mux.HandleFunc("/check_cookie", handlers.CheckCookieHandler)
	mux.HandleFunc("/create_tables", handlers.CreateTablesHandler)
	mux.HandleFunc("/delete_tables", handlers.DeleteTablesHandler)

	handler := cors.Default().Handler(mux)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
