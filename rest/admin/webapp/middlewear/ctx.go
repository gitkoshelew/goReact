package middlewear

import (
	"context"
	"fmt"
)

// CtxKey ...
type CtxKey int8

// CtxKeys
const (
	CtxKeyFile CtxKey = iota
)

// CtxFile ...
func CtxFile(ctx context.Context) (string, error) {
	path := ctx.Value(CtxKeyFile)
	if path == "" {

		return "", fmt.Errorf("path is empty")
	}
	return path.(string), nil
}
