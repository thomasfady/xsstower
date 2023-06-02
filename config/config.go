package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	RegistrationEnabled bool `mapstructure:"REGISTRATION_ENABLED"`
}

var ConfigPath string

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(strings.Split(ConfigPath, ".")[0])
	viper.SetConfigType(strings.Split(ConfigPath, ".")[1])
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func GetConfig(key string) interface{} {
	viper.AddConfigPath(".")
	viper.SetConfigName(strings.Split(ConfigPath, ".")[0])
	viper.SetConfigType(strings.Split(ConfigPath, ".")[1])

	viper.AutomaticEnv()

	viper.ReadInConfig()
	fmt.Println(viper.AllKeys())
	return viper.Get(key)
}
