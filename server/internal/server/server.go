package server

import (
	"fmt"
	"main/internal/db"
	// "main/internal/handler"
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
	// http.HandleFunc("/", handler.HandlerRoot)

	// globalRouter := routes.NewGlobalRouter()

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

// func InitializeServer() (*http.Server, error) {
// 	// Gestion des routes
// 	http.HandleFunc("/", handler.HandlerGo)

// 	// Connexion à la base de données
// 	dbConnection, err := db.ConnectToDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer dbConnection.Close()
// 	fmt.Println("Connexion à la base de données réussie!")

// 	globalRouter := routes.NewGlobalRouter()

// 	// Configuration du serveur
// 	port := 8080
// 	addr := fmt.Sprintf(":%d", port)

// 	server := &http.Server{
// 		Addr:    addr,
// 		Handler: globalRouter,
// 	}

// 	return server, nil
// }