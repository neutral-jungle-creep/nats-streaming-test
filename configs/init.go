package configs

import (
	"flag"
	"github.com/spf13/viper"
)

func InitConfig() error {
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
			URL:       viper.GetString("URL"),
			ClientID:  viper.GetString("client"),
			ClusterID: viper.GetString("cluster"),
			Subj:      viper.GetString("subj"),
		},
		DataBase: &DataBaseConfig{
			ConnLink: viper.GetString("connLink"),
		},
	}
}
