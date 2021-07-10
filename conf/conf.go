package conf

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	MongoURL string
}

func getConfig() *AppConfig {
	var c AppConfig
	err := envconfig.Process("app", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &c
}
