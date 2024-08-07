package api

import "context"

type userCtxKeyType string

var userCtxKey = userCtxKeyType("user_ctx_key")

type ApiUser struct {
	UserID int32
}

func GetContextUser(ctx context.Context) (*ApiUser, bool) {
	v := ctx.Value(userCtxKey)
	if v == nil {
		return nil, false
	}

	user, ok := v.(*ApiUser)
	return user, ok
}

func SetContextUser(ctx context.Context, u *ApiUser) context.Context {
	return context.WithValue(ctx, userCtxKey, u)
}
