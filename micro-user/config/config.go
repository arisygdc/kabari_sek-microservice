package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Database DbConfig     `mapstructure:"db"`
	Token    TokenConfig  `mapstructure:"token"`
	Server   ServerConfig `mapstructure:"server"`
}
type ServerConfig struct {
	SvcName string `mapstructure:"service_name"`
}

type TokenConfig struct {
	Secret string        `mapstructure:"secret"`
	Dur    time.Duration `mapstructure:"dur"`
}

type DbConfig struct {
	Name    string       `mapstructure:"name"`
	Host    string       `mapstructure:"host"`
	Port    int32        `mapstructure:"port"`
	User    string       `mapstructure:"user"`
	Pass    string       `mapstructure:"pass"`
	SslMode string       `mapstructure:"sslmode"`
	Pool    DbPoolConfig `mapstructure:"pool"`
}

type DbPoolConfig struct {
	MinSize int `mapstructure:"minSize"`
	MaxSize int `mapstructure:"maxSize"`
}

func LoadConfig(path, filename string, config *Config) error {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	return err
}
