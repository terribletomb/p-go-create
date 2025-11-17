package services

import (
	"errors"
	"strings"

	"github.com/danysoftdev/p-go-create/models"
	"github.com/danysoftdev/p-go-create/repositories"
)

var Repo repositories.PersonaRepository

func SetPersonaRepository(r repositories.PersonaRepository) {
	Repo = r
}

func ValidarPersona(p models.Persona) error {
	if strings.TrimSpace(p.Documento) == "" {
		return errors.New("el documento no puede estar vacío")
	}
	if strings.TrimSpace(p.Nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	if strings.TrimSpace(p.Apellido) == "" {
		return errors.New("el apellido no puede estar vacío")
	}
	if p.Edad <= 0 {
		return errors.New("la edad debe ser un número entero mayor a 0")
	}
	if strings.TrimSpace(p.Correo) == "" || !strings.Contains(p.Correo, "@") {
		return errors.New("el correo es inválido")
	}
	if strings.TrimSpace(p.Telefono) == "" {
		return errors.New("el teléfono no puede estar vacío")
	}
	if strings.TrimSpace(p.Direccion) == "" {
		return errors.New("la dirección no puede estar vacía")
	}
	return nil
}

func CrearPersona(p models.Persona) error {
	if err := ValidarPersona(p); err != nil {
		return err
	}

	_, err := Repo.ObtenerPersonaPorDocumento(p.Documento)
	if err == nil {
		return errors.New("ya existe una persona con ese documento")
	}

	return Repo.InsertarPersona(p)
}

// ObtenerPersona devuelve una persona por documento
func ObtenerPersona(documento string) (models.Persona, error) {
	return Repo.ObtenerPersonaPorDocumento(documento)
}

// ObtenerTodasPersonas devuelve todas las personas
func ObtenerTodasPersonas() ([]models.Persona, error) {
	return Repo.ObtenerTodasPersonas()
}

// ActualizarPersona actualiza una persona por documento usando un mapa parcial
func ActualizarPersona(documento string, update map[string]interface{}) error {
	// No permitir cambiar el documento
	if _, ok := update["documento"]; ok {
		return errors.New("no está permitido cambiar el documento")
	}
	return Repo.ActualizarPersona(documento, update)
}

// EliminarPersona elimina una persona por documento
func EliminarPersona(documento string) error {
	return Repo.EliminarPersona(documento)
}
