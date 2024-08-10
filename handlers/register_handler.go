package handlers

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "authors")
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
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "authors table does not exist"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	var author models.Author
	if encodeErr := json.NewDecoder(r.Body).Decode(&author); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}

	if author.Username == "" || author.Email == "" || author.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	author, err = auth.RegisterUser(database.DB, author.Username, author.Email, author.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Error creating user"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(author); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}
}
