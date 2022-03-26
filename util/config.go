package util

import "github.com/spf13/viper"

type Config struct {
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
