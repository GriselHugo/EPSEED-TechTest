package handler

import (
	// "fmt"
	"net/http"
	// "github.com/gorilla/mux"

	"main/internal/db"
	"encoding/json"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gestions de la requête
	var deleteRequest OnlyId

	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	err = db.DeleteUser(deleteRequest.Id)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la suppression de l'utilisateur")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Utilisateur supprimé",
		Token: "",
	})

	w.Write(returnJson)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var getRequest OnlyId

	err := json.NewDecoder(r.Body).Decode(&getRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	user, err := db.GetUserById(getRequest.Id)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la récupération de l'utilisateur")
		return
	}

	if user == nil {
		writeErrorResponse(w, "Utilisateur inconnu")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(user)

	w.Write(returnJson)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	users, err := db.GetUsers()
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la récupération des utilisateurs")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(users)

	w.Write(returnJson)
}


// AddUserHandler gère la requête pour ajouter un utilisateur
// func AddUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour ajouter un utilisateur
// 	fmt.Fprintln(w, "Ajout d'un utilisateur")
// }

// DeleteUserHandler gère la requête pour supprimer un utilisateur
// func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour supprimer un utilisateur
// 	vars := mux.Vars(r)
// 	userID := vars["id"]
// 	fmt.Fprintf(w, "Suppression de l'utilisateur avec l'ID %s\n", userID)
// }

// GetUserHandler gère la requête pour récupérer un utilisateur
// func GetUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour récupérer un utilisateur
// 	vars := mux.Vars(r)
// 	userID := vars["id"]
// 	fmt.Fprintf(w, "Récupération de l'utilisateur avec l'ID %s\n", userID)
// }

// GetAllUsersHandler gère la requête pour récupérer tous les utilisateurs
// func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour récupérer tous les utilisateurs
// 	fmt.Fprintln(w, "Récupération de tous les utilisateurs")
// }
