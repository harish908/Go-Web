package main

import (
	"./config"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	config.ConfigFile()
	fmt.Print(viper.Get("ApiInfo.PortalServer"))
}
