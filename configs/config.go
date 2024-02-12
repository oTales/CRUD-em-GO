package configs

import (
	"errors"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.user", "user_todo")
	viper.SetDefault("db.password", "1212")
	viper.SetDefault("db.name", "api_todo")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
