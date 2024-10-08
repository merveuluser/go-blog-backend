package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"blog-backend/posts"
	"encoding/json"
	"log"
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
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

	validJSON := helpers.ValidateJSONPost("create", post)
	if !validJSON {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	post.AuthorID = userID

	post.Summary, err = helpers.PostContentSummarizer(post.Content)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error summarizing post content")
		return
	}

	post.URL, err = helpers.PostUrlCreator(post.Title)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating post URL")
		return
	}

	post, err = posts.CreatePost(database.DB, post.Title, post.Content, post.Summary, post.URL, post.AuthorID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(post); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
