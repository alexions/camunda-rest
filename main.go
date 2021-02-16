package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/alexions/camunda-rest/app"
)

func main() {
	config, err := readConfig()
	if err != nil {
		panic(fmt.Errorf("unable to init config: %s \n", err))
	}
	app, err := app.NewApp(config)
	if err != nil {
		panic(fmt.Errorf("failed to start app: %s \n", err))
	}

	app.Run()
}

func readConfig() (app.AppConfig, error) {
	var config app.AppConfig

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/camunda-rest/")

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
