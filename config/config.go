package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	RegistrationEnabled bool `mapstructure:"REGISTRATION_ENABLED"`
}

var ConfigPath string = "app.yaml"

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(strings.Split(ConfigPath, ".")[0])
	viper.SetConfigType(strings.Split(ConfigPath, ".")[1])

	viper.ReadInConfig()
	viper.SetEnvPrefix("XT")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

}

func GetConfig(key string) interface{} {
	LoadConfig()
	return viper.Get(key)
}
