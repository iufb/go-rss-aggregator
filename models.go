package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/iufb/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	APIKey    string    `json:"apiKey"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.Apikey,
	}
}
