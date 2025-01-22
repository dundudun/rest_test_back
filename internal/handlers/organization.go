package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/dundudun/rest_test_back/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *Handler) CreateOrganization(c *gin.Context) {
	//TOCHECK: dynamic body and number of fields
	//TOCHECK: what values will be in *pgtype.Int4 (*pgtype.Int4.Int32 and *pgtype.Int4.Valid)
	//TOCHECK: if name is required as expected and other fields are optional as expected (should be in test also)
	var requestBody OrganizationCreate
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.CreateOrganizationParams{
		Name:             requestBody.Name,
		PlasticLimit:     utils.OptionalInt4(requestBody.PlasticLimit),
		GlassLimit:       utils.OptionalInt4(requestBody.GlassLimit),
		BiowasteLimit:    utils.OptionalInt4(requestBody.BiowasteLimit),
		ProducedPlastic:  utils.OptionalInt4(requestBody.ProducedPlastic),
		ProducedGlass:    utils.OptionalInt4(requestBody.ProducedGlass),
		ProducedBiowaste: utils.OptionalInt4(requestBody.ProducedBiowaste),
	}
	organization, err := handler.Queries.CreateOrganization(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create organization:\n%v", err)})
		return
	}

	c.JSON(http.StatusOK, organization)
}

func (handler *Handler) GetOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}
	organization, err := handler.Queries.GetOrganization(handler.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get organization:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, organization)
}

func (handler *Handler) ListOrganizations(c *gin.Context) {
	organizations, err := handler.Queries.ListOrganizations(handler.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get organizations:\n%v", err)})
		return
	}
	c.JSON(http.StatusOK, organizations)
}

func (handler *Handler) PartlyChangeOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	var requestBody OrganizationPartlyUpdate
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	//TOCHECK: that i can't set values to null accidently
	params := sqlc.PartlyUpdateOrganizationParams{
		ID:               id,
		Name:             utils.OptionalText(requestBody.Name),
		PlasticLimit:     utils.OptionalInt4(requestBody.PlasticLimit),
		GlassLimit:       utils.OptionalInt4(requestBody.GlassLimit),
		BiowasteLimit:    utils.OptionalInt4(requestBody.BiowasteLimit),
		ProducedPlastic:  utils.OptionalInt4(requestBody.ProducedPlastic),
		ProducedGlass:    utils.OptionalInt4(requestBody.ProducedGlass),
		ProducedBiowaste: utils.OptionalInt4(requestBody.ProducedBiowaste),
	}

	// reqVal := reflect.ValueOf(req)
	// reqType := reflect.TypeOf(req)
	// paramVal := reflect.ValueOf(&params).Elem()

	// for i := 0; i < reqVal.NumField(); i++ {
	// 	fieldName := reqType.Field(i)
	// 	reqField := reqVal.Field(i)
	// 	paramsField := paramVal.FieldByName(fieldName.Name)

	// 	for reqField.IsValid() && !reqField.IsNil() {
	// 		paramsField.Set(reqField.Elem())
	// 	}
	// }

	organization, err := handler.Queries.PartlyUpdateOrganization(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update organization"})
		return
	}

	c.JSON(http.StatusOK, organization)
}

func (handler *Handler) ChangeOrganization(c *gin.Context) {
	//TOCHECK: if i can send whole body but still have optional field - so that would be as null values, would they parse and work as expected
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	var requestBody OrganizationUpdate
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body:\n%v", err)})
		return
	}

	params := sqlc.UpdateOrganizationParams{
		ID:               id,
		Name:             requestBody.Name,
		PlasticLimit:     requestBody.PlasticLimit,
		GlassLimit:       requestBody.GlassLimit,
		BiowasteLimit:    requestBody.BiowasteLimit,
		ProducedPlastic:  requestBody.ProducedPlastic,
		ProducedGlass:    requestBody.ProducedGlass,
		ProducedBiowaste: requestBody.ProducedBiowaste,
	}
	organization, err := handler.Queries.UpdateOrganization(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update organization:\n%v", err)})
		return
	}

	c.JSON(http.StatusOK, organization)
}

func (handler *Handler) DeleteOrganization(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}

	//TOCHECK: if pgx.ErrNoRows works correctly
	if _, err = handler.Queries.DeleteOrganization(handler.Ctx, id); err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("no organization with such id: %d", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete organization:\n%v", err)})
		return
	}
	c.Status(http.StatusOK)
}

func (handler *Handler) ProduceWaste(c *gin.Context) {
	tx, err := handler.Db.Begin(handler.Ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to start transaction:\n%v", err)})
		return
	}
	defer tx.Rollback(handler.Ctx)
	qtx := handler.Queries.WithTx(tx)

	overproduce := int32(0)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid organization ID:\n%v", err)})
		return
	}
	organization, err := qtx.GetOrganization(handler.Ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to get organization from db:\n%v", err)})
		return
	}

	wasteAmountToAddInt64, err := strconv.ParseInt(c.Query("amount"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failure of getting 'amount' query parameter:\n%v", err)})
		return
	}
	wasteAmountToAdd := int32(wasteAmountToAddInt64)
	if wasteAmountToAdd <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "waste amount to add shouldn't be less or equal to 0"})
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
		query next level (1st level from Organizations is outside loop)
		add wasteStorages from new level to all wasteStorages (for ones from wasteStorage accumulate Distance with parent)
		find across all wasteStorages that we haven't worked with one that have min distance
		mark it somehow or pop it?
		work with one that have min distance
			if have free space in storage subract from amount
			if no free space just do nothing, go further
		check amount on zero
			if greater than zero repeat for next level
			in other case exit loop
	*/

	//TODO: try slice first, then see if map[level]fetchedWasteStorages would be better
	var wasteStorages []fetchedWasteStorages
	var updatedWasteStorages []sqlc.WasteStorage

	level := 1
	// 1st level (one next to Organization)
	fmt.Printf("going to fetch from Organization: %d\n", id)
	fetched, err := qtx.FromOrgPlasticStors(handler.Ctx, pgtype.Int8{Int64: id, Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to get waste storages for organization:\n%v", err)})
		return
	}
	if len(fetched) == 0 {
		//TOCHECK: query to update Organization produced_wastetype value if there isn't WasteStorages for it
		//TOCHECK: overproduce
		org := handler.addWasteToOrg(c, &organization, wasteAmountToAdd)
		if org == nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":                        "no wasteStorages for organization",
			"organization":                   org,
			"changed_waste_storages":         updatedWasteStorages,
			"waste_above_organization_limit": overproduce,
		})
		return
	}
	for _, storage := range fetched {
		wasteStorages = append(wasteStorages, fetchedWasteStorages{
			WorkedWith:     false,
			level:          level,
			ID:             storage.ID,
			Name:           storage.Name,
			WasteLimit:     storage.PlasticLimit,
			StoredWaste:    storage.StoredPlastic,
			DistanceMeters: storage.DistanceMeters,
		})
	}
	idClosest := findClosestWasteStorage(wasteStorages)
	wasteStorages[idClosest].WorkedWith = true

	freeSpaceInStore := wasteStorages[idClosest].WasteLimit.Int32 - wasteStorages[idClosest].StoredWaste.Int32
	if freeSpaceInStore > 0 {
		var newStoredWaste int32
		if freeSpaceInStore >= wasteAmountToAdd {
			wasteAmountToAdd = 0
			newStoredWaste = wasteStorages[idClosest].StoredWaste.Int32 + wasteAmountToAdd
		} else {
			wasteAmountToAdd -= freeSpaceInStore
			newStoredWaste = wasteStorages[idClosest].WasteLimit.Int32
		}

		params := sqlc.PartlyUpdateWasteStorageParams{
			ID:            wasteStorages[idClosest].ID,
			StoredPlastic: pgtype.Int4{Int32: newStoredWaste, Valid: true},
		}
		updatedWasteStorage, err := qtx.PartlyUpdateWasteStorage(handler.Ctx, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage's value of stored_plastic:\n%v", err)})
			return
		}
		updatedWasteStorages = append(updatedWasteStorages, updatedWasteStorage)
	}

	// all next levels (ones after WasteStorages)
	nextLevel := true
	for wasteAmountToAdd > 0 {
		//TOCHECK: to not get in this loop if waste was distributed on 1st level
		//TOCHECK: exiting right when produced waste is distributed and no more db calls or loops
		//TOCHECK: to correctly get new closest wasteStorage each time and skip ones that we looked up already
		//TOCHECK: if wasteStorages are full but amount of waste still left
		//TOCHECK: do not call db if there isn't more connections
		//TOCHECK: to correctly update/distribute values across wasteStorages
		//TODO: remore this if statement after testing
		if !nextLevel {
			fmt.Println("no WasteStorages have connections to another ones")
		} else if nextLevel {
			level++
			counter := 0
			// fetch next level from previous one
			for _, wasteStorage := range wasteStorages {
				if wasteStorage.level-1 != level {
					counter++
					continue
				}
				fmt.Printf("going to fetch from WasteStorage: %d, %s\n", wasteStorage.ID, wasteStorage.Name.String)
				fetched, err := qtx.FromStorsPlasticStors(handler.Ctx, pgtype.Int8{Int64: wasteStorage.ID, Valid: true})
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
						level:          level,
						ID:             stor.ID,
						Name:           stor.Name,
						WasteLimit:     stor.PlasticLimit,
						StoredWaste:    stor.StoredPlastic,
						DistanceMeters: stor.DistanceMeters + wasteStorage.DistanceMeters,
					})
				}
			}

			if counter == len(wasteStorages) {
				nextLevel = false
			}
		}

		idClosest = findClosestWasteStorage(wasteStorages)
		if idClosest == -1 {
			//TOCHECK: update Organization produced_wastetype value and list of WasteStorages
			org := handler.addWasteToOrg(c, &organization, wasteAmountToAdd)
			if org == nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":                        "no more wasteStorages available to receive waste",
				"organization":                   org,
				"changed_waste_storages":         updatedWasteStorages,
				"waste_above_organization_limit": overproduce,
			})
			return
		}

		wasteStorages[idClosest].WorkedWith = true

		freeSpaceInStorage := wasteStorages[idClosest].WasteLimit.Int32 - wasteStorages[idClosest].StoredWaste.Int32
		if freeSpaceInStorage > 0 {
			var newStoredWaste int32
			if freeSpaceInStorage >= wasteAmountToAdd {
				wasteAmountToAdd = 0
				newStoredWaste = wasteStorages[idClosest].StoredWaste.Int32 + wasteAmountToAdd
			} else {
				wasteAmountToAdd -= freeSpaceInStorage
				newStoredWaste = wasteStorages[idClosest].WasteLimit.Int32
			}

			params := sqlc.PartlyUpdateWasteStorageParams{
				ID:            wasteStorages[idClosest].ID,
				StoredPlastic: pgtype.Int4{Int32: newStoredWaste, Valid: true},
			}
			updatedWasteStorage, err := qtx.PartlyUpdateWasteStorage(handler.Ctx, params)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to update waste storage's value of stored_plastic:\n%v", err)})
				return
			}
			updatedWasteStorages = append(updatedWasteStorages, updatedWasteStorage)
		}
	}

	if err := tx.Commit(handler.Ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("fail to Commit transaction:\n%v", err)})
		return
	}
	//TOCHECK: final output
	c.JSON(http.StatusOK, gin.H{
		"message":                        "waste was fully distributed across waste storages",
		"organization":                   organization,
		"updated_waste_storages":         updatedWasteStorages,
		"waste_above_organization_limit": overproduce,
	})
}

type fetchedWasteStorages struct {
	WorkedWith     bool
	level          int
	ID             int64
	Name           pgtype.Text
	WasteLimit     pgtype.Int4
	StoredWaste    pgtype.Int4
	DistanceMeters int32
}

func findClosestWasteStorage(wasteStorages []fetchedWasteStorages) int {
	StorageIdWithMinDistance := -1

	for i, wasteStorage := range wasteStorages {
		if wasteStorage.WorkedWith {
			continue
		}

		if StorageIdWithMinDistance == -1 {
			StorageIdWithMinDistance = i
		} else if wasteStorage.DistanceMeters < wasteStorages[StorageIdWithMinDistance].DistanceMeters {
			StorageIdWithMinDistance = i
		}
	}

	return StorageIdWithMinDistance
}

func (handler *Handler) addWasteToOrg(c *gin.Context, organization *sqlc.Organization, wasteAmountToAdd int32) *sqlc.Organization {
	//TODO: flexibily on wastetype to add (and add all three wastetypes)
	overproduce := (organization.ProducedPlastic.Int32 + wasteAmountToAdd) % organization.PlasticLimit.Int32
	newValue := (organization.ProducedPlastic.Int32 + wasteAmountToAdd) - overproduce

	params := sqlc.PartlyUpdateOrganizationParams{ID: organization.ID, ProducedPlastic: pgtype.Int4{Int32: newValue, Valid: true}}
	newOrganization, err := handler.Queries.PartlyUpdateOrganization(handler.Ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failure to update organization's value of produced plastic:\n%v", err)})
		return nil
	}
	return &newOrganization
}
