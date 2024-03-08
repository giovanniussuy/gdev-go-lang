package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var C *AppConfigurations

type AppConfigurations struct {
	AppName       string `mapstructure:"APP_NAME"`
	LoggingLevel  string `mapstructure:"logging_level"`
	ServerPort    string `mapstructure:"server_port"`
	ServerTimeout string `mapstructure:"server_timeout"`
	ServerMode    string `mapstructure:"server_mode"`
	SomethingHost string `mapstructure:"something_host"`
}

func LoadAppConfig() (*AppConfigurations, error) {
	var appConfigurations AppConfigurations
	configurations := bindConfig()

	if err := configurations.Unmarshal(&appConfigurations); err != nil {
		return &AppConfigurations{}, err
	}

	C = &appConfigurations
	return &appConfigurations, nil
}

func bindConfig() *viper.Viper {
	viperHandler := viper.New()

	bindEnvironmentVairable("APP_NAME", viperHandler)
	bindEnvironmentVairable("SERVER_MODE", viperHandler)

	viperHandler.SetDefault("LOGGING_LEVEL", "debug")
	viperHandler.SetDefault("SERVER_PORT", "8080")
	viperHandler.SetDefault("SERVER_TIMEOUT", "10")
	viperHandler.SetDefault("SOMETHING_HOST", "host.com.br/something")

	bindEnvironmentVairable("LOGGING_LEVEL", viperHandler)
	bindEnvironmentVairable("SERVER_PORT", viperHandler)
	bindEnvironmentVairable("SERVER_TIMEOUT", viperHandler)
	bindEnvironmentVairable("SOMETHING_HOST", viperHandler)

	viperHandler.AutomaticEnv()
	return viperHandler
}

func bindEnvironmentVairable(environmentVariable string, viperHandler *viper.Viper) {
	if err := viperHandler.BindEnv(environmentVariable); err != nil {
		log.Fatal(fmt.Sprintf("Environment Variable %s not found", environmentVariable), err)
	}
}
