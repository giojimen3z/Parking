package config

import (
	"os"
	"strings"
)

const (
	appId                  = ""
	goEnvironment          = "GO_ENVIRONMENT"
	production             = "production"
	readScope              = ""
	scope                  = "SCOPE"
	dbName                 = "bikeparking"
	dbHost                 = "localhost:3306"
	dbTestHost             = "localhost:3306"
	dbTestName             = "bikeparking"
	readDBAdminUser        = "root"
	dbAdminPwd             = "root"
	readDBTestUser         = "root"
	readDBTestPwd          = "root"
	writeDBAdminUser       = "root"
	writeDBAdminPwd        = "root"
	writeDBTestUser        = "root"
	writeDBTestPwd         = "root"
	writeScope             = "write"
	localScope             = "local"

)

func GetAppId() string {
	return appId
}

func IsProductiveScope() bool {
	return isProduction() && isInProductiveScopes()
}
func isProduction() bool {
	return strings.EqualFold(os.Getenv(goEnvironment), production)
}

func isInProductiveScopes() bool {
	var productiveScopes = []string{writeScope, readScope}

	actualScope := getActualScope()

	for _, productiveScope := range productiveScopes {
		if strings.EqualFold(actualScope, productiveScope) {
			return true
		}
	}

	return false
}
func getActualScope() string {
	return os.Getenv(scope)
}


// IsLocalScope return true if environment is locally, false otherwise
func IsLocalScope() bool {
	return strings.EqualFold(getActualScope(), localScope)
}



func GetRoutePrefix() string {
	if !IsProductiveScope() {
		return "/" + getActualScope()
	}

	return ""
}
