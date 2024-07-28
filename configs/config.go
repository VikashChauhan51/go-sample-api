package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App      AppSettings
	Server   ServerSettings
	Database DatabaseSettings
}

type AppSettings struct {
	Name    string
	Version string
}

type ServerSettings struct {
	Port int
}

type DatabaseSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var Config AppConfig

func LoadConfig(env string) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")

	if env != "" {
		viper.SetConfigName("config_" + env)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func init() {
	LoadConfig("") // Load default config
}
