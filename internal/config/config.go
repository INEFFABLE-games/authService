package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	SqlConnString string `env:"SQL_URI,required,notEmpty"`
	GrpcPort      string `env:"GRPC_PORT,required,notEmpty"`
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
