package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig() (host, port string) {
	// Set up Viper
	viper.SetConfigName("redis") // Name of config file (without extension)
	viper.AddConfigPath(".")     // Look for config file in the current directory
	viper.SetConfigType("yaml")  // Optional, set config file type. Defaults to "yaml"

	// Read in the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Get values from the config file
	redisHost := viper.GetString("host")
	redisPort := viper.GetInt("port")

	return redisHost, fmt.Sprintf("%d", redisPort)
}
