package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	// mongo
	Mongo_URI = os.Getenv("MONGO_URI")

	// github
	GH_Token = os.Getenv("GH_TOKEN")
)
