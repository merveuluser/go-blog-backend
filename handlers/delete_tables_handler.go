package handlers

import (
	"blog-backend/database"
	"encoding/json"
	"log"
	"net/http"
)

func DeleteTablesHandler(w http.ResponseWriter, r *http.Request) {
	err := database.DeleteTables()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Error deleting tables"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Tables deleted"}); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}
}
