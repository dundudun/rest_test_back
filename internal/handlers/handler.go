package handlers

import (
	"context"

	"github.com/dundudun/rest_test_back/db/sqlc"
)

type Handler struct {
	Queries *sqlc.Queries
	Ctx     context.Context
}
