package utils

import "github.com/jackc/pgx/v5/pgtype"

func ToPgText(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}
