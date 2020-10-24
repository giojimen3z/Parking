package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWriteConnection(t *testing.T) {
	con, _ := GetWriteConnection()
	assert.NotNil(t, con)
}

func TestGetReadConnection(t *testing.T) {
	con, _ := GetReadConnection()
	assert.NotNil(t, con)
}
