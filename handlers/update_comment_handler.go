package handlers

import (
	"blog-backend/comments"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "comments")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `comments` does not exist")
		return
	}

	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Unable to retrieve user ID")
		return
	}

	var comment *models.Comment
	if encodeErr := json.NewDecoder(r.Body).Decode(&comment); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	validJSON := helpers.ValidateJSONComment("update", comment)
	if !validJSON {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	commentExists, err := helpers.CheckCommentByID(database.DB, comment.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !commentExists {
		helpers.RespondWithError(w, http.StatusNotFound, "Comment not found")
		return
	}

	comment.AuthorID = userID

	authOK, err := helpers.AuthOnComment(database.DB, comment.ID, comment.AuthorID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !authOK {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	comment, err = comments.UpdateComment(database.DB, comment.ID, comment.Content)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error updating comment")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(comment); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
