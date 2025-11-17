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
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "read service") })
	r.HandleFunc("/personas", controllers.ListarPersonas).Methods("GET")
	r.HandleFunc("/personas/{documento}", controllers.ObtenerPersona).Methods("GET")

	port := ":8082"
	fmt.Printf("Read service escuchando en %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

