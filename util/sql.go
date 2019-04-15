package util

import "database/sql"

// StringOrNull ..
func StringOrNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: str, Valid: true}
}
