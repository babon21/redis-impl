package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port      string
		ServerUrl string
	}
}

func Init() Config {
	var config Config
	viper.AutomaticEnv()
	config.Server.Port = viper.GetString("CLIENT_PORT")
	config.Server.ServerUrl = viper.GetString("SERVER_URL")
	return config
}
