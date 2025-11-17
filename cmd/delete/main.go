package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danysoftdev/p-go-create/config"
	"github.com/danysoftdev/p-go-create/controllers"
	"github.com/danysoftdev/p-go-create/repositories"
	"github.com/danysoftdev/p-go-create/services"

	"github.com/gorilla/mux"
)

func main() {
	if err := config.ConectarMongo(); err != nil {
		log.Fatal("Error conectando a Mongo: ", err)
	}
	defer config.CerrarMongo()

	services.SetPersonaRepository(repositories.RealPersonaRepository{})
	repositories.SetCollection(config.Collection)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "delete service") })
	r.HandleFunc("/personas/{documento}", controllers.EliminarPersona).Methods("DELETE")

	port := ":8084"
	fmt.Printf("Delete service escuchando en %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
