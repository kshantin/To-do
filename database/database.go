package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDB() {
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatal("Database: .env is required but not set")
	}

	db, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Database: unable to connect to database: %v\n", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable ping to database: %v\n", err)
	}
	fmt.Print("Database connected!")
}

func GetDB() *pgxpool.Pool {
	return db
}
