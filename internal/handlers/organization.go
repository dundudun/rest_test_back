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
	if _, err := h.Queries.PartlyUpdateOrganization(h.Ctx, params); err != nil {
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
	overProduce := int32(0)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}
	org, err := h.Queries.GetOrganization(h.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to get organization from db:\n%v", err)})
		return
	}

	amount64, err := strconv.ParseInt(c.Param("amount"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failure of getting 'amount' query parameter:\n%v", err)})
		return
	}
	amount := int32(amount64)
	if amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount shouldn't be less or equal to 0"})
		return
	}

	//TODO: flexibily for waste_type
	wasteType := c.Query("waste_type")
	switch wasteType {
	case "plastic":
	case "glass":
	case "biowaste":
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong waste type, should be 'plastic', 'glass' or 'biowaste'"})
		return
	}

	/*
		query next layer (1st layer from Organizations is outside loop)
		add wasteStorages from new layer to all wasteStorages (for ones from wasteStorage accumulate Distance with parent)
		find across all wasteStorages that we haven't worked with one that have min distance
		mark it somehow or pop it?
		work with one that have min distance
			if have free space in storage subract from amount
			if no free space just do nothing, go further
		check amount on zero
			if greater than zero repeat for next layer
			in other case exit loop
	*/

	//TODO: try slice first, then see if map[layer]fetchedWasteStorages would be better
	var wasteStorages []fetchedWasteStorages
	var updatedWasteStorages []sqlc.WasteStorage

	layer := 1
	// 1st layer (one next to Organization)
	fmt.Printf("going to fetch from Organization: %d\n", id)
	fetched, err := h.Queries.FromOrgPlasticStors(h.Ctx, pgtype.Int8{Int64: id, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to get waste storages for organization:\n%v", err)})
		return
	}
	if len(fetched) == 0 {
		//TOCHECK: query to update Organization produced_wastetype value if there isn't WasteStorages for it
		//TOCHECK: overProduce
		org := h.addWasteToOrg(c, &org, amount)
		if org == nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":                        "no wasteStorages for organization",
			"organization":                   org,
			"waste_storages":                 updatedWasteStorages,
			"waste_above_organization_limit": overProduce,
		})
		return
	}
	for _, stor := range fetched {
		wasteStorages = append(wasteStorages, fetchedWasteStorages{
			WorkedWith:     false,
			Layer:          layer,
			ID:             stor.ID,
			Name:           stor.Name,
			WasteLimit:     stor.PlasticLimit,
			StoredWaste:    stor.StoredPlastic,
			DistanceMeters: stor.DistanceMeters,
		})
	}
	idClosest := findClosestWasteStorage(wasteStorages)
	wasteStorages[idClosest].WorkedWith = true

	freeSpaceInStore := wasteStorages[idClosest].WasteLimit.Int32 - wasteStorages[idClosest].StoredWaste.Int32
	if freeSpaceInStore > 0 {
		var newStoredWaste int32
		if freeSpaceInStore >= amount {
			amount = 0
			newStoredWaste = wasteStorages[idClosest].StoredWaste.Int32 + amount
		} else {
			amount -= freeSpaceInStore
			newStoredWaste = wasteStorages[idClosest].WasteLimit.Int32
		}

		params := sqlc.PartlyUpdateWasteStorageParams{
			ID:            wasteStorages[idClosest].ID,
			StoredPlastic: pgtype.Int4{Int32: newStoredWaste, Valid: true},
		}
		updatedWasteStorage, err := h.Queries.PartlyUpdateWasteStorage(h.Ctx, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage's value of stored_plastic:\n%v", err)})
			return
		}
		updatedWasteStorages = append(updatedWasteStorages, updatedWasteStorage)
	}

	// all next layers (ones after WasteStorages)
	nextLayer := true
	for amount > 0 {
		//TOCHECK: to not get in this loop if waste was distributed on 1st layer
		//TOCHECK: exiting right when produced waste is distributed and no more db calls or loops
		//TOCHECK: to correctly get new closest wasteStorage each time and skip ones that we looked up already
		//TOCHECK: if wasteStorages are full but amount of waste still left
		//TOCHECK: do not call db if there isn't more connections
		//TOCHECK: to correctly update/distribute values across wasteStorages
		//TODO: remore this if statement after testing
		if !nextLayer {
			fmt.Println("no WasteStorages have connections to another ones")
		} else if nextLayer {
			layer++
			counter := 0
			// fetch next layer from previous one
			for _, wasteStorage := range wasteStorages {
				if wasteStorage.Layer-1 != layer {
					counter++
					continue
				}
				fmt.Printf("going to fetch from WasteStorage: %d, %s\n", wasteStorage.ID, wasteStorage.Name)
				fetched, err := h.Queries.FromStorsPlasticStors(h.Ctx, pgtype.Int8{Int64: wasteStorage.ID, Valid: true})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to get waste storages for waste storage:\n%v", err)})
					return
				}
				if len(fetched) == 0 {
					counter++
					continue
				}
				for _, stor := range fetched {
					wasteStorages = append(wasteStorages, fetchedWasteStorages{
						WorkedWith:     false,
						Layer:          layer,
						ID:             stor.ID,
						Name:           stor.Name,
						WasteLimit:     stor.PlasticLimit,
						StoredWaste:    stor.StoredPlastic,
						DistanceMeters: stor.DistanceMeters + wasteStorage.DistanceMeters,
					})
				}
			}

			if counter == len(wasteStorages) {
				nextLayer = false
			}
		}

		idClosest = findClosestWasteStorage(wasteStorages)
		if idClosest == -1 {
			//TOCHECK: update Organization produced_wastetype value and list of WasteStorages
			org := h.addWasteToOrg(c, &org, amount)
			if org == nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":                        "no more wasteStorages available to receive waste",
				"organization":                   org,
				"waste_storages":                 updatedWasteStorages,
				"waste_above_organization_limit": overProduce,
			})
			return
		}

		wasteStorages[idClosest].WorkedWith = true

		freeSpaceInStore := wasteStorages[idClosest].WasteLimit.Int32 - wasteStorages[idClosest].StoredWaste.Int32
		if freeSpaceInStore > 0 {
			var newStoredWaste int32
			if freeSpaceInStore >= amount {
				amount = 0
				newStoredWaste = wasteStorages[idClosest].StoredWaste.Int32 + amount
			} else {
				amount -= freeSpaceInStore
				newStoredWaste = wasteStorages[idClosest].WasteLimit.Int32
			}

			params := sqlc.PartlyUpdateWasteStorageParams{
				ID:            wasteStorages[idClosest].ID,
				StoredPlastic: pgtype.Int4{Int32: newStoredWaste, Valid: true},
			}
			updatedWasteStorage, err := h.Queries.PartlyUpdateWasteStorage(h.Ctx, params)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage's value of stored_plastic:\n%v", err)})
				return
			}
			updatedWasteStorages = append(updatedWasteStorages, updatedWasteStorage)
		}
	}

	//TOCHECK: final output
	c.JSON(http.StatusOK, gin.H{
		"message":                        "waste was fully distributed across waste storages",
		"organization":                   org,
		"waste_storages":                 updatedWasteStorages,
		"waste_above_organization_limit": overProduce,
	})
}

type fetchedWasteStorages struct {
	WorkedWith     bool
	Layer          int
	ID             int64
	Name           string
	WasteLimit     pgtype.Int4
	StoredWaste    pgtype.Int4
	DistanceMeters int32
}

func findClosestWasteStorage(wasteStorages []fetchedWasteStorages) int {
	idStorMinDist := -1

	for i, wasteStorage := range wasteStorages {
		if wasteStorage.WorkedWith {
			continue
		}

		if idStorMinDist == -1 {
			idStorMinDist = i
		} else if wasteStorage.DistanceMeters < wasteStorages[idStorMinDist].DistanceMeters {
			idStorMinDist = i
		}
	}

	return idStorMinDist
}

func (h *Handler) addWasteToOrg(c *gin.Context, org *sqlc.Organization, amount int32) *sqlc.Organization {
	overProduce := (org.ProducedPlastic.Int32 + amount) % org.PlasticLimit.Int32
	newValue := (org.ProducedPlastic.Int32 + amount) - overProduce

	params := sqlc.PartlyUpdateOrganizationParams{ID: org.ID, ProducedPlastic: pgtype.Int4{Int32: newValue, Valid: true}}
	new_org, err := h.Queries.PartlyUpdateOrganization(h.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to update organization's value of produced plastic:\n%v", err)})
		return nil
	}
	return &new_org
}
