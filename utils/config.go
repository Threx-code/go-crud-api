package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Environment struct {
	DBName              string        `mapstructure:"DB_NAME"`
	DBHost              string        `mapstructure:"DB_HOST"`
	DBRoot              string        `mapstructure:"DB_ROOT"`
	RootPassword        string        `mapstructure:"ROOT_PASSWOORD"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (env Environment, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&env)
	return
}
