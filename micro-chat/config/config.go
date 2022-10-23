package config

import "github.com/spf13/viper"

type Config struct {
	Db DbConfig `mapstructure:"db"`
}

type DbConfig struct {
	Driver     string `mapstructure:"driver"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"pass"`
	Name       string `mapstructure:"name"`
	Collection string `mapstructure:"coll"`
}

func LoadConfig(path, name string) (Config, error) {
	cfg := Config{}
	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	return cfg, err
}
