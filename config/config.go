package config

import (
	"encoding/json"
	"fmt"
	"go-redis-url-shortener/utils/log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name string
	Port string
}

type RedisConfig struct {
	Host string
	Port string
	Pass string
	Db   int
}

var config *Config

type Config struct {
	App   *AppConfig
	Redis *RedisConfig
}

func AllConfig() *Config {
	return config
}

func App() *AppConfig {
	return config.App
}

func Load() {
	config = &Config{}
	loadDefault()

	_ = viper.BindEnv("CONSUL_URL")
	_ = viper.BindEnv("CONSUL_PATH")

	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")

	if consulURL != "" && consulPath != "" {
		_ = viper.AddRemoteProvider("consul", consulURL, consulPath)

		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()

		if err != nil {
			log.Error(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPath))
		}

		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}

		if r, err := json.MarshalIndent(config, "", "  "); err == nil {
			log.Info(string(r))
		}
	} else {
		log.Error("CONSUL_URL or CONSUL_PATH missing! Serving with default config...")
	}
}

func loadDefault() {
	config.App = &AppConfig{
		Name: "url-shortener",
		Port: "1323",
	}

	config.Redis = &RedisConfig{
		Host: "localhost",
		Port: "6379",
		Pass: "",
		Db:   1,
	}
}
