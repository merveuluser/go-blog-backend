package handlers

import (
	"blog-backend/comments"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "comments")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `comments` does not exist")
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

	var comment *models.Comment
	if encodeErr := json.NewDecoder(r.Body).Decode(&comment); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	if comment.Content == "" || comment.PostID == 0 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	postExists, err := helpers.CheckPostByID(database.DB, comment.PostID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !postExists {
		helpers.RespondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	comment.AuthorID = userID

	comment, err = comments.CreateComment(database.DB, comment.Content, comment.PostID, comment.AuthorID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating comment")
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(comment); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
