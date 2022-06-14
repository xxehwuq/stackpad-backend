package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Http struct {
		Port string `mapstructure:"PORT"`
	}

	Db struct {
		Name     string `mapstructure:"DB_NAME"`
		Username string `mapstructure:"DB_USERNAME"`
		Password string `mapstructure:"DB_PASSWORD"`
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
	}

	Hash struct {
		PasswordSalt string `mapstructure:"PASSWORD_SALT"`
	}

	Jwt struct {
		Ttl        time.Duration `mapstructure:"ttl"`
		SigningKey string        `mapstructure:"JWT_SIGNING_KEY"`
	}
}

func Init() (*Config, error) {
	var cfg Config

	if err := readYaml(&cfg); err != nil {
		return nil, err
	}

	if err := readEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readYaml(cfg *Config) error {
	viper.SetConfigFile("config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.Http); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("jwt", &cfg.Jwt); err != nil {
		return err
	}

	return nil
}

func readEnv(cfg *Config) error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg.Http.Port); err != nil {
		return err
	}
	if err := viper.Unmarshal(&cfg.Db); err != nil {
		return err
	}
	if err := viper.Unmarshal(&cfg.Hash); err != nil {
		return err
	}
	if err := viper.Unmarshal(&cfg.Jwt); err != nil {
		return err
	}

	return nil
}
