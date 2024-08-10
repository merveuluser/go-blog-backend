package handlers

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	var login models.Login
	if encodeErr := json.NewDecoder(r.Body).Decode(&login); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}

	if login.Username == "" || login.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	userID, err := auth.AuthLogin(database.DB, login.Username, login.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Invalid username or password"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	cookie := &http.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(userID),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(cookie); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}
}
