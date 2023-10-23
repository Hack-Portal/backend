package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgText(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func ToPgDate(value time.Time) pgtype.Date {
	return pgtype.Date{
		Time:  value,
		Valid: true,
	}
}
