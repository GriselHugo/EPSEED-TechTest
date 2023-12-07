package routes

import (
	"github.com/gorilla/mux"
	"main/internal/handler"
)

// NewUserRouter crée un routeur pour l'entité "user"
func NewUserRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/add", handler.AddUserHandler).Methods("POST")
	router.HandleFunc("/user/delete/{id}", handler.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/user/get/{id}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/user/getall", handler.GetAllUsersHandler).Methods("GET")
	// Ajoutez d'autres routes pour l'entité "user" au besoin
	return router
}

// NewNoteRouter crée un routeur pour l'entité "note"
func NewNoteRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/note/add", handler.AddNoteHandler).Methods("POST")
	router.HandleFunc("/note/delete/{id}", handler.DeleteNoteHandler).Methods("DELETE")
	router.HandleFunc("/note/update/{id}", handler.UpdateNoteHandler).Methods("PUT")
	router.HandleFunc("/note/get/{id}", handler.GetNoteHandler).Methods("GET")
	router.HandleFunc("/note/getall", handler.GetAllNotesHandler).Methods("GET")
	router.HandleFunc("/note/getbyuser/{id}", handler.GetNotesByUserHandler).Methods("GET")
	// Ajoutez d'autres routes pour l'entité "note" au besoin
	return router
}

// NewGlobalRouter crée un routeur global en incluant les routeurs pour "user" et "note"
func NewGlobalRouter() *mux.Router {
	globalRouter := mux.NewRouter()

	// Routeur pour l'entité "user"
	userRouter := NewUserRouter()
	globalRouter.PathPrefix("/user").Handler(userRouter)

	// Routeur pour l'entité "note"
	noteRouter := NewNoteRouter()
	globalRouter.PathPrefix("/note").Handler(noteRouter)

	// Ajoutez d'autres routeurs pour d'autres entités au besoin

	return globalRouter
}
