package handlers

import (
	"blog-backend/comments"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	var comment models.Comment
	if encodeErr := json.NewDecoder(r.Body).Decode(&comment); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	if strconv.Itoa(comment.ID) == "" {
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

	err = comments.DeleteComment(database.DB, comment.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error deleting comment")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Comment deleted"}); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
