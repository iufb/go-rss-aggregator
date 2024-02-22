package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iufb/rssagg/internal/database"
)

func (apiCnf *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parametrs struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parametrs{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}
	feed, err := apiCnf.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error while creating feed %v", err))
		return

	}
	respondWithJson(w, 201, databaseFeedToFeed(feed))
}

func (apiCnf *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCnf.DB.GetFeeds(r.Context(), user.ID)
	if err != nil {
		respondWithErr(w, 404, fmt.Sprintf("Feeds not found,", err))
	}
	respondWithJson(w, 200, databaseFeedsToFeeds(feeds))
}
