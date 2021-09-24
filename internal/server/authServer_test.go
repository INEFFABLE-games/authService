package server

import (
	"context"
	"database/sql"
	"github.com/INEFFABLE-games/authService/internal/repository"
	"github.com/INEFFABLE-games/authService/internal/service"
	"github.com/INEFFABLE-games/authService/protocol"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func getConn() *sql.DB {

	conn, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.Errorf("main: unable to open sql connection %v,", err)
	}

	return conn
}

func TestAuthServer_Refresh(t *testing.T) {
	s := NewAuthServer(service.NewAuthService(repository.NewTokenRepository(getConn())))

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	jwtTok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InRlc3QiLCJleHAiOjE2MzEwMTc2MzV9.1fsX030czrGcCLQK3AJdrPXRdRawPq_ttaLUJj7BNUg"

	resultjwt, resultrt := s.Refresh(ctx, &protocol.RefreshRequest{Login: &jwtTok})

	log.WithFields(log.Fields{
		"jwt": resultjwt,
		"rt":  resultrt,
	})

}
