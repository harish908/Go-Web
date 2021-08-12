package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Config file error: %w \n", err))
	}
}
