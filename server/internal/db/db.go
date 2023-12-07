package db

import (
	"main/internal/config"

	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	// Charger le fichier .env
	config.LoadEnv()

	// Récupérer les variables d'environnement
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Construire la chaîne de connexion à la base de données
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Ouvrir une connexion à la base de données
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Vérifier la connexion à la base de données
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Retourner la connexion à la base de données
	return db, nil
}
