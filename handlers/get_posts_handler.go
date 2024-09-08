package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/posts"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "posts")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `posts` does not exist")
		return
	}

	allPosts, err := posts.GetPosts(database.DB)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get posts")
		return
	}

	numOfPosts := len(allPosts)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Total-Count", strconv.Itoa(numOfPosts))
	w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count") // Ensure CORS exposes the header
	if encodeErr := json.NewEncoder(w).Encode(allPosts); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}

}
