package handler

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour ajouter une note
	fmt.Fprintln(w, "Ajout d'une note")
}

// DeleteNoteHandler gère la requête pour supprimer une note
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour supprimer une note
	vars := mux.Vars(r)
	noteID := vars["id"]
	fmt.Fprintf(w, "Suppression de la note avec l'ID %s\n", noteID)
}

// UpdateNoteHandler gère la requête pour mettre à jour une note
func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour mettre à jour une note
	vars := mux.Vars(r)
	noteID := vars["id"]
	fmt.Fprintf(w, "Mise à jour de la note avec l'ID %s\n", noteID)
}

// GetNoteHandler gère la requête pour récupérer une note
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour récupérer une note
	vars := mux.Vars(r)
	noteID := vars["id"]
	fmt.Fprintf(w, "Récupération de la note avec l'ID %s\n", noteID)
}

// GetAllNotesHandler gère la requête pour récupérer toutes les notes
func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour récupérer toutes les notes
	fmt.Fprintln(w, "Récupération de toutes les notes")
}

// GetNotesByUserHandler gère la requête pour récupérer toutes les notes d'un utilisateur
func GetNotesByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour récupérer toutes les notes d'un utilisateur
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "Récupération de toutes les notes de l'utilisateur avec l'ID %s\n", userID)
}
