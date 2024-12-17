package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"github.com/dundudun/rest_test_back/queries"
)

func getOrganization(c *gin.Context) {
	id := c.Param("id")
	//select from db
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file: %v\n", err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	queries := queries.New(conn)

	r := gin.Default()

	org := r.Group("/organizations")
	org.POST("", createOrganization)
	org.GET("", listOrganizations)
	org.GET("/:id", getOrganization)
	org.PUT("/:id", changeOrganization)
	org.PATCH("/:id", partlyChangeOrganization)
	org.DELETE("/:id", deleteOrganization)

	stor := r.Group("/waste_storages")
	stor.POST("", createWasteStorage)
	stor.GET("", listWasteStorages)
	stor.GET("/:id", getWasteStorage)
	stor.PUT("/:id", changeWasteStorage)
	stor.PATCH("/:id", partlyChangeWasteStorage)
	stor.DELETE("/:id", deleteWasteStorage)

	/*
		r.POST("/organizations", addOrganization)
		r.POST("/waste_storages", addWasteStorage)

		r.GET("/organizations", listOrganizations)
		r.GET("/waste_storages", listWasteStorages)
		r.GET("/organizations/:id", getOrganization)
		r.GET("/waste_storages/:id", getWasteStorage)

		r.PUT("/organizations/:id", changeOrganization)
		r.PUT("/waste_storages/:id", changeWasteStorage)
		r.PATCH("/organizations/:id", partlyChangeOrganization)
		r.PATCH("/waste_storages/:id", partlyChangeWasteStorage)

		r.DELETE("/organizations/:id", deleteOrganization)
		r.DELETE("/waste_storages/:id", deleteWasteStorage)
	*/

	r.Run(":8080")
}
