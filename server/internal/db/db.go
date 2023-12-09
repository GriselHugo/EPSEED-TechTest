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
