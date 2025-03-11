package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DatabaseURL  string
	DatabaseName string
	JWTSecret    string
	Port         string
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier '.env' trouvé.")
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	DatabaseName = os.Getenv("DATABASE_NAME")
	JWTSecret = os.Getenv("JWT_SECRET")
	Port = os.Getenv("PORT")
}
