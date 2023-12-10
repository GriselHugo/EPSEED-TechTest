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

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", handler.LogInHandler).Methods("POST")
	router.HandleFunc("/signup", handler.SignUpHandler).Methods("POST")
	router.HandleFunc("/getall", handler.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/createnote", handler.AddNoteHandler).Methods("POST")
	router.HandleFunc("/deletenote", handler.DeleteNoteHandler).Methods("DELETE")
	router.HandleFunc("/updatenote", handler.UpdateNoteHandler).Methods("PUT")
	router.HandleFunc("/getnote", handler.GetNotesByUserHandler).Methods("GET")
	return router
}

func NewGlobalRouter() {
	globalRouter := mux.NewRouter()

	globalRouter.Use(logCallMiddleware)

	globalRouter.HandleFunc("/", handler.HandlerRoot)

	Router := Router()
	globalRouter.PathPrefix("").Handler(Router)

	http.Handle("/", globalRouter)
}
