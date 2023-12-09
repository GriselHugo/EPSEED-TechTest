package handler

import (
	"net/http"
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
