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
	// Debug: imprimir variables de entorno relevantes
	uri := os.Getenv("MONGO_URI")
	db := os.Getenv("MONGO_DB")
	col := os.Getenv("COLLECTION_NAME")
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		if uri == "" {
			uri = "mongodb://admin:secret@127.0.0.1:27017/?authSource=admin"
			_ = os.Setenv("MONGO_URI", uri)
		}
		if db == "" {
			db = "personas_db"
			_ = os.Setenv("MONGO_DB", db)
		}
		if col == "" {
			col = "personas"
			_ = os.Setenv("COLLECTION_NAME", col)
		}
	}
	log.Printf("ENV MONGO_URI=%s MONGO_DB=%s COLLECTION_NAME=%s", uri, db, col)

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
