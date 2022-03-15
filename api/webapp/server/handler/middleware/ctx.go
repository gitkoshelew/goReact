package middleware

import (
	"context"
	"errors"
)

// CtxKey ...
type CtxKey int8

// CtxKeys
const (
	CtxKeyUser  CtxKey = 1
	CtxKeyEmail CtxKey = 2
)

var (
	// ErrEmptyEmail ...
	ErrEmptyEmail = errors.New("email is empty")
)

// ContextEmail ...
func ContextEmail(ctx context.Context) (string, error) {
	email := ctx.Value(CtxKeyEmail).(string)
	if email == "" {
		return "", ErrEmptyEmail
	}
	return email, nil
}
