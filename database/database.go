package database

import (
	"context"
	"log"

	"github.com/StevenYAMBOS/wilo-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
    clientOptions := options.Client().ApplyURI(config.DatabaseURL)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Erreur de connexion à la base de données:", err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal("Base de données non accessible:", err)
    }

    log.Println("Connexion à la base de données réussie !")
    DB = client.Database(config.DatabaseName)
}
