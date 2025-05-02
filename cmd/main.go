package main

import (
	"log"

	"github.com/MATGILL/GIN_V2/cmd/api"
	"github.com/MATGILL/GIN_V2/config"
	"github.com/MATGILL/GIN_V2/db"
)

func main() {
	//Database configuration
	dbConfig := db.DbConfig{
		Username: config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Host:     config.Envs.PublicHost,
		Port:     config.Envs.Port,
		DBName:   config.Envs.DBName,
	}
	db, err := db.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
