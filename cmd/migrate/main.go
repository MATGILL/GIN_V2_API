package main

import (
	"log"
	"os"

	"github.com/MATGILL/GIN_V2/config"
	"github.com/MATGILL/GIN_V2/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbConfig := db.DbConfig{
		Username: config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Host:     config.Envs.DBHost,
		Port:     config.Envs.DBPort,
		DBName:   config.Envs.DBName,
	}
	db, err := db.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
