package handler

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// AddUserHandler gère la requête pour ajouter un utilisateur
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour ajouter un utilisateur
	fmt.Fprintln(w, "Ajout d'un utilisateur")
}

// DeleteUserHandler gère la requête pour supprimer un utilisateur
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour supprimer un utilisateur
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "Suppression de l'utilisateur avec l'ID %s\n", userID)
}

// GetUserHandler gère la requête pour récupérer un utilisateur
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour récupérer un utilisateur
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "Récupération de l'utilisateur avec l'ID %s\n", userID)
}

// GetAllUsersHandler gère la requête pour récupérer tous les utilisateurs
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour récupérer tous les utilisateurs
	fmt.Fprintln(w, "Récupération de tous les utilisateurs")
}
