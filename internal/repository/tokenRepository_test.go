package repository

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func getConn() *sql.DB {

	conn, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=positions sslmode=disable")
	if err != nil {
		log.Errorf("main: unable to open sql connection %v,", err)
	}

	return conn
}

func TestTokenRepository_InsertOrUpdate(t *testing.T) {

	conn := getConn()

	r := NewTokenRepository(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.InsertOrUpdate(ctx, "45287ty794gh93gb39ubg9u3b 93", "USERUUID")

	require.Nil(t, err)
}
