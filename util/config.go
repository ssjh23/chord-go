package util

import "github.com/spf13/viper"

type Config struct {
	Predecessor_Address string `mapstructure:"PREDECESSOR_ADDRESS"`
	SuccessorAddress    string `mapstructure:"SUCCESSOR_ADDRESS"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	ChordId             string `mapstructure:"CHORD_ID"`
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
