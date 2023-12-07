package server

import (
	"fmt"
	"main/internal/db"
	"main/internal/handler"
	"main/internal/routes"
	"net/http"
)

func InitializeServer() (*http.Server, error) {
	// Gestion des routes
	http.HandleFunc("/", handler.HandlerGo)

	// Connexion à la base de données
	dbConnection, err := db.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer dbConnection.Close()
	fmt.Println("Connexion à la base de données réussie!")

	globalRouter := routes.NewGlobalRouter()

	// Configuration du serveur
	port := 8080
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:    addr,
		Handler: globalRouter,
	}

	return server, nil
}