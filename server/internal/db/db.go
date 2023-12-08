package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := dbUser + ":" + dbPass + "@tcp(mariadb:" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("Attempting to connect to database...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Println("Connexion à la base de données réussie!")
	return db, nil
}

// func ConnectToDB() (*sql.DB, error) {
// 	// Charger le fichier .env
// 	config.LoadEnv()

// 	// Récupérer les variables d'environnement
// 	dbHost := os.Getenv("DB_HOST")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbName := os.Getenv("DB_NAME")
// 	dbPort := os.Getenv("DB_PORT")

// 	// Construire la chaîne de connexion à la base de données
// 	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

// 	// Ouvrir une connexion à la base de données
// 	db, err := sql.Open("mysql", dbURI)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Vérifier la connexion à la base de données
// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	// Retourner la connexion à la base de données
// 	return db, nil
// }
