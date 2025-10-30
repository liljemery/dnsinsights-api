package database

import (
	"log"
)

// ApplyMigrations runs database schema migrations.
func ApplyMigrations() {
	if DB == nil {
		log.Println("database not initialized; skipping migrations")
		return
	}
	// Add AutoMigrate calls here, e.g. DB.AutoMigrate(&YourModel{})
	log.Println("migrations completed")
}
