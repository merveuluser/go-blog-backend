package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"blog-backend/postcategories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func AddCategoryToPostHandler(w http.ResponseWriter, r *http.Request) {
	tableNames := []string{"posts", "categories"}

	for _, tableName := range tableNames {
		tableExists, err := helpers.CheckTable(database.DB, tableName)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		switch tableExists {
		case false:
			helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Table `%s` does not exist", tableName))
			return
		}
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

	var postCategory *models.PostCategory
	if encodeErr := json.NewDecoder(r.Body).Decode(&postCategory); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	if postCategory.PostID == 0 || postCategory.CategoryName == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	postExists, err := helpers.CheckPostByID(database.DB, postCategory.PostID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !postExists {
		helpers.RespondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	categoryExists, err := helpers.CheckCategoryByName(database.DB, postCategory.CategoryName)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !categoryExists {
		helpers.RespondWithError(w, http.StatusNotFound, "Category not found")
		return
	}

	authOK, err := helpers.AuthOnPost(database.DB, postCategory.PostID, userID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !authOK {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	
	postCategoryExists, err := helpers.CheckPostCategory(database.DB, postCategory.PostID, postCategory.CategoryName)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if postCategoryExists {
		helpers.RespondWithError(w, http.StatusConflict, "Category already added to post")
		return
	}

	err = postcategories.AddCategoryToPost(database.DB, postCategory.PostID, postCategory.CategoryName)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error adding category to post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Category added to post"}); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
