package main

import (
	"log"

	"github.com/jeremyinoa/dnsinsight-api/configs"
	"github.com/jeremyinoa/dnsinsight-api/database"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	if err := database.Init(cfg); err != nil {
		log.Fatalf("db: %v", err)
	}
	database.ApplyMigrations()
}
