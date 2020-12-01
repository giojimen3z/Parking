package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProductiveScopes(t *testing.T) {
	os.Setenv(scope, writeScope)
	defer os.Clearenv()
	want := true
	got := isInProductiveScopes()
	assert.Equal(t, want, got)
}

func TestIsNotProductiveScopes(t *testing.T) {
	os.Setenv(scope, "test")
	defer os.Clearenv()
	want := false
	got := isInProductiveScopes()
	assert.Equal(t, want, got)
}
