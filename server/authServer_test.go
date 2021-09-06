package server

import (
	"authService/protocol"
	"authService/repository"
	"authService/service"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func getConn() *sql.DB{

	conn, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable",
	)
	if err != nil {
		log.Errorf("main: unable to open sql connection %v,", err)
	}

	return conn
}


func TestAuthServer_Refresh(t *testing.T) {
	s := NewAuthServer(service.NewAuthService(repository.NewTokenRepository(getConn())))

	ctx,cancel := context.WithTimeout(context.Background(),time.Minute*10)
	defer cancel()

	jwtTok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InRlc3QiLCJleHAiOjE2MzEwMTc2MzV9.1fsX030czrGcCLQK3AJdrPXRdRawPq_ttaLUJj7BNUg"

	resultjwt,resultrt := s.Refresh(ctx,&protocol.RefreshRequest{Rt: &jwtTok})

	log.WithFields(log.Fields{
		"jwt" : resultjwt,
		"rt" : resultrt,
	})

}