package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/wbso/dtsgotask/app"
	"github.com/wbso/dtsgotask/restapi"
	"github.com/wbso/dtsgotask/store"
)

func run() error {
	// Connect to the database
	log.Println("Connecting to the database...")
	db, err := store.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		return err
	}
	defer db.Close()

	// initialize repository
	repo := store.New(db)
	// initialize application
	app := app.New(repo)
	// initialize server
	restServer := restapi.New(app)

	srv := &http.Server{
		Handler:           restServer.Router(),
		Addr:              fmt.Sprintf(":%s", os.Getenv("PORT")),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       time.Minute,
	}

	log.Printf("starting server on port %s", os.Getenv("PORT"))
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("error starting server: %s", err)
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
