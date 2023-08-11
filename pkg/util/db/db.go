package dbutil

import "database/sql"

func ToSqlNullString(n string) sql.NullString {
	return sql.NullString{
		String: n,
		Valid:  true,
	}
}

func ToSqlNullInt32(n int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: n,
		Valid: true,
	}
}
