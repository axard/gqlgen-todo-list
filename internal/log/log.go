package log

import (
	"fmt"

	"github.com/axard/gqlgen-todo-list/internal/cfg"
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

func init() {
	var err error

	l := cfg.LogLevel()
	if !isValidLevel(l) {
		panic(fmt.Sprintf("unknown log level: %s", l))
	}

	switch l {
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
