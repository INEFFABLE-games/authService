package repository

import (
	"context"
	"database/sql"
)

// TokenRepository struct for tokens repository
type TokenRepository struct {
	db *sql.DB
}

// InsertOrUpdate insert new token in tokens table
func (t *TokenRepository) InsertOrUpdate(ctx context.Context, RT string, userUid string) error {

	_, err := t.db.ExecContext(ctx,
		"insert into tokens(uid,value) values($1,$2) ON CONFLICT (uid) DO update set uid=$1, value=$2",
		userUid,
		RT,
	)

	return err
}

// NewTokenRepository creates new TokenRepository object
func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db: db}
}
