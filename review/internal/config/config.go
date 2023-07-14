package config

import (
	"log"

	"github.com/ctfrancia/go-code-review/review/cmd/api"

	"github.com/brumhard/alligotor"
)

// Config is the configuration for api.
type Config struct {
	API api.Config
}

// New creates new config instance.
func New() Config {
	cfg := Config{
		API: api.Config{
			Host: "localhost",
			Port: 8080,
		},
	}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
