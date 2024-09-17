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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "authors")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `authors` does not exist")
		return
	}

	var author *models.Author
	if encodeErr := json.NewDecoder(r.Body).Decode(&author); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}

	validate := validator.New()
	if err := validate.Struct(author); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usernameExists, err := helpers.CheckUserByName(database.DB, author.Username)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if usernameExists {
		helpers.RespondWithError(w, http.StatusConflict, "Username already exists")
		return
	}

	emailExists, err := helpers.CheckEmail(database.DB, author.Email)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if emailExists {
		helpers.RespondWithError(w, http.StatusConflict, "Email already exists")
		return
	}

	author, err = auth.RegisterAuthor(database.DB, author.Username, author.Email, author.Password)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating author")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(author); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
