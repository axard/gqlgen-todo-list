package api

import (
	"context"
	"fmt"
	"strconv"
)

type contextKey string

const (
	UserID = contextKey("UserID")
)

// UserIDFromContext извлекает из контекста ctx пользовательский ID в строковом виде, переводит его в
// число, которое и возвращает. В случае ошибке на этапе извлечения или конвертации возвращается
// ошибка
func UserIDFromContext(ctx context.Context) (int, error) {
	str, ok := ctx.Value(UserID).(string)
	if !ok {
		return 0, fmt.Errorf("invalid type of UserID: \"%T\"", ctx.Value(UserID))
	}

	userID, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
