package configs

import (
	"os"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	getEnvsFile := os.Getenv("GET_ENV_FILE")
	if getEnvsFile == "true" {
		viper.SetConfigName("")
		viper.SetConfigType("env")
		viper.AddConfigPath(path)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	viper.AutomaticEnv()
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
