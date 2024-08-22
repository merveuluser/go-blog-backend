package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"encoding/json"
	"log"
	"net/http"
)

func CreateTablesHandler(w http.ResponseWriter, r *http.Request) {
	err := database.CreateTables()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating tables")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Tables created"}); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
