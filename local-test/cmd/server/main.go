package main

import (
	"fmt"
	config "local-test/configs"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println(cfg.FirebaseAuthClient)
}