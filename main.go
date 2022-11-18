package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SilAronica/labora-proyectoFinal.git/controllers"
	"github.com/SilAronica/labora-proyectoFinal.git/services"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func main() {
	fmt.Println("Bienvenidos!")
	services.InitDb()
	router := mux.NewRouter().StrictSlash(true)
	//ruta principal
	router.HandleFunc("/", controllers.Index)
	//ruta Post
	router.HandleFunc("/search", controllers.CreateEventSearch).Methods("POST")
	//ruta Get
	router.HandleFunc("/get", controllers.GetEvent).Methods("GET")
	server := http.ListenAndServe(":8080", router)
	if server != nil {
		fmt.Println("Erroral iniciar el servidor")
		log.Fatal(server)
	}

	log.Fatal()
}
