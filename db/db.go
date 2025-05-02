package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewPostgresDB(conf DbConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName))
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
