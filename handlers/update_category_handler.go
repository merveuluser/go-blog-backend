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

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "categories")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `categories` does not exist")
		return
	}

	_, ok := r.Context().Value("user_id").(int)
	if !ok {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized: Unable to retrieve user ID")
		return
	}

	var category *models.Category
	if encodeErr := json.NewDecoder(r.Body).Decode(&category); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	if category.ID == 0 || category.Name == "" {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	categoryExists, err := helpers.CheckCategoryByID(database.DB, category.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if !categoryExists {
		helpers.RespondWithError(w, http.StatusNotFound, "Category not found")
		return
	}

	categoryNameUnchanged, err := helpers.CheckCategoryByName(database.DB, category.Name)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if categoryNameUnchanged {
		helpers.RespondWithError(w, http.StatusNotFound, "Category name unchanged")
		return
	}

	category, err = categories.UpdateCategory(database.DB, category.ID, category.Name)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error updating category")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(category); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
