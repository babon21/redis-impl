package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
}

func Init() Config {
	var config Config
	viper.AutomaticEnv()
	config.Server.Port = viper.GetString("SERVER_PORT")
	return config
}
