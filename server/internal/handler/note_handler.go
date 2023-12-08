package handler

import (
	// "fmt"
	"net/http"
	// "github.com/gorilla/mux"

	"main/internal/db"
	"encoding/json"
)

type NotesRequest struct {
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

type NotesWithIdRequest struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var noteRequest NotesRequest

	err := json.NewDecoder(r.Body).Decode(&noteRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	err = db.CreateNote(noteRequest.UserId, noteRequest.Title, noteRequest.Content)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de l'ajout de la note")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Note ajoutée",
		Token: "",
	})

	w.Write(returnJson)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var deleteRequest OnlyId

	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	err = db.DeleteNoteByID(deleteRequest.Id)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la suppression de la note")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Note supprimée",
		Token: "",
	})

	w.Write(returnJson)
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var noteRequest NotesWithIdRequest

	err := json.NewDecoder(r.Body).Decode(&noteRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	err = db.UpdateNote(noteRequest.Id, noteRequest.Title, noteRequest.Content)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la mise à jour de la note")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Note mise à jour",
		Token: "",
	})

	w.Write(returnJson)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var getRequest OnlyId

	err := json.NewDecoder(r.Body).Decode(&getRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	note, err := db.GetNoteByID(getRequest.Id)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la récupération de la note")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(note)

	w.Write(returnJson)
}

func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	// var getRequest OnlyId

	// err := json.NewDecoder(r.Body).Decode(&getRequest)
	// if err != nil {
	// 	writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
	// 	return
	// }

	notes, err := db.GetAllNotes()
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la récupération des notes")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(notes)

	w.Write(returnJson)
}

func GetNotesByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var getRequest OnlyId

	err := json.NewDecoder(r.Body).Decode(&getRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	notes, err := db.GetNotesByUserID(getRequest.Id)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la récupération des notes")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(notes)

	w.Write(returnJson)
}


// AddNoteHandler gère la requête pour ajouter une note
// func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour ajouter une note
// 	fmt.Fprintln(w, "Ajout d'une note")
// }

// // AddNoteHandler gère la requête pour ajouter une note
// func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Extraction des données du corps de la requête
// 	var noteData struct {
// 		UserID   int    `json:"user_id"`
// 		Title    string `json:"title"`
// 		Content  string `json:"content"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&noteData)
// 	if err != nil {
// 		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// Vérification des données obligatoires
// 	if noteData.UserID == 0 || noteData.Title == "" || noteData.Content == "" {
// 		http.Error(w, "Données incomplètes", http.StatusBadRequest)
// 		return
// 	}

// 	// Appel de la fonction AddNote pour ajouter la note
// 	err = db.AddNote(dbConnection, noteData.UserID, noteData.Title, noteData.Content)
// 	if err != nil {
// 		http.Error(w, "Erreur lors de l'ajout de la note", http.StatusInternalServerError)
// 		return
// 	}

// 	// Réponse réussie
// 	fmt.Fprintln(w, "Ajout d'une note avec succès")
// }

// DeleteNoteHandler gère la requête pour supprimer une note
// func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour supprimer une note
// 	vars := mux.Vars(r)
// 	noteID := vars["id"]
// 	fmt.Fprintf(w, "Suppression de la note avec l'ID %s\n", noteID)
// }

// UpdateNoteHandler gère la requête pour mettre à jour une note
// func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour mettre à jour une note
// 	vars := mux.Vars(r)
// 	noteID := vars["id"]
// 	fmt.Fprintf(w, "Mise à jour de la note avec l'ID %s\n", noteID)
// }

// GetNoteHandler gère la requête pour récupérer une note
// func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour récupérer une note
// 	vars := mux.Vars(r)
// 	noteID := vars["id"]
// 	fmt.Fprintf(w, "Récupération de la note avec l'ID %s\n", noteID)
// }

// GetAllNotesHandler gère la requête pour récupérer toutes les notes
// func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour récupérer toutes les notes
// 	fmt.Fprintln(w, "Récupération de toutes les notes")
// }

// GetNotesByUserHandler gère la requête pour récupérer toutes les notes d'un utilisateur
// func GetNotesByUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Logique pour récupérer toutes les notes d'un utilisateur
// 	vars := mux.Vars(r)
// 	userID := vars["id"]
// 	fmt.Fprintf(w, "Récupération de toutes les notes de l'utilisateur avec l'ID %s\n", userID)
// }
