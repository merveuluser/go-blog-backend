package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"blog-backend/posts"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "posts")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `posts` does not exist")
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Cookie not found")
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid user ID")
		return
	}

	userExists, err := helpers.CheckUserByID(database.DB, userID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !userExists {
		helpers.RespondWithError(w, http.StatusUnauthorized, "User not found")
		return
	}

	var post models.Post
	if encodeErr := json.NewDecoder(r.Body).Decode(&post); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	if strconv.Itoa(post.ID) == "" || post.Title == "" || post.Content == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	postExists, err := helpers.CheckPostByID(database.DB, post.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !postExists {
		helpers.RespondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	post.AuthorID = userID

	authOK, err := helpers.AuthOnPost(database.DB, post.ID, post.AuthorID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !authOK {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	post, err = posts.UpdatePost(database.DB, post.ID, post.Title, post.Content)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error updating post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(post); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
