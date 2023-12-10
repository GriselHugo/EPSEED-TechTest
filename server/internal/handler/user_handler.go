package handler

import (
	"net/http"
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
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, "Erreur de requête des utilisateurs", http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Erreur d'encodage JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(usersJSON)
}
