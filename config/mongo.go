package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var client *mongo.Client

func ConectarMongo() error {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	collectionName := os.Getenv("COLLECTION_NAME")

	if uri == "" || dbName == "" || collectionName == "" {
		return fmt.Errorf("faltan variables de entorno: MONGO_URI, MONGO_DB o COLLECTION_NAME")
	}

	// Configurable retry policy
	maxAttempts := 15
	delay := 2 * time.Second
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// use a temp client until connection is confirmed
		tmpClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			lastErr = err
			cancel()
			log.Printf("intento %d/%d: mongo.Connect error: %v", attempt, maxAttempts, err)
			if attempt < maxAttempts {
				time.Sleep(delay)
				continue
			}
			break
		}

		// try ping
		err = tmpClient.Ping(ctx, nil)
		cancel()
		if err != nil {
			// disconnect temporary client to avoid leaks
			_ = tmpClient.Disconnect(context.Background())
			lastErr = err
			log.Printf("intento %d/%d: ping a mongo falló: %v", attempt, maxAttempts, err)
			if attempt < maxAttempts {
				time.Sleep(delay)
				continue
			}
			break
		}

		// éxito: asignar cliente global y colección
		client = tmpClient
		Collection = client.Database(dbName).Collection(collectionName)
		log.Println("✅ Conectado a MongoDB correctamente.")
		return nil
	}

	if lastErr != nil {
		return fmt.Errorf("no se pudo conectar a MongoDB después de %d intentos: %w", maxAttempts, lastErr)
	}
	return fmt.Errorf("no se pudo conectar a MongoDB")
}

func CerrarMongo() error {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return client.Disconnect(ctx)
	}
	return nil
}
