package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"blog-backend/postcategories"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func RemoveCategoryFromPostHandler(w http.ResponseWriter, r *http.Request) {
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

	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Unable to retrieve user ID")
		return
	}

	var postCategory *models.PostCategory
	if encodeErr := json.NewDecoder(r.Body).Decode(&postCategory); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	validate := validator.New()
	if err := validate.Struct(postCategory); err != nil {
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
	if !postCategoryExists {
		helpers.RespondWithError(w, http.StatusConflict, "Category not added to post")
		return
	}

	err = postcategories.RemoveCategoryFromPost(database.DB, postCategory.PostID, postCategory.CategoryName)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error adding category to post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Category removed from post"}); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
