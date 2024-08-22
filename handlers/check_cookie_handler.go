package handlers

import (
	"blog-backend/helpers"
	"encoding/json"
	"log"
	"net/http"
)

func CheckCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Cookie not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{"user_id": cookie.Value}); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
