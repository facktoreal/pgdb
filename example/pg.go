package main

import (
	"context"
	"log"

	"github.com/facktoreal/pgdb"
)

func main() {
	ctx := context.Background()

	cfg := pgdb.Config{
		Hostname: "localhost",
		Username: "stiks",
		Password: "",
		Database: "postgres",
		Port:     "", // optional
		Debug:    false,
	}

	db, err := pgdb.Init(cfg)
	if err != nil {
		log.Fatalf("Unable to connect new database, err: %s", err.Error())

		return
	}

	// check database connection
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Database error: %s", err.Error())

		return
	}

	log.Println("Database connected")
}
