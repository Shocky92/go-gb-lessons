package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	ConfigVar string
}

func LoadConfig(configType string, configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType(configType)
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to load config: %s", err)
	}

	return v.Unmarshal(&Config)
}
