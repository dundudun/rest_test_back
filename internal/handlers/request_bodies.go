package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
)

func OrganizationValidation(orgStructLevel validator.StructLevel) {
	organization := orgStructLevel.Current().Interface().(OrganizationCreate)

	switch {
	case organization.ProducedPlastic.Valid && !organization.PlasticLimit.Valid:
		orgStructLevel.ReportError(organization.ProducedPlastic, "ProducedPlastic", "produced_plastic", "plasticnotvalid", "")
		orgStructLevel.ReportError(organization.PlasticLimit, "PlasticLimit", "plastic_limit", "plasticnotvalid", "")

	case organization.ProducedGlass.Valid && !organization.GlassLimit.Valid:
		orgStructLevel.ReportError(organization.ProducedGlass, "ProducedGlass", "produced_glass", "glassnotvalid", "")
		orgStructLevel.ReportError(organization.GlassLimit, "GlassLimit", "glass_limit", "glassnotvalid", "")

	case organization.ProducedBiowaste.Valid && !organization.BiowasteLimit.Valid:
		orgStructLevel.ReportError(organization.ProducedBiowaste, "ProducedBiowaste", "produced_biowaste", "biowastenotvalid", "")
		orgStructLevel.ReportError(organization.BiowasteLimit, "BiowasteLimit", "biowaste_limit", "biowastenotvalid", "")
	}
}

type OrganizationCreate struct {
	Name             pgtype.Text  `json:"name" binding:"required"`
	PlasticLimit     *pgtype.Int4 `json:"plastic_limit"`
	GlassLimit       *pgtype.Int4 `json:"glass_limit"`
	BiowasteLimit    *pgtype.Int4 `json:"biowaste_limit"`
	ProducedPlastic  *pgtype.Int4 `json:"produced_plastic"`
	ProducedGlass    *pgtype.Int4 `json:"produced_glass"`
	ProducedBiowaste *pgtype.Int4 `json:"produced_biowaste"`
}

type OrganizationPartlyUpdate struct {
	Name             *pgtype.Text `json:"name"`
	PlasticLimit     *pgtype.Int4 `json:"plastic_limit"`
	GlassLimit       *pgtype.Int4 `json:"glass_limit"`
	BiowasteLimit    *pgtype.Int4 `json:"biowaste_limit"`
	ProducedPlastic  *pgtype.Int4 `json:"produced_plastic"`
	ProducedGlass    *pgtype.Int4 `json:"produced_glass"`
	ProducedBiowaste *pgtype.Int4 `json:"produced_biowaste"`
}

type OrganizationUpdate struct {
	Name             pgtype.Text `json:"name" binding:"required"`
	PlasticLimit     pgtype.Int4 `json:"plastic_limit" binding:"required"`
	GlassLimit       pgtype.Int4 `json:"glass_limit" binding:"required"`
	BiowasteLimit    pgtype.Int4 `json:"biowaste_limit" binding:"required"`
	ProducedPlastic  pgtype.Int4 `json:"produced_plastic" binding:"required"`
	ProducedGlass    pgtype.Int4 `json:"produced_glass" binding:"required"`
	ProducedBiowaste pgtype.Int4 `json:"produced_biowaste" binding:"required"`
}

func WasteStorageValidation(wasteStorageStructLevel validator.StructLevel) {
	wasteStorage := wasteStorageStructLevel.Current().Interface().(WasteStorageCreate)

	switch {
	case wasteStorage.StoredPlastic.Valid && !wasteStorage.PlasticLimit.Valid:
		wasteStorageStructLevel.ReportError(wasteStorage.StoredPlastic, "StoredPlastic", "stored_plastic", "plasticnotvalid", "")
		wasteStorageStructLevel.ReportError(wasteStorage.PlasticLimit, "PlasticLimit", "plastic_limit", "plasticnotvalid", "")

	case wasteStorage.StoredGlass.Valid && !wasteStorage.GlassLimit.Valid:
		wasteStorageStructLevel.ReportError(wasteStorage.StoredGlass, "StoredGlass", "stored_glass", "glassnotvalid", "")
		wasteStorageStructLevel.ReportError(wasteStorage.GlassLimit, "GlassLimit", "glass_limit", "glassnotvalid", "")

	case wasteStorage.StoredBiowaste.Valid && !wasteStorage.BiowasteLimit.Valid:
		wasteStorageStructLevel.ReportError(wasteStorage.StoredBiowaste, "StoredBiowaste", "stored_biowaste", "biowastenotvalid", "")
		wasteStorageStructLevel.ReportError(wasteStorage.BiowasteLimit, "BiowasteLimit", "biowaste_limit", "biowastenotvalid", "")
	}
}

type WasteStorageCreate struct {
	Name           pgtype.Text  `json:"name" binding:"required"`
	PlasticLimit   *pgtype.Int4 `json:"plastic_limit"`
	GlassLimit     *pgtype.Int4 `json:"glass_limit"`
	BiowasteLimit  *pgtype.Int4 `json:"biowaste_limit"`
	StoredPlastic  *pgtype.Int4 `json:"stored_plastic"`
	StoredGlass    *pgtype.Int4 `json:"stored_glass"`
	StoredBiowaste *pgtype.Int4 `json:"stored_biowaste"`
}

type WasteStoragePartlyUpdate struct {
	Name           *pgtype.Text `json:"name"`
	PlasticLimit   *pgtype.Int4 `json:"plastic_limit"`
	GlassLimit     *pgtype.Int4 `json:"glass_limit"`
	BiowasteLimit  *pgtype.Int4 `json:"biowaste_limit"`
	StoredPlastic  *pgtype.Int4 `json:"stored_plastic"`
	StoredGlass    *pgtype.Int4 `json:"stored_glass"`
	StoredBiowaste *pgtype.Int4 `json:"stored_biowaste"`
}

type WasteStorageUpdate struct {
	Name           pgtype.Text `json:"name" binding:"required"`
	PlasticLimit   pgtype.Int4 `json:"plastic_limit" binding:"required"`
	GlassLimit     pgtype.Int4 `json:"glass_limit" binding:"required"`
	BiowasteLimit  pgtype.Int4 `json:"biowaste_limit" binding:"required"`
	StoredPlastic  pgtype.Int4 `json:"stored_plastic" binding:"required"`
	StoredGlass    pgtype.Int4 `json:"stored_glass" binding:"required"`
	StoredBiowaste pgtype.Int4 `json:"stored_biowaste" binding:"required"`
}
