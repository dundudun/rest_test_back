package handlers

import (
	"context"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	Queries *sqlc.Queries
	Ctx     context.Context
	Db      *pgx.Conn
}
