package repositories

import (
	"context"
	"time"

	"github.com/danysoftdev/p-go-create/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Permite inyectar la colección desde fuera (ideal para pruebas)
func SetCollection(c *mongo.Collection) {
	collection = c
}

// InsertarPersona guarda una nueva persona en la base de datos
func InsertarPersona(persona models.Persona) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, persona)
	return err
}

// ObtenerPersonaPorDocumento busca una persona por su Documento
func ObtenerPersonaPorDocumento(documento string) (models.Persona, error) {
	var persona models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"documento": documento}).Decode(&persona)
	return persona, err
}

// ObtenerTodasPersonas devuelve todas las personas en la colección
func ObtenerTodasPersonas() ([]models.Persona, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Persona
	for cursor.Next(ctx) {
		var p models.Persona
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}

// ActualizarPersona actualiza campos de una persona por documento usando un mapa (actualización parcial)
func ActualizarPersona(documento string, update map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"documento": documento}
	updateDoc := bson.M{"$set": update}
	_, err := collection.UpdateOne(ctx, filter, updateDoc, &options.UpdateOptions{})
	return err
}

// EliminarPersona elimina físicamente la persona por documento
func EliminarPersona(documento string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"documento": documento})
	return err
}

type RealPersonaRepository struct{}

func (r RealPersonaRepository) InsertarPersona(p models.Persona) error {
	return InsertarPersona(p)
}

func (r RealPersonaRepository) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	return ObtenerPersonaPorDocumento(doc)
}

func (r RealPersonaRepository) ObtenerTodasPersonas() ([]models.Persona, error) {
	return ObtenerTodasPersonas()
}

func (r RealPersonaRepository) ActualizarPersona(documento string, update map[string]interface{}) error {
	return ActualizarPersona(documento, update)
}

func (r RealPersonaRepository) EliminarPersona(documento string) error {
	return EliminarPersona(documento)
}
