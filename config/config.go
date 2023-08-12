package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	App
}

type App struct {
	ApiPort int `env:"API_PORT" env-default:"8080"`
}

var instance sync.Once
var configErr error
var appCfg *Config

func ReadConfig() (*Config, error) {
	if appCfg == nil {
		instance.Do(func() {
			appCfg = &Config{}
			if err := cleanenv.ReadEnv(appCfg); err != nil {
				configErr = err
			}
		})
	}

	return appCfg, configErr
}
