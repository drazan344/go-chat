package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context) error {
	// Implement the logic to create a user in the database
	return nil
}