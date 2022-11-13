package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"nats-listener/internal/delivery"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configuration, %s", err.Error())
	}

	delivery.Run()
}

func initConfig() error {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "configs", "Path to config file")
	flag.StringVar(&configFile, "config", "config", "Name of config file")
	flag.StringVar(&configPath, "p", "configs", "Path to config file")
	flag.StringVar(&configFile, "c", "config", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)
	return viper.ReadInConfig()
}
