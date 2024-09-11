package main

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/handlers"
	corsHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	headersOk := corsHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := corsHandlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := corsHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	r.HandleFunc("/posts", handlers.GetPostsHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.GetPostByIDHandler).Methods("GET")
	r.HandleFunc("/authors", handlers.GetAuthorsHandler).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.GetAuthorByID).Methods("GET")

	r.HandleFunc("/posts", auth.AuthMiddleware(handlers.CreatePostHandler)).Methods("POST")
	r.HandleFunc("/posts/{id}", auth.AuthMiddleware(handlers.UpdatePostHandler)).Methods("PUT")
	r.HandleFunc("/posts/{id}", auth.AuthMiddleware(handlers.DeletePostHandler)).Methods("DELETE")
	r.HandleFunc("/add_comment", auth.AuthMiddleware(handlers.AddCommentHandler)).Methods("POST")
	r.HandleFunc("/update_comment", auth.AuthMiddleware(handlers.UpdateCommentHandler)).Methods("PUT")
	r.HandleFunc("/delete_comment", auth.AuthMiddleware(handlers.DeleteCommentHandler)).Methods("DELETE")
	r.HandleFunc("/create_category", auth.AuthMiddleware(handlers.CreateCategoryHandler)).Methods("POST")
	r.HandleFunc("/update_category", auth.AuthMiddleware(handlers.UpdateCategoryHandler)).Methods("PUT")
	r.HandleFunc("/delete_category", auth.AuthMiddleware(handlers.DeleteCategoryHandler)).Methods("DELETE")
	r.HandleFunc("/add_category_to_post", auth.AuthMiddleware(handlers.AddCategoryToPostHandler)).Methods("POST")
	r.HandleFunc("/remove_category_from_post", auth.AuthMiddleware(handlers.RemoveCategoryFromPostHandler)).Methods("DELETE")

	r.HandleFunc("/create_tables", handlers.CreateTablesHandler).Methods("POST")
	r.HandleFunc("/delete_tables", handlers.DeleteTablesHandler).Methods("DELETE")

	err = http.ListenAndServe(":8080", corsHandlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		return
	}
}
