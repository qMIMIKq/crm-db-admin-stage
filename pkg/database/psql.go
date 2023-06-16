package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PGConfig struct {
	Host      string `yaml:"host" env:"DB_HOST"`
	Port      string `yaml:"port" env:"DB_PORT"`
	Username  string `yaml:"user" env:"DB_USER"`
	Password  string `yaml:"password" env:"DB_PASSWORD"`
	DBName    string `yaml:"name" env:"DB_NAME"`
	SSLMode   string `yaml:"sslmode" env:"DB_SSLMODE"`
	PgConnect string `yaml:"db_connect" env:"DB_CONNECT"`
}

func NewPostgresDB(cfg PGConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password,
		),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
