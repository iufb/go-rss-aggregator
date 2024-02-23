package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/iufb/rssagg/internal/database"
)

func (apiCnf *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parametrs struct {
		FeedID uuid.UUID `json:"feedId"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parametrs{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}
	feed, err := apiCnf.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error while creating feed follows %v", err))
		return

	}
	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feed))
}

func (apiCnf *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCnf.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithErr(w, 404, fmt.Sprintf("Feeds not found, %v", err))
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feeds))
}

func (apiCnf *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	if feedFollowIdStr == "" {
		respondWithErr(w, 400, "Provided wrong id.")
		return
	}
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithErr(w, 400, "Provided wrong id.")
		return
	}
	err = apiCnf.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{ID: feedFollowId, UserID: user.ID})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error while trying to delete: %v", err))
		return
	}
	w.WriteHeader(200)
}
