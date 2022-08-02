package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wbso/dtsgotask/store"
)

func run() error {
	_, err := store.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
