package main

import (
	"log"
	"net/http"

	"github.com/SilAronica/Labora-API.git/models"
	"github.com/gorilla/mux"
)

func main() {
	//Creamos un nuevo router, usando la libreria mux.
	router := mux.NewRouter().StrictSlash(true)
	//Le agregamos una ruta por defecto al router.
	router.HandleFunc("/", models.HomeLink)
	//Levantamos el servidor de forma normal, pero como handlers le pasamos el router
	log.Fatal(http.ListenAndServe(":8080", router))
}
