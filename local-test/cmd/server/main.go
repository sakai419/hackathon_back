package main

import (
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	db, err := database.ConnectToDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	fmt.Println("connected to db")
	defer db.Close()
}