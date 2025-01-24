package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/dundudun/rest_test_back/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (handler *Handler) CreateWasteStorage(c *gin.Context) {
	var requestBody WasteStorageCreate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
	}

	params := sqlc.CreateWasteStorageParams{
		Name:           requestBody.Name,
		PlasticLimit:   utils.OptionalInt4(requestBody.PlasticLimit),
		GlassLimit:     utils.OptionalInt4(requestBody.GlassLimit),
		BiowasteLimit:  utils.OptionalInt4(requestBody.BiowasteLimit),
		StoredPlastic:  utils.OptionalInt4(requestBody.StoredPlastic),
		StoredGlass:    utils.OptionalInt4(requestBody.StoredGlass),
		StoredBiowaste: utils.OptionalInt4(requestBody.StoredBiowaste),
	}
	wasteStorage, err := handler.Queries.CreateWasteStorage(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create waste storage:\n%v", err)})
		return
	}

	c.JSON(http.StatusCreated, wasteStorage)
}

func (handler *Handler) GetWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid waste storage ID:\n%v", err)})
		return
	}
	wasteStorage, err := handler.Queries.GetWasteStorage(handler.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get waste storage:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, wasteStorage)
}

func (handler *Handler) ListWasteStorages(c *gin.Context) {
	wasteStorages, err := handler.Queries.ListWasteStorage(handler.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get waste storages:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, wasteStorages)
}

func (handler *Handler) PartlyChangeWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid waste storage ID"})
		return
	}

	var requestBody WasteStoragePartlyUpdate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.PartlyUpdateWasteStorageParams{
		ID:             id,
		Name:           utils.OptionalText(requestBody.Name),
		PlasticLimit:   utils.OptionalInt4(requestBody.PlasticLimit),
		GlassLimit:     utils.OptionalInt4(requestBody.GlassLimit),
		BiowasteLimit:  utils.OptionalInt4(requestBody.BiowasteLimit),
		StoredPlastic:  utils.OptionalInt4(requestBody.StoredPlastic),
		StoredGlass:    utils.OptionalInt4(requestBody.StoredGlass),
		StoredBiowaste: utils.OptionalInt4(requestBody.StoredBiowaste),
	}
	wasteStorage, err := handler.Queries.PartlyUpdateWasteStorage(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update waste storage"})
		return
	}

	c.JSON(http.StatusOK, wasteStorage)
}

func (handler *Handler) ChangeWasteStorage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid waste storage ID"})
		return
	}

	var requestBody WasteStorageUpdate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.UpdateWasteStorageParams{
		ID:             id,
		Name:           requestBody.Name,
		PlasticLimit:   requestBody.PlasticLimit,
		GlassLimit:     requestBody.GlassLimit,
		BiowasteLimit:  requestBody.BiowasteLimit,
		StoredPlastic:  requestBody.StoredPlastic,
		StoredGlass:    requestBody.StoredGlass,
		StoredBiowaste: requestBody.StoredBiowaste,
	}
	wasteStorage, err := handler.Queries.UpdateWasteStorage(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage:\n%v", err)})
		return
	}

	c.JSON(http.StatusOK, wasteStorage)
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
