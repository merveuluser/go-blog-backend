package handlers

import (
	"blog-backend/authors"
	"blog-backend/database"
	"blog-backend/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "authors")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `authors` does not exist")
		return
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Cannot convert id to int")
	}

	author, err := authors.GetAuthor(database.DB, id)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get author")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(author); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
