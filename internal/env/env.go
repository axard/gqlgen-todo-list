package env

import (
	"fmt"
	"os"
)

const (
	development = "develop"
	production  = "product"
)

var Environment = development

func isValid(s string) bool {
	return s == development ||
		s == production
}

func init() {
	if v := os.Getenv("ENV"); v != "" {
		if !isValid(v) {
			panic(fmt.Errorf("invalid value of $ENV: %s", v))
		}

		Environment = v
	}
}
