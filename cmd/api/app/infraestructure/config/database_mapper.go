package config

import (
	"database/sql"
	"time"
)

const emptyStringLength = 0

func NewDatabaseTimeStamp(value time.Time) interface{} {
	if value.Equal(time.Time{}) {
		return sql.NullTime{}
	}
	return value
}

func NewDatabaseString(value string) interface{} {
	if len(value) == emptyStringLength {
		return sql.NullString{}
	}
	return value
}
