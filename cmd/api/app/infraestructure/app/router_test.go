package app

import (
	"os"
	"testing"
)

func TestStartApplication(t *testing.T) {

	os.Setenv("PORT", "0")
	StartApp()
}
