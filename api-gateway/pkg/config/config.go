package config

import "github.com/spf13/viper"

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthSvcUrl        string `mapstructure:"AUTH_SVC_URL"`
	ArticleSvcUrl     string `mapstructure:"ARTICLE_SVC_URL"`
	FileSvcUrl        string `mapstructure:"FILE_SVC_URL"`
	SummaryTextSvcUrl string `mapstructure:"SUMMARY_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
