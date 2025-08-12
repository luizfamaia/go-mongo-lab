package config

import "os"

type AppConfig struct {
	Port     string
	MongoURI string
}

func Load() AppConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		// local dev default
		mongoURI = "mongodb://localhost:27017"
	}
	return AppConfig{Port: port, MongoURI: mongoURI}
}
