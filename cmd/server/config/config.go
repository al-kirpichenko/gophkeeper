package config

import (
	"fmt"
	"time"
)

const (
	Host       = "localhost:8080"
	DBhost     = "localhost"
	DBuser     = "postgres"
	DBpassword = "123"
	DBname     = "postgres"
)

type Config struct {
	Host        string
	DataBaseURI string
	TokenTTL    time.Duration
}

func NewConfig() *Config {

	return &Config{
		Host: Host,
		DataBaseURI: fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			DBhost, DBuser, DBpassword, DBname),
		TokenTTL: 10 * time.Hour,
	}
}
