package apiconfig

import (
	"github.com/jinzhu/configor"
)

type APIConfig struct {
	Host         string `env:"API_HOST" default:"localhost"`
	Port         string `env:"API_PORT" default:"8080"`
	Version      string `env:"API_VERSION" default:"v1"`
	BaseURL      string `env:"BASE_URL" default:"challenge"`
	ContactGroup string `env:"CONTACT_GROUP" default:"contact"`
}

func NewAPIConfig() (conf APIConfig, err error) {
	err = configor.New(&configor.Config{Environment: "development"}).Load(&conf)
	return
}
