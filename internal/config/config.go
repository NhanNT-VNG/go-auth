package config

import "github.com/go-playground/validator"

type Config struct {
	Validate      *validator.Validate
	MongoUsername string
	MongoPassword string
}

func NewConfigurations(validate *validator.Validate) *Config {
	configs := &Config{
		Validate:      validate,
		MongoUsername: "",
		MongoPassword: "",
	}
	return configs
}
