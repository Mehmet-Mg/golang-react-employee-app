package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		ConnectionString string
		Port             string
	}
)

func NewConfig(environment string) (cfg *Config, err error) {
	cfg = &Config{}
	viper.SetConfigName(fmt.Sprintf("config.%s", environment))
	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
	}
	if viper.Unmarshal(cfg) != nil {
		err = fmt.Errorf("Config Error")
	} else {
		fmt.Println("Config loaded succesfully", cfg)
	}
	return
}
