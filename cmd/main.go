package main

import (
	"config"
	"fmt"
)

func main() {
	if err := config.LoadConfig("json", "./conf.json"); err != nil {
		panic(fmt.Errorf("cannot open config file %s", err))
	}
	if err := config.LoadConfig("yaml", "./conf.yaml"); err != nil {
		panic(fmt.Errorf("cannot open config file %s", err))
	}
}
