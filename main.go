package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	setRuntimeConfig()
}

func main() {
	fmt.Println("Hello Go !!")
}

func setRuntimeConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("Config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
