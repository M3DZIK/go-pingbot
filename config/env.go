package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// mongo
var Mongo_URI = os.Getenv("MONGODB_URI")
var Mongo_DB = os.Getenv("MONGODB_DB")
var Mongo_Collection = os.Getenv("MONGODB_COLLECTION")
