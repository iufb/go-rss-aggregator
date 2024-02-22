package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iufb/rssagg/internal/auth"
	"github.com/iufb/rssagg/internal/database"
)

func (apiCnf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parametrs struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parametrs{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}
	user, err := apiCnf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error while creating user %v", err))
		return

	}
	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCnf *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithErr(w, 401, "Unauthorized")
		return
	}
	user, err := apiCnf.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithErr(w, 404, "User not found.")
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}
