package config

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapToDatabaseTimeStamp(t *testing.T) {
	timeStamp := time.Now()
	nilTimeStamp := time.Time{}
	sqlTimeStamp := NewDatabaseTimeStamp(timeStamp)
	sqlNullTimeStamp := NewDatabaseTimeStamp(nilTimeStamp)

	assert.IsType(t, sql.NullTime{}, sqlNullTimeStamp)
	assert.IsType(t, time.Time{}, sqlTimeStamp)
}

func TestMapToDatabaseString(t *testing.T) {
	timeStamp := "garex"
	nilTimeStamp := ""
	sqlTimeStamp := NewDatabaseString(timeStamp)
	sqlNullTimeStamp := NewDatabaseString(nilTimeStamp)

	assert.IsType(t, sql.NullString{}, sqlNullTimeStamp)
	assert.IsType(t, "string", sqlTimeStamp)
}
