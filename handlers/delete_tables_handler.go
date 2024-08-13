package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"encoding/json"
	"log"
	"net/http"
)

func DeleteTablesHandler(w http.ResponseWriter, r *http.Request) {
	err := database.DeleteTables()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error deleting tables")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Tables deleted"}); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", encodeErr)
		return
	}
}
