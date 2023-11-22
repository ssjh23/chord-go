package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SuccessorAddress   string `mapstructure:"SUCCESSOR_ADDRESS"`
	ServerAddress      string `mapstructure:"SERVER_ADDRESS"`
	PredecessorAddress string `mapstructure:"PREDECESSOR_ADDRESS"`
	ChordId            string `mapstructure:"CHORD_ID"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func ChangeEnvVariable(path string, variableName string, newValue string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	viper.Set(variableName, newValue)
	if err := viper.WriteConfig(); err != nil {
		log.Fatal("Error writing to app.env file")
	}
	return
}
