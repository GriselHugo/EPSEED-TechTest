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

type NoteResponse struct {
	Message string `json:"message"`
	Token 	string `json:"token"`
	Id 	int `json:"id"`
}

type DeleteNoteRequest struct {
	UserId int `json:"user_id"`
	NoteId int `json:"id"`
}

type UpdateNoteRequest struct {
	UserId  int   `json:"user_id"`
	NoteId  int   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddNoteHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var deleteNoteRequest DeleteNoteRequest
	err := json.NewDecoder(r.Body).Decode(&deleteNoteRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.DeleteNoteForUser(deleteNoteRequest.UserId, deleteNoteRequest.NoteId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(NoteResponse{
		Message: "Note supprimée avec succès",
		Token: "token",
		Id: deleteNoteRequest.NoteId,
	})

	w.Write(returnJson)
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updateNoteRequest UpdateNoteRequest
	err := json.NewDecoder(r.Body).Decode(&updateNoteRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	note, err := db.UpdateNoteForUser(updateNoteRequest.UserId, updateNoteRequest.NoteId, updateNoteRequest.Title, updateNoteRequest.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(NoteResponse{
		Message: "Note mise à jour avec succès",
		Token: "token",
		Id: note.ID,
	})

	w.Write(returnJson)
}

func GetNotesByUserHandler(w http.ResponseWriter, r *http.Request) {
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
