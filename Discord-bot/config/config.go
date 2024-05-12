package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DiscordToken    string `json:"discord_token" validate:"required"`
	DBName          string `json:"dbname" validate:"required"`
	DBUserName      string `json:"dbusername" validate:"required"`
	DBPassword      string `json:"dbpassword" validate:"required"`
	DBAddress       string `json:"dbaddress" validate:"required"`
	ServerAddress   string `json:"serveraddress" validate:"required"`
	MaxDBRetryCount int    `json:"maxdbretrycount" validate:"required"`
	TextChannelID   string `json:"text_channel_id" vaidate:"required"`
}

func NewConfig() (*Config, error) {
	configFile, err := os.ReadFile("config.json")

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
