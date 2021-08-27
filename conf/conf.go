package conf

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	MongoURL string
	LogLevel uint8
}

func (c *AppConfig) GetConfig() *AppConfig {
	err := envconfig.Process("app", c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
