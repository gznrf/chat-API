package viper

import (
	"github.com/spf13/viper"
	"log"
)

func SetConfiguration() {
	viper.SetConfigName("home_config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err)
	}
}
