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

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "categories")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	if !tableExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "categories table does not exist"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Cookie not found"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Invalid user ID"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	userExists, err := helpers.CheckUser(database.DB, userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	if !userExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "User not found"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	var category models.Category
	if encodeErr := json.NewDecoder(r.Body).Decode(&category); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}

	if category.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	categoryExists, err := helpers.CheckExistsByName(database.DB, "categories", category.Name)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}
	if !categoryExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Category not found"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	err = categories.DeleteCategory(database.DB, category.Name)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"}); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}
}
