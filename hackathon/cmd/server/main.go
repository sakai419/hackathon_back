package main

import (
	v1 "local-test/api/v1"
	"local-test/internal/config"
	"local-test/pkg/database"
	"local-test/pkg/firebase"
	"local-test/pkg/utils"
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

	log.Println("Firebase config: ", cfg.FirebaseConfig)

	// Initialize Firebase client
	firebaseClient, err := firebase.InitFirebaseClient(cfg.FirebaseConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Close DB connection on syscall
	utils.CloseDBWithSysCall(db)

	// Setup server
	server := v1.NewServer(db, firebaseClient)

	// Start server
	if err := server.Start(cfg.ServerConfig.Port); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
