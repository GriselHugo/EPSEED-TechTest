package main

import (
	// "fmt"
	// "log"
	"main/internal/config"
	"main/internal/server"
)

func main() {
	// Récupération du serveur initialisé
	// server, err := server.InitializeServer()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Démarrage du serveur
	// port := 8080
	// // addr := fmt.Sprintf(":%d", port)
	// fmt.Printf("Serveur écoutant sur le port %d...\n", port)

	// err = server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println("Erreur:", err)
	// }
	config.LoadEnv()
	server.InitializeServer()
}