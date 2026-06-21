package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/zhalisher/ip-task-manager/config"
	"github.com/zhalisher/ip-task-manager/internal/app"
)

func main() {
	godotenv.Load()

	cfg := config.Load()
	if cfg.PostgresURI == "" {
		log.Fatal("POSTGRES_URI is required")
	}
	app.Run(cfg)
}
