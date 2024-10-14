package config

import (
	"go-simpler.org/env"
)

type Config struct {
	AppConf
	ConsulConf
	PrometheusConf
	RedisConf
	PostgreConf
}

func GetConfig() (*Config, error) {
	var cfg Config

	//Не забыть за коментить!
	//if err := godotenv.Load("C:\\dev\\GameSparks\\go-gs-architecture-team_service\\.env"); err != nil {
	//	return nil, err
	//}

	if err := env.Load(&cfg, nil); err != nil {
		return nil, err
	}

	return &cfg, nil
}
