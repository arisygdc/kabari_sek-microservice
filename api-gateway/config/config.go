package config

import "github.com/spf13/viper"

const configType string = "yaml"

type Config struct {
	Endpoint ConfigServiceEndpoint `mapstructure:"service_endpoint"`
	Server   ConfigServer          `mapstructure:"server"`
}

type ConfigServiceEndpoint struct {
	User string `mapstructure:"user"`
}

type ConfigServer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func LoadConfig(path, configName string, config *Config) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.AutomaticEnv()

	return viper.Unmarshal(config)
}
