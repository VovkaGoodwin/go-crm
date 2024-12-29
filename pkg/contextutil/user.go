package contextutil

import (
	"context"
	"errors"
)

func SetCurrentUserID(ctx context.Context, id int32) context.Context {
	return context.WithValue(ctx, "userID", id)
}

func GetCurrentUserID(ctx context.Context) (int32, error) {
	if id, ok := ctx.Value("userID").(int32); ok {
		return id, nil
	}

	return 0, errors.New("user id does not exist")
}
