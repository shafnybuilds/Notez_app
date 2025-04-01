package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment         string        `mapstructure:"ENV"`
	LogLevel            string        `mapstructure:"LOG_LEVEL"`
	DBConnString        string        `mapstructure:"DB_CONN_STRING"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	PASETOSecret        string        `mapstructure:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstucture:"ACCESS_TOKEN_DURATION"`
}

// Load and read the configs from file or environment variables
func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("development")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadConfig()
	if err != nil {
		log.Fatalf("Viper couldn't read in the config file. %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Viper couldn't unmarshal the configuration. %v", err)
	}

	return
}
