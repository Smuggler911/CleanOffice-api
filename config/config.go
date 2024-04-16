package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	DbUrl             string `mapstructure:"DB_URL"`
	ImgPath           string `mapstructure:"IMG_PATH_PROD"`
	Secret            string `mapstructure:"SECRET"`
	TwilioAccountSid  string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToke    string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioPhoneNumber string `mapstructure:"TWILIO_PHONE_NUMBER"`
}

func LoadConfig() (config Config, err error) {

	viper.SetConfigFile("././.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
