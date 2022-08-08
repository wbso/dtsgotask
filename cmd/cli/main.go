package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/wbso/dtsgotask/app"
	"github.com/wbso/dtsgotask/store"
)

func run(db *pgxpool.Pool) error {
	a := app.New(store.New(db))
	task, err := a.CreateTask(context.Background(), app.CreateTaskRequest{
		Detail:   "test",
		Assignee: "test",
		Deadline: "2020-01-01",
	})
	if err != nil {
		return err
	}

	datastring, _ := json.Marshal(task)
	fmt.Println("DATA:")
	fmt.Println(string(datastring))
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
