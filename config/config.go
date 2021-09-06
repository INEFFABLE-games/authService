package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DbPort   string `env:"SQLPORT,required,notEmpty"`
	GrpcPort string `env:"GRPCPORT,required,notEmpty"`
	Host     string `env:"SQLHOST,required,notEmpty"`
	User     string `env:"SQLUSER,required,notEmpty"`
	Password string `env:"SQLPASSWORD,required,notEmpty"`
	Dbname   string `env:"SQLDBNAME,required,notEmpty"`
	Sslmode  string `env:"SQLSSLMODE,required,notEmpty"`
	//TokenKey string `env:"TOKENKEY,required,notEmpty"`
}

func NewConfig() *Config {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		log.WithFields(log.Fields{
			"handler": "config",
			"action":  "initialize",
		}).Errorf("unable to pars environment variables %v,", err)
	}

	return &cfg
}
