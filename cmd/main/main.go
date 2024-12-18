package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dundudun/rest_test_back/db/sqlc"
	h "github.com/dundudun/rest_test_back/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var (
	server  *gin.Engine
	queries *sqlc.Queries
	pgxConn *pgx.Conn
	ctx     context.Context
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Unable to load .env file: %v\n", err)
	}

	DATABASE_URL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	ctx = context.Background()
	pgxConn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	queries = sqlc.New(pgxConn)
	server = gin.Default()
}

func main() {
	defer pgxConn.Close(context.Background())

	h := h.Handler{Queries: queries, Ctx: ctx}
	r := server.Group("/api")

	org := r.Group("/organizations")
	{
		org.POST("", h.CreateOrganization)
		org.GET("", h.ListOrganizations)
		org.GET("/:id", h.GetOrganization)
		org.PUT("/:id", h.ChangeOrganization)
		org.PATCH("/:id", h.PartlyChangeOrganization)
		org.DELETE("/:id", h.DeleteOrganization)
		org.POST("/:id/produce", h.ProduceWaste)
	}

	stor := r.Group("/waste_storages")
	{
		stor.POST("", h.CreateWasteStorage)
		stor.GET("", h.ListWasteStorages)
		stor.GET("/:id", h.GetWasteStorage)
		stor.PUT("/:id", h.ChangeWasteStorage)
		stor.PATCH("/:id", h.PartlyChangeWasteStorage)
		stor.DELETE("/:id", h.DeleteWasteStorage)
	}

	server.Run(":8080")
}
