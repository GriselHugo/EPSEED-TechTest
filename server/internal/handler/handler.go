package handler

import (
	"fmt"
	"net/http"
)

func HandlerGo(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	fmt.Fprintf(w, "Bonjour, c'est un serveur en Go !")
}
