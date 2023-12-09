package routes

import (
	"github.com/gorilla/mux"
	"main/internal/handler"

	"fmt"
	"net/http"
)

func logCallMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Appel de la route %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// IdentificationRouter crée un routeur pour l'identification
func IdentificationRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", handler.LogInHandler).Methods("POST")
	router.HandleFunc("/signup", handler.SignUpHandler).Methods("POST")
	return router
}

// NewUserRouter crée un routeur pour les "user"
func NewUserRouter() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/user/add", handler.AddUserHandler).Methods("POST")
	router.HandleFunc("/delete", handler.DeleteUserHandler).Methods("DELETE")
	// router.HandleFunc("/get/{id}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/getall", handler.GetAllUsersHandler).Methods("GET")
	// Ajoutez d'autres routes pour les "user" au besoin
	return router
}

// NewNoteRouter crée un routeur pour les "note"
func NewNoteRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/add", handler.AddNoteHandler).Methods("POST")
	router.HandleFunc("/delete", handler.DeleteNoteHandler).Methods("DELETE")
	router.HandleFunc("/update", handler.UpdateNoteHandler).Methods("PUT")
	router.HandleFunc("/get/", handler.GetNoteHandler).Methods("GET")
	router.HandleFunc("/getall", handler.GetAllNotesHandler).Methods("GET")
	router.HandleFunc("/getbyuser", handler.GetNotesByUserHandler).Methods("GET")
	// Ajoutez d'autres routes pour les "note" au besoin
	return router
}

// NewGlobalRouter crée un routeur global en incluant les routeurs pour "user" et "note"
func NewGlobalRouter() {
	globalRouter := mux.NewRouter()

	globalRouter.Use(logCallMiddleware)

	// Routeur pour la racine
	globalRouter.HandleFunc("/", handler.HandlerRoot)

	// Routeur pour l'identification
	identificationRouter := IdentificationRouter()
	globalRouter.PathPrefix("").Handler(identificationRouter)

	// Routeur pour les "user"
	userRouter := NewUserRouter()
	globalRouter.PathPrefix("/user").Handler(userRouter)

	// Routeur pour les "note"
	noteRouter := NewNoteRouter()
	globalRouter.PathPrefix("/note").Handler(noteRouter)

	// Ajoutez d'autres routeurs pour d'autres routes au besoin

	http.Handle("/", globalRouter)
}
