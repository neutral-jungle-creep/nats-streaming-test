package configs

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	Nats     *NatsConfig
	DataBase *DataBaseConfig
}

type NatsConfig struct {
	URL       string
	ClientID  string
	ClusterID string
	Subj      string
}

type DataBaseConfig struct {
	ConnLink string
}

func NewConfig() *Config {
	return &Config{
		Nats: &NatsConfig{
			URL:       viper.GetString("nats.URL"),
			ClientID:  viper.GetString("nats.client"),
			ClusterID: viper.GetString("nats.cluster"),
			Subj:      viper.GetString("nats.subj"),
		},
		DataBase: &DataBaseConfig{
			ConnLink: viper.GetString("dataBase.pgConnLink"),
		},
	}
}

func init() {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "configs", "Path to config file")
	flag.StringVar(&configFile, "config", "config", "Name of config file")
	flag.StringVar(&configPath, "p", "configs", "Path to config file")
	flag.StringVar(&configFile, "c", "config", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
