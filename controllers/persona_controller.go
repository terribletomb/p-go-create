package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/danysoftdev/p-go-create/models"
	"github.com/danysoftdev/p-go-create/services"
	"github.com/gorilla/mux"
)

func CrearPersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona

	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "El formato del cuerpo es inválido", http.StatusBadRequest)
		return
	}

	err = services.CrearPersona(persona)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Persona creada exitosamente"})
}

// ObtenerPersona devuelve una persona por documento
func ObtenerPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doc := vars["documento"]
	p, err := services.ObtenerPersona(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

// ListarPersonas devuelve todas las personas
func ListarPersonas(w http.ResponseWriter, r *http.Request) {
	ps, err := services.ObtenerTodasPersonas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ps)
}

// ActualizarPersona actualiza una persona por documento
func ActualizarPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doc := vars["documento"]
	var update map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Formato inválido", http.StatusBadRequest)
		return
	}
	if len(update) == 0 {
		http.Error(w, "cuerpo vacío", http.StatusBadRequest)
		return
	}
	if err := services.ActualizarPersona(doc, update); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Actualización exitosa"})
}

// EliminarPersona elimina una persona por documento
func EliminarPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doc := vars["documento"]
	if err := services.EliminarPersona(doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Eliminación exitosa"})
}
