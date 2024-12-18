package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dundudun/rest_test_back/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

type Handler struct {
	queries *sqlc.Queries
}

func (h *Handler) getOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err) //??
		return
	}
	org, err := h.queries.GetOrganization(ctx, id)
	if err != nil {
		c.Error(err) //??
		//ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, org)
}

type PartlyUpdateOrganizationRequest struct {
	Name             *string      `json:"name"`
	PlasticLimit     *pgtype.Int4 `json:"plastic_limit"`
	GlassLimit       *pgtype.Int4 `json:"glass_limit"`
	BiowasteLimit    *pgtype.Int4 `json:"biowaste_limit"`
	ProducedPlastic  *pgtype.Int4 `json:"produced_plastic"`
	ProducedGlass    *pgtype.Int4 `json:"produced_glass"`
	ProducedBiowaste *pgtype.Int4 `json:"produced_biowaste"`
}

func (h *Handler) partlyChangeOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	var req PartlyUpdateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	org, err := h.queries.PartlyUpdateOrganization(
		c,
		sqlc.PartlyUpdateOrganizationParams{
			ID:               id,
			Name:             *req.Name,
			PlasticLimit:     *req.PlasticLimit,
			GlassLimit:       *req.GlassLimit,
			BiowasteLimit:    *req.BiowasteLimit,
			ProducedPlastic:  *req.ProducedPlastic,
			ProducedGlass:    *req.ProducedGlass,
			ProducedBiowaste: *req.ProducedBiowaste,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update organization"})
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) changeOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	var req struct {
		Name             string      `json:"name" binding:"required"`
		PlasticLimit     pgtype.Int4 `json:"plastic_limit" binding:"required"`
		GlassLimit       pgtype.Int4 `json:"glass_limit" binding:"required"`
		BiowasteLimit    pgtype.Int4 `json:"biowaste_limit" binding:"required"`
		ProducedPlastic  pgtype.Int4 `json:"produced_plastic" binding:"required"`
		ProducedGlass    pgtype.Int4 `json:"produced_glass" binding:"required"`
		ProducedBiowaste pgtype.Int4 `json:"produced_biowaste" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	org, err := h.queries.UpdateOrganization(
		c,
		sqlc.UpdateOrganizationParams{
			ID:               id,
			Name:             req.Name,
			PlasticLimit:     req.PlasticLimit,
			GlassLimit:       req.GlassLimit,
			BiowasteLimit:    req.BiowasteLimit,
			ProducedPlastic:  req.ProducedPlastic,
			ProducedGlass:    req.ProducedGlass,
			ProducedBiowaste: req.ProducedBiowaste,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update organization"})
		return
	}

	c.JSON(http.StatusOK, org)
}

func main() {
	defer pgxConn.Close(context.Background())

	h := Handler{queries}
	r := server.Group("/api")

	org := r.Group("/organizations")
	// org.POST("", createOrganization)
	// org.GET("", listOrganizations)
	org.GET("/:id", h.getOrganization)
	org.PUT("/:id", h.changeOrganization)
	org.PATCH("/:id", h.partlyChangeOrganization)
	// org.DELETE("/:id", deleteOrganization)

	// stor := r.Group("/waste_storages")
	// stor.POST("", createWasteStorage)
	// stor.GET("", listWasteStorages)
	// stor.GET("/:id", getWasteStorage)
	// stor.PUT("/:id", changeWasteStorage)
	// stor.PATCH("/:id", partlyChangeWasteStorage)
	// stor.DELETE("/:id", deleteWasteStorage)

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

	server.Run(":8080")
}
