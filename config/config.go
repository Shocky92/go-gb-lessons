package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Port        int    `yaml:"port"`
	DBUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func NewConfig(configPath string) (*config, error) {
	config := &config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidatePath(path string) error {
	stat, err := os.Stat(path)

	if err != nil {
		return err
	}

	if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./conf.yaml", "path to config file")
	flag.Parse()

	if err := ValidatePath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
