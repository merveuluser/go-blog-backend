package handlers

import (
	"blog-backend/auth"
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "authors")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `authors` does not exist")
		return
	}

	var login *models.Login
	if encodeErr := json.NewDecoder(r.Body).Decode(&login); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	validate := validator.New()
	if err := validate.Struct(login); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err = auth.AuthenticateUser(database.DB, login.Username, login.Password) // What if QueryRow returns an error? How to handle that?
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	token, err := auth.GenerateJWT(login.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message":      "Login successful",
		"access_token": token,
	}
	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
