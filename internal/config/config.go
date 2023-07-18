package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     Mongo
	Server Server
}

type Mongo struct {
	URI string
	// Username string
	// Password string
	Database string
}

type Server struct {
	Port int
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := new(Config)

	dbURI := os.Getenv("DB_URI")
	database := os.Getenv("DB_DATABASE")
	serverPort := os.Getenv("SERVER_PORT")
	cfg.DB.URI = dbURI
	cfg.DB.Database = database
	cfg.Server.Port, err = strconv.Atoi(serverPort)
	if err != nil {
		log.Fatalf("can't parse server port: %v", err)
	}
	return cfg, nil
}
