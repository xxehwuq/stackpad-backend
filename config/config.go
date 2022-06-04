package config

import "github.com/spf13/viper"

type Config struct {
	Http struct {
		Port string
	}
}

func Init() (*Config, error) {
	var cfg Config

	viper.SetConfigFile("config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
