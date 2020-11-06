package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

const (
	development = "develop"
	production  = "product"
)

var (
	Logger *zap.Logger
)

func isValidLevel(s string) bool {
	return s == development ||
		s == production
}

func level() string {
	if v := os.Getenv("LOG_LVL"); v != "" {
		if !isValidLevel(v) {
			panic(fmt.Errorf("invalid value of $ENV: %s", v))
		}

		return v
	}

	return production
}

func init() {
	var err error

	switch level() {
	case development:
		Logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

	default:
		Logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
}
