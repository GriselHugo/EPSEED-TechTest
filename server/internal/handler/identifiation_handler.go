package handler

import (
	"net/http"
	"main/internal/db"
	"encoding/json"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func writeErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: message,
		Token: "",
	}

	returnJson, _ := json.Marshal(response)
	w.Write(returnJson)
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var loginRequest LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	user, err := db.GetUserByEmailandPassword(loginRequest.Email, loginRequest.Password)
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

	var returnJson, _ = json.Marshal(Response{
		Message: "Connexion réussie",
		Token: "token",
	})

	w.Write(returnJson)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var signupRequest SignupRequest

	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	user, err := db.GetUserByUsername(signupRequest.Username)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la récupération de l'utilisateur")
		return
	}

	if user != nil {
		writeErrorResponse(w, "Utilisateur déjà existant")
		return
	}

	err = db.CreateUser(signupRequest.Username, signupRequest.Password, signupRequest.Email)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la création de l'utilisateur")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Utilisateur créé",
		Token: "token",
	})

	w.Write(returnJson)
}
