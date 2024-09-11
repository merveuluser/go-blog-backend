package handlers

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"blog-backend/posts"
	"encoding/json"
	"github.com/gorilla/mux"
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

	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Unable to retrieve user ID")
		return
	}

	var post *models.Post
	if encodeErr := json.NewDecoder(r.Body).Decode(&post); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	validJSON := helpers.ValidateJSONPost("update", post)
	if !validJSON {
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

	authOK, err := auth.AuthOnPost(database.DB, post.ID, post.AuthorID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !authOK {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Cannot convert id to int")
		return
	}

	post, err = posts.UpdatePost(database.DB, id, post.Title, post.Content)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error updating post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(post); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
