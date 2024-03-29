package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Config file error: %w \n", err))
	}
}