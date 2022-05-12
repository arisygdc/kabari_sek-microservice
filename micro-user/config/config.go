package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database DbConfig `mapstructure:"db"`
}

type DbConfig struct {
	Name string       `mapstructure:"name"`
	Host string       `mapstructure:"host"`
	Port string       `mapstructure:"port"`
	User string       `mapstructure:"user"`
	Pass string       `mapstructure:"pass"`
	Pool DbPoolConfig `mapstructure:"pool"`
}

type DbPoolConfig struct {
	MinSize int `mapstructure:"minPoolSize"`
	MaxSize int `mapstructure:"maxPoolSize"`
}

func LoadConfig(path, filename string, config *Config) error {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	return err
}
