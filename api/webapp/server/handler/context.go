package handler

import (
	"context"
	"errors"
)

// CtxKey ...
type CtxKey int8

// CtxKeys
const (
	CtxKeyUser                         CtxKey = 1
	CtxKeyUserValidation               CtxKey = 2
	CtxKeyEmail                        CtxKey = 3
	CtxKeyLoginValidation              CtxKey = 4
	CtxKeyBookingValidation            CtxKey = 5
	CtxKeyFreeSeatsSearchReqValidation CtxKey = 6
	CTXKeyAccessTokenGitOAuth          CtxKey = 7
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
