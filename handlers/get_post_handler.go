package handlers

import (
	"blog-backend/database"
	"blog-backend/helpers"
	"blog-backend/posts"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	tableExists, err := helpers.CheckTable(database.DB, "posts")
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if !tableExists {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Table `posts` does not exist")
		return
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Cannot convert id to int")
		return
	}

	post, err := posts.GetPost(database.DB, id)
	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(post); encodeErr != nil {
		log.Println("Error encoding JSON response:", encodeErr)
	}
}
