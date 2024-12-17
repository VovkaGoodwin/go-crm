package main

import (
	"crm-backend/internal/rybakcrm/app/infrastructure/database"
	"crm-backend/internal/rybakcrm/config"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

const (
	up   = "up"
	down = "down"
)

func runMigration() error {
	cfg, _ := config.NewConfig()

	args := os.Args[1:2]
	if len(args) == 0 {
		return errors.New("empty action")
	}

	if args[0] != down && args[0] != up {
		return errors.New("wrong action use aup or down")
	}

	db, err := database.NewPostgresDb(cfg)

	if err != nil {
		return err
	}

	pg, err := postgres.WithInstance(db.DB, &postgres.Config{
		MigrationsTable: cfg.Migrations.Table,
		SchemaName:      cfg.Migrations.SchemaName,
	})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://"+cfg.Migrations.Path,
		cfg.DB.Postgres.DbName,
		pg,
	)
	if err != nil {
		return err
	}

	switch args[0] {
	case down:
		err = migrator.Down()
	case up:
		err = migrator.Up()
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := runMigration()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("all changes has been applied")
}
