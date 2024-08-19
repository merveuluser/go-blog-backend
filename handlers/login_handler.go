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
	"strconv"
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

	userID, err := auth.AuthenticateUser(database.DB, login.Username, login.Password) // What if QueryRow returns an error? How to handle that?
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	cookie := &http.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(userID),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(cookie); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
