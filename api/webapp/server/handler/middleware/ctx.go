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

func ContextEmail(ctx context.Context) string {
	fmt.Println("HELLO FROM CONTEXT EMAIL start ")
	email := ctx.Value(CtxKeyEmail)
	if email == "" {

		return " email is 0"
	}

	fmt.Println("HELLO FROM CONTEXT EMAIL ", email)
	return email.(string)
}
