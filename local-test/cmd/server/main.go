package main

import (
	"fmt"
	"local-test/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println(cfg.FirebaseAuthClient)
}