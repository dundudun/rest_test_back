package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dundudun/rest_test_back/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) CreateWasteStorage(c *gin.Context) {
	var req struct {
		Name           string      `json:"name" binding:"required"`
		PlasticLimit   pgtype.Int4 `json:"plastic_limit"`
		GlassLimit     pgtype.Int4 `json:"glass_limit"`
		BiowasteLimit  pgtype.Int4 `json:"biowaste_limit"`
		StoredPlastic  pgtype.Int4 `json:"stored_plastic"`
		StoredGlass    pgtype.Int4 `json:"stored_glass"`
		StoredBiowaste pgtype.Int4 `json:"stored_biowaste"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
	}

	params := sqlc.CreateWasteStorageParams{
		Name:           req.Name,
		PlasticLimit:   req.PlasticLimit,
		GlassLimit:     req.GlassLimit,
		BiowasteLimit:  req.BiowasteLimit,
		StoredPlastic:  req.StoredPlastic,
		StoredGlass:    req.StoredGlass,
		StoredBiowaste: req.StoredBiowaste,
	}
	if err := h.Queries.CreateWasteStorage(h.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create waste storage:\n%v", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid waste storage ID:\n%v", err)})
		return
	}
	stor, err := h.Queries.GetWasteStorage(h.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get waste storage:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, stor)
}

func (h *Handler) ListWasteStorages(c *gin.Context) {
	stors, err := h.Queries.ListWasteStorage(h.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get waste storages:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, stors)
}

func (h *Handler) PartlyChangeWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid waste storage ID"})
		return
	}

	var req struct {
		Name           *string      `json:"name"`
		PlasticLimit   *pgtype.Int4 `json:"plastic_limit"`
		GlassLimit     *pgtype.Int4 `json:"glass_limit"`
		BiowasteLimit  *pgtype.Int4 `json:"biowaste_limit"`
		StoredPlastic  *pgtype.Int4 `json:"stored_plastic"`
		StoredGlass    *pgtype.Int4 `json:"stored_glass"`
		StoredBiowaste *pgtype.Int4 `json:"stored_biowaste"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.PartlyUpdateWasteStorageParams{
		ID:             id,
		Name:           *req.Name,
		PlasticLimit:   *req.PlasticLimit,
		GlassLimit:     *req.GlassLimit,
		BiowasteLimit:  *req.BiowasteLimit,
		StoredPlastic:  *req.StoredPlastic,
		StoredGlass:    *req.StoredGlass,
		StoredBiowaste: *req.StoredBiowaste,
	}
	if _, err := h.Queries.PartlyUpdateWasteStorage(h.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update waste storage"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) ChangeWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid waste storage ID"})
		return
	}

	var req struct {
		Name           string      `json:"name" binding:"required"`
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
	if err := h.Queries.UpdateWasteStorage(h.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage:\n%v", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid waste storage ID:\n%v", err)})
		return
	}
	if err := h.Queries.DeleteWasteStorage(h.Ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete waste storage:\n%v", err)})
		return
	}
	c.Status(http.StatusOK)
}
