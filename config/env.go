package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// mongo
var Mongo_URI = os.Getenv("MONGODB_URI")
var Mongo_DB = os.Getenv("MONGODB_DB")
var Mongo_Collection = os.Getenv("MONGODB_COLLECTION")

// redis
var Redis_Addr = os.Getenv("REDIS_ADDR")
var Redis_Pass = os.Getenv("REDIS_PASS")
