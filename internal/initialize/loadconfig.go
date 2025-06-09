package initialize

import (
	"fmt"
	"go-crm-api-shopdev/global"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("Server Key:", viper.GetString("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode %v", err)
	}

	fmt.Println("Config port:", config.Server.Port)
}
