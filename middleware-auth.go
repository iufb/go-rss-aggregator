package main

import (
	"net/http"

	"github.com/iufb/rssagg/internal/auth"
	"github.com/iufb/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCnf *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
