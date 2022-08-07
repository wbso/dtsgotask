package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/wbso/dtsgotask/store"
)

func run(db *pgxpool.Pool) error {
	// TODO: implement
	return nil
}

func prepare() error {
	db, err := store.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return run(db)
}

func main() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := prepare()
	if err != nil {
		log.Fatal(err)
	}
}
