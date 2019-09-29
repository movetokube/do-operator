package config

import (
	"errors"
	"os"
	"sync"
)

var doOnce sync.Once

type config struct {
	DOToken string
}

func GetConfig() *config {
	cfg := &config{}
	doOnce.Do(func() {
		if token, found := os.LookupEnv("DO_TOKEN"); found {
			cfg.DOToken = token
		} else {
			panic(errors.New("DO_TOKEN env variable must be set"))
		}
	})
	return cfg
}
