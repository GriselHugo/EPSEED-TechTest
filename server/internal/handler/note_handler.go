package handler

import (
	"net/http"
	"main/internal/db"
	"encoding/json"
	"strconv"
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

type NoteResponse struct {
	Message string `json:"message"`
	Token 	string `json:"token"`
	Id 	int `json:"id"`
}

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	var noteRequest NotesRequest

	err := json.NewDecoder(r.Body).Decode(&noteRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de la lecture du formulaire")
		return
	}

	note, err := db.CreateNote(noteRequest.UserId, noteRequest.Title, noteRequest.Content)
	if err != nil {
		writeErrorResponse(w, "Erreur lors de l'ajout de la note")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(NoteResponse{
		Message: "Note ajoutée",
		Token: "",
		Id: note.ID,
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
		Token: "token",
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
		Token: "token",
	})

	w.Write(returnJson)
}

func GetNotesByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	notes, err := db.GetNotesByUserID(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(notes)

	w.Write(returnJson)
}
