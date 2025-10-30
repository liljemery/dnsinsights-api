package main

import (
	"fmt"
	"log"

	"github.com/jeremyinoa/dnsinsight-api/configs"
	"github.com/jeremyinoa/dnsinsight-api/database"
	"github.com/jeremyinoa/dnsinsight-api/routes"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := database.Init(cfg); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	app := routes.Initialize(cfg)

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
