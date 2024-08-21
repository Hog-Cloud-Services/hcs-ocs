package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("config");
	viper.SetConfigType("yaml");
	viper.AddConfigPath(".");
	viper.AddConfigPath("/etc/ocs");
	setDefaults();
	err := viper.ReadInConfig();
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		log.Warn("LoadConfig: could not locate config file, using default values")
	}
	return err;
}

func setDefaults() {
	viper.SetDefault("port", "80")
}
