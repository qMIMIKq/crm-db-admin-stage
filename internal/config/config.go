package config

import (
	"crm_admin/pkg/database"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"sync"
)

type Config struct {
	AppPort string            `yaml:"app_port" env:"PORT"`
	Salt    string            `yaml:"salt" env:"SALT"`
	CrmDB   database.PGConfig `yaml:"db"`
}

var once sync.Once

func GetConfig() *Config {
	config := &Config{}

	once.Do(func() {
		log.Info().Caller().Msg("reading app config")

		if err := cleanenv.ReadConfig("./settings/dev-config.yml", config); err != nil {
			if err := cleanenv.ReadEnv(config); err != nil {
				help, _ := cleanenv.GetDescription(config, nil)
				if help != "" {
					log.Warn().Caller().Msgf("env helpful info %s", help)
				}

				log.Fatal().Caller().Err(err).Msg("error reading app config")
			}
		}
	})

	return config
}
