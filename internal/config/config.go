package config

import (
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do"
	"github.com/spf13/viper"
)

func init() {
	do.Provide(bootstrap.Injector, NewConfig)
}

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Sql struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Name     string `mapstructure:"name"`
		} `mapstructure:"sql"`
	} `mapstructure:"database"`
}

func NewConfig(i *do.Injector) (*Config, error) {
	v := viper.New()

	v.SetConfigType("env")
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
