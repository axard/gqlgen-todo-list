package cfg

import (
	"fmt"
	"os"
)

func getenvOrPanic(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("variable %s should be set", key))
	}

	return v
}

// PgURL return value of environment variable PG_URL
func PgURL() string {
	return getenvOrPanic("PG_URL")
}

// LogLevel return value of environment variable LOG_LEVEL
func LogLevel() string {
	return getenvOrPanic("LOG_LEVEL")
}
