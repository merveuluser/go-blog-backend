package handlers

import (
	"blog-backend/authors"
	"blog-backend/database"
	"blog-backend/helpers"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "authors")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `authors` does not exist")
		return
	}

	allAuthors, err := authors.GetAuthors(database.DB)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	numOfAuthors := len(allAuthors)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Total-Count", strconv.Itoa(numOfAuthors))
	w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count") // Ensure CORS exposes the header
	if encodeErr := json.NewEncoder(w).Encode(allAuthors); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}

}
