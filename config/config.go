package config

import "github.com/spf13/viper"

type Config struct {
	DBURI  string `mapstructure:"MONGO_URI"`
	Api    string `mapstructure:"API"`
	DBName string `mapstructure:"MONGO_DB"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
