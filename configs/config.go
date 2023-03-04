package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	viper.SetConfigFile("./configs/weatherapp.toml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	return viper.GetViper()
}
