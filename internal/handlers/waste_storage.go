package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dundudun/rest_test_back/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *Handler) CreateWasteStorage(c *gin.Context) {
	var requestBody struct {
		Name           pgtype.Text  `json:"name" binding:"required"`
		PlasticLimit   *pgtype.Int4 `json:"plastic_limit"`
		GlassLimit     *pgtype.Int4 `json:"glass_limit"`
		BiowasteLimit  *pgtype.Int4 `json:"biowaste_limit"`
		StoredPlastic  *pgtype.Int4 `json:"stored_plastic"`
		StoredGlass    *pgtype.Int4 `json:"stored_glass"`
		StoredBiowaste *pgtype.Int4 `json:"stored_biowaste"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
	}

	params := sqlc.CreateWasteStorageParams{
		Name:           requestBody.Name,
		PlasticLimit:   *requestBody.PlasticLimit,
		GlassLimit:     *requestBody.GlassLimit,
		BiowasteLimit:  *requestBody.BiowasteLimit,
		StoredPlastic:  *requestBody.StoredPlastic,
		StoredGlass:    *requestBody.StoredGlass,
		StoredBiowaste: *requestBody.StoredBiowaste,
	}
	if err := handler.Queries.CreateWasteStorage(handler.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create waste storage:\n%v", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (handler *Handler) GetWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid waste storage ID:\n%v", err)})
		return
	}
	storage, err := handler.Queries.GetWasteStorage(handler.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get waste storage:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, storage)
}

func (handler *Handler) ListWasteStorages(c *gin.Context) {
	storages, err := handler.Queries.ListWasteStorage(handler.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get waste storages:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, storages)
}

func (handler *Handler) PartlyChangeWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid waste storage ID"})
		return
	}

	var requestBody struct {
		Name           *pgtype.Text `json:"name"`
		PlasticLimit   *pgtype.Int4 `json:"plastic_limit"`
		GlassLimit     *pgtype.Int4 `json:"glass_limit"`
		BiowasteLimit  *pgtype.Int4 `json:"biowaste_limit"`
		StoredPlastic  *pgtype.Int4 `json:"stored_plastic"`
		StoredGlass    *pgtype.Int4 `json:"stored_glass"`
		StoredBiowaste *pgtype.Int4 `json:"stored_biowaste"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.PartlyUpdateWasteStorageParams{
		ID:             id,
		Name:           *requestBody.Name,
		PlasticLimit:   *requestBody.PlasticLimit,
		GlassLimit:     *requestBody.GlassLimit,
		BiowasteLimit:  *requestBody.BiowasteLimit,
		StoredPlastic:  *requestBody.StoredPlastic,
		StoredGlass:    *requestBody.StoredGlass,
		StoredBiowaste: *requestBody.StoredBiowaste,
	}
	if _, err := handler.Queries.PartlyUpdateWasteStorage(handler.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update waste storage"})
		return
	}

	c.Status(http.StatusOK)
}

func (handler *Handler) ChangeWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid waste storage ID"})
		return
	}

	var req struct {
		Name           pgtype.Text `json:"name" binding:"required"`
		PlasticLimit   pgtype.Int4 `json:"plastic_limit" binding:"required"`
		GlassLimit     pgtype.Int4 `json:"glass_limit" binding:"required"`
		BiowasteLimit  pgtype.Int4 `json:"biowaste_limit" binding:"required"`
		StoredPlastic  pgtype.Int4 `json:"stored_plastic" binding:"required"`
		StoredGlass    pgtype.Int4 `json:"stored_glass" binding:"required"`
		StoredBiowaste pgtype.Int4 `json:"stored_biowaste" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.UpdateWasteStorageParams{
		ID:             id,
		Name:           req.Name,
		PlasticLimit:   req.PlasticLimit,
		GlassLimit:     req.GlassLimit,
		BiowasteLimit:  req.BiowasteLimit,
		StoredPlastic:  req.StoredPlastic,
		StoredGlass:    req.StoredGlass,
		StoredBiowaste: req.StoredBiowaste,
	}
	if err := handler.Queries.UpdateWasteStorage(handler.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage:\n%v", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (handler *Handler) DeleteWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid waste storage ID:\n%v", err)})
		return
	}

	if _, err := handler.Queries.DeleteWasteStorage(handler.Ctx, id); err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("no waste storage with such id: %d", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete waste storage:\n%v", err)})
		return
	}
	c.Status(http.StatusOK)
}
