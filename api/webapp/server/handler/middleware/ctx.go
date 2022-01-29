package middleware

import (
	"context"
	"fmt"
)

// CtxKey ...
type CtxKey int8

// CtxKeys
const (
	CtxKeyUser  CtxKey = iota
	CtxKeyEmail CtxKey = iota
)

// ContextEmail ...
func ContextEmail(ctx context.Context) (string, error) {
	email := ctx.Value(CtxKeyEmail)
	if email == "" {

		return "", fmt.Errorf("email is empty")
	}
	return email.(string), nil
}
