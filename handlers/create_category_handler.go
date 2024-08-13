package handlers

import (
	"blog-backend/categories"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	var category models.Category
	if encodeErr := json.NewDecoder(r.Body).Decode(&category); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	if category.Name == "" {
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
	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(category); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
