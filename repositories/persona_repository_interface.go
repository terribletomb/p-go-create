package repositories

import "github.com/danysoftdev/p-go-create/models"

type PersonaRepository interface {
	InsertarPersona(persona models.Persona) error
	ObtenerPersonaPorDocumento(documento string) (models.Persona, error)
	ObtenerTodasPersonas() ([]models.Persona, error)
	ActualizarPersona(documento string, update map[string]interface{}) error
	EliminarPersona(documento string) error
}
