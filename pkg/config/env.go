package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppPort   string `mapstructure:"APP_PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`

	MlApi    string `mapstructure:"ML_API"`
	MlApiKey string `mapstructure:"ML_API_KEY"`

	DbDialect  string `mapstructure:"DB_DIALECT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int    `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbName     string `mapstructure:"DB_NAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbSslMode  string `mapstructure:"DB_SSL_MODE"`
}

func LoadEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("File env not found: %v", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Printf("Error on unmarshal env: %v", err)
	}

	return &env
}
