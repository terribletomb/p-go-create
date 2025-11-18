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
	// Si estamos en GitHub Actions y las variables vienen vacías, usar fallback al servicio job (127.0.0.1)
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
