package util

type Config struct {
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
}
