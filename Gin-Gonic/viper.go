package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.Gin-Gonic")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s\n", err)
	}
}