package utils

import "github.com/jackc/pgx/v5/pgtype"

func OptionalInt4(ptr *pgtype.Int4) pgtype.Int4 {
	if ptr != nil {
		return *ptr
	}
	return pgtype.Int4{Valid: false}
}

func OptionalText(ptr *pgtype.Text) pgtype.Text {
	if ptr != nil {
		return *ptr
	}
	return pgtype.Text{Valid: false}
}
