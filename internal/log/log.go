package log

import (
	"github.com/axard/gqlgen-todo-list/internal/env"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func init() {
	var err error

	switch env.Environment {
	case "development":
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
