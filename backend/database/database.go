package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DBConn *pgxpool.Pool
)

func init() {
	var err error
	DBConn, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to create database connection: %v\n", err)
	}
}
