package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"crm-backend/internal/rybakcrm/config"
)

func NewPostgresDb(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.DB.Postgres.Host,
			cfg.DB.Postgres.Port,
			cfg.DB.Postgres.Username,
			cfg.DB.Postgres.DbName,
			cfg.DB.Postgres.Password,
			cfg.DB.Postgres.SslMode,
		),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
