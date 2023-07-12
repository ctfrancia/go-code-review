package config

import (
	"log"

	"github.com/ctfrancia/go-code-review/review/cmd/api"

	"github.com/brumhard/alligotor"
)

type Config struct {
	API api.Config
}

func New() Config {
	cfg := Config{
		API: api.Config{},
	}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
