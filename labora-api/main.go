package main

import (
	"log"
	"net/http"
	"github.com/SilAronica/Labora-API.git/models"
	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// allEvents va a ser otra forma de decir []event en una struct
type allEvents []event

// Creamos la variable tipo allEvents que va a ser nuestra "db", con un registro
var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(value true)
	  //ruta principal
		router.HandleFunc("/", Index)
	
	  //ruta post
	  router.HandleFunc("/search", models.post).Methods ("POST")
	
	  //ruta get
	  router.HandleFunc("/get", models.get).Methods("GET")
	
		server := http.ListenAndServe(":8080", router)
	  
	  log.Fatal(server)
	}

