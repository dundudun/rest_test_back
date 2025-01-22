package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/dundudun/rest_test_back/internal/handlers"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var server *gin.Engine
var handler = handlers.Handler{}

func init() {
	//TOCHECK: do i need to load .env file if docker-compose will set environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file: %v\n", err)
	}

	DATABASE_URL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		//"127.0.0.1",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	handler.Ctx = context.Background()
	handler.Db, err = pgx.Connect(handler.Ctx, DATABASE_URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	handler.Queries = sqlc.New(handler.Db)
	server = gin.Default()

	//TOTHINK: maybe it's place somewhere else, closer to validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(handlers.OrganizationValidation, handlers.OrganizationCreate{})
		v.RegisterStructValidation(handlers.WasteStorageValidation, handlers.WasteStorageCreate{})
	}
}

func main() {
	defer handler.Db.Close(handler.Ctx)

	router := server.Group("/api")

	organization := router.Group("/organizations")
	{
		organization.POST("", handler.CreateOrganization)
		organization.GET("", handler.ListOrganizations)
		organization.GET("/:id", handler.GetOrganization)
		organization.PUT("/:id", handler.ChangeOrganization)
		organization.PATCH("/:id", handler.PartlyChangeOrganization)
		organization.DELETE("/:id", handler.DeleteOrganization)
		organization.POST("/:id/produce", handler.ProduceWaste)
	}

	storage := router.Group("/waste_storages")
	{
		storage.POST("", handler.CreateWasteStorage)
		storage.GET("", handler.ListWasteStorages)
		storage.GET("/:id", handler.GetWasteStorage)
		storage.PUT("/:id", handler.ChangeWasteStorage)
		storage.PATCH("/:id", handler.PartlyChangeWasteStorage)
		storage.DELETE("/:id", handler.DeleteWasteStorage)
	}

	server.Run(":8080")
}
