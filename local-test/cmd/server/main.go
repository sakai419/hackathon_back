package main

import (
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/database"
	"log"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Connect to database
	db, err := database.ConnectToDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("connected to db")
	defer db.Close()
}