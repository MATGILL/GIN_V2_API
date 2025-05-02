package main

import (
	"database/sql"
	"log"

	"github.com/MATGILL/GIN_V2/api"
	"github.com/MATGILL/GIN_V2/config"
	"github.com/MATGILL/GIN_V2/db"
)

func main() {
	//Database configuration
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

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB : Successfully connected.")
}
