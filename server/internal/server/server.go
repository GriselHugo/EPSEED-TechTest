package server

import (
	"fmt"
	"main/internal/db"
	"main/internal/routes"
	"net/http"
	"github.com/rs/cors"
)

func InitializeServer() {
	var err error
	db.DbInstance, err = db.ConnectToDB()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	sqlDB, err := db.DbInstance.DB()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer sqlDB.Close()

	// Gestion des routes

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(http.DefaultServeMux)

	routes.NewGlobalRouter()

	http.ListenAndServe(":8080", handler)
	fmt.Println("Serveur lancé sur le port 8080")
}
