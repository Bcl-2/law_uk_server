package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
    Port  string `mapstructure:"PORT"`
    DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./pkg/common/config/envs")
    viper.SetConfigName("dev")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    if err = viper.ReadInConfig(); err != nil {
        return Config{}, fmt.Errorf("failed to read config file: %w", err)
    }

    if err = viper.Unmarshal(&c); err != nil {
        return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
    }

    return c, nil
}