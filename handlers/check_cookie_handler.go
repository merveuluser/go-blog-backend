package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Cookie not found"}); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("Error encoding JSON response:", encodeErr)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"user_id": cookie.Value}); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding JSON response:", encodeErr)
	}
}
