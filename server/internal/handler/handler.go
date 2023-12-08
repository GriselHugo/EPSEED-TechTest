package handler

import (
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Token 	string `json:"token"`
}

type OnlyId struct {
	Id int `json:"id"`
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	// Gestion de la requête
	fmt.Fprintf(w, "Bonjour, c'est un serveur en Go !")
}
