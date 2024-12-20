package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/dundudun/rest_test_back/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var server *gin.Engine
var h = handlers.Handler{}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file: %v\n", err)
	}

	DATABASE_URL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	h.Ctx = context.Background()
	h.Db, err = pgx.Connect(h.Ctx, DATABASE_URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	h.Queries = sqlc.New(h.Db)
	server = gin.Default()
}

func main() {
	defer h.Db.Close(context.Background())

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
