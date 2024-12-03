package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func CloseDBWithSysCall(db *sql.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		fmt.Println()
		log.Printf("Received syscall: %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Success: DB connection closed")
		os.Exit(0)
	}()
}