package handlers

import (
	"blog-backend/categories"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "categories")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `categories` does not exist")
		return
	}

	var category *models.Category
	if encodeErr := json.NewDecoder(r.Body).Decode(&category); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	validJSON := helpers.ValidateJSONCategory("create", category)
	if !validJSON {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	categoryExists, err := helpers.CheckCategoryByName(database.DB, category.Name)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if categoryExists {
		helpers.RespondWithError(w, http.StatusConflict, "Category already exists")
		return
	}

	category, err = categories.CreateCategory(database.DB, category.Name)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating category")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(category); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
