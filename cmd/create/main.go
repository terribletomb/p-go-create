package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danysoftdev/p-go-create/config"
	"github.com/danysoftdev/p-go-create/controllers"
	"github.com/danysoftdev/p-go-create/repositories"
	"github.com/danysoftdev/p-go-create/services"

	"github.com/gorilla/mux"
)

func main() {
	// Debug: imprimir variables de entorno relevantes (ocultar password)
	uri := os.Getenv("MONGO_URI")
	db := os.Getenv("MONGO_DB")
	col := os.Getenv("COLLECTION_NAME")
	masked := uri
	// no exponer credenciales en logs si están en la URI; opcional
	log.Printf("ENV MONGO_URI=%s MONGO_DB=%s COLLECTION_NAME=%s", masked, db, col)

	if err := config.ConectarMongo(); err != nil {
		log.Fatal("Error conectando a Mongo: ", err)
	}
	defer config.CerrarMongo()

	services.SetPersonaRepository(repositories.RealPersonaRepository{})
	repositories.SetCollection(config.Collection)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "create service") })
	r.HandleFunc("/crear-personas", controllers.CrearPersona).Methods("POST")

	port := ":8081"
	fmt.Printf("Create service escuchando en %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
