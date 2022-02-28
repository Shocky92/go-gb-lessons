package config

type config struct {
	Port        int    `yaml:"port"`
	DBUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}
