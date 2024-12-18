package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dundudun/rest_test_back/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) CreateOrganization(c *gin.Context) {
	var req struct {
		Name             string      `json:"name" binding:"required"`
		PlasticLimit     pgtype.Int4 `json:"plastic_limit"`
		GlassLimit       pgtype.Int4 `json:"glass_limit"`
		BiowasteLimit    pgtype.Int4 `json:"biowaste_limit"`
		ProducedPlastic  pgtype.Int4 `json:"produced_plastic"`
		ProducedGlass    pgtype.Int4 `json:"produced_glass"`
		ProducedBiowaste pgtype.Int4 `json:"produced_biowaste"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
	}

	params := sqlc.CreateOrganizationParams{
		Name:             req.Name,
		PlasticLimit:     req.PlasticLimit,
		GlassLimit:       req.GlassLimit,
		BiowasteLimit:    req.BiowasteLimit,
		ProducedPlastic:  req.ProducedPlastic,
		ProducedGlass:    req.ProducedGlass,
		ProducedBiowaste: req.ProducedBiowaste,
	}
	if err := h.Queries.CreateOrganization(h.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create organization:\n%v", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}
	org, err := h.Queries.GetOrganization(h.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get organization:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, org)
}

func (h *Handler) ListOrganizations(c *gin.Context) {
	orgs, err := h.Queries.ListOrganizations(h.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get organizations:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func (h *Handler) PartlyChangeOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	var req struct {
		Name             *string      `json:"name"`
		PlasticLimit     *pgtype.Int4 `json:"plastic_limit"`
		GlassLimit       *pgtype.Int4 `json:"glass_limit"`
		BiowasteLimit    *pgtype.Int4 `json:"biowaste_limit"`
		ProducedPlastic  *pgtype.Int4 `json:"produced_plastic"`
		ProducedGlass    *pgtype.Int4 `json:"produced_glass"`
		ProducedBiowaste *pgtype.Int4 `json:"produced_biowaste"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.PartlyUpdateOrganizationParams{
		ID:               id,
		Name:             *req.Name,
		PlasticLimit:     *req.PlasticLimit,
		GlassLimit:       *req.GlassLimit,
		BiowasteLimit:    *req.BiowasteLimit,
		ProducedPlastic:  *req.ProducedPlastic,
		ProducedGlass:    *req.ProducedGlass,
		ProducedBiowaste: *req.ProducedBiowaste,
	}
	if err := h.Queries.PartlyUpdateOrganization(h.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update organization"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) ChangeOrganization(c *gin.Context) {
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
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.UpdateOrganizationParams{
		ID:               id,
		Name:             req.Name,
		PlasticLimit:     req.PlasticLimit,
		GlassLimit:       req.GlassLimit,
		BiowasteLimit:    req.BiowasteLimit,
		ProducedPlastic:  req.ProducedPlastic,
		ProducedGlass:    req.ProducedGlass,
		ProducedBiowaste: req.ProducedBiowaste,
	}
	if err := h.Queries.UpdateOrganization(h.Ctx, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update organization:\n%v", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}
	if err := h.Queries.DeleteOrganization(h.Ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete organization:\n%v", err)})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) ProduceWaste(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}

	amount, err := strconv.Atoi(c.Query("amount"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("failure of getting 'amount' query parameter:\n%v", err)})
		return
	}
	if amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "amount shouldn't be less or equal to 0"})
		return
	}

	wasteType := c.Query("waste_type")
	switch wasteType {
	case "plastic":
	case "glass":
	case "biowaste":
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong waste type, should be 'plastic', 'glass' or 'biowaste'"})
		return
	}

	type fetchedWasteStorages struct {
		Layer		int
		ID          int64
		Name        string
		WasteLimit  pgtype.Int4
		StoredWaste pgtype.Int4
		DistanceMeters int32
	}

	//var wasteStorages []fetchedWasteStorages
	/*
		query next layer
		work with one that have min distance
			if have free space in storage subract from amount
			if no free space just do nothing, go further
		check amount on zero
			if greater than zero repeat for next layer
			in other case exit loop
	*/
	for amount > 0 {
		fetched, err := h.Queries.FromOrgPlasticStors(h.Ctx, pgtype.Int8{Int64: id, Valid: true})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failure to get waste storages for organization:\n%v", err)})
			return
		}
		for _, stor := range fetched {
			stor.
		}
	}
}
