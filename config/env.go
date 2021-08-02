package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	// mongo
	Mongo_URI        = os.Getenv("MONGODB_URI")
	Mongo_DB         = os.Getenv("MONGODB_DB")
	Mongo_Collection = os.Getenv("MONGODB_COLLECTION")

	// jwt
	JWT_Secret = os.Getenv("JWT_SECRET")
)
