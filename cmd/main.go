package main

import (
	"config"
	"fmt"
)

func main() {
	cfgPath, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Port)
	fmt.Println(cfg.DBUrl)
	fmt.Println(cfg.JaegerUrl)
	fmt.Println(cfg.SentryUrl)
	fmt.Println(cfg.KafkaBroker)
	fmt.Println(cfg.SomeAppId)
	fmt.Println(cfg.SomeAppKey)
}
