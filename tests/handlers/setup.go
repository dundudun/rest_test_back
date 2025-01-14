package handlers_tests

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dundudun/rest_test_back/internal/handlers"
	"github.com/jackc/pgx/v5"
)

var handler = handlers.Handler{}

func SetupDB() func() {
	DATABASE_URL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	handler.Ctx = context.Background()
	var err error
	if handler.Db, err = pgx.Connect(handler.Ctx, DATABASE_URL); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return func() { handler.Db.Close(handler.Ctx) }
}
