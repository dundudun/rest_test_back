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
