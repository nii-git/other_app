package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBName        string `json:"dbname" validate:"required"`
	DBUserName    string `json:"dbusername" validate:"required"`
	DBPassword    string `json:"dbpassword" validate:"required"`
	DBAddress     string `json:"dbaddress" validate:"required"`
	ServerAddress string `json:"serveraddress" validate:"required"`
	LogLevel      string `json:"loglevel" validate:"required"`
}

func NewConfig() (*Config, error) {
	configFilePath := "config/config_" + os.Getenv("ENV") + ".json"
	configFile, err := os.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = json.Unmarshal(configFile, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
