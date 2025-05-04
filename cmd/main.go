package main

import (
	"database/sql"
	"log"

	"github.com/MATGILL/GIN_V2/api"
	"github.com/MATGILL/GIN_V2/config"
	"github.com/MATGILL/GIN_V2/db"
)

func main() {
	// Configuration de la base de données
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

	// Initialisation du stockage
	initStorage(db)

	// Création et démarrage du serveur API
	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	// Test de connexion à la base de données
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB : Successfully connected.")
}
