package authentication_test

import (
	"fmt"
	"goReact/webapp/server/handler/authentication"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokens_CreateCustomToken(t *testing.T) {
	payload := make(map[string]interface{})
	payload["user_id"] = 1
	expireTime := time.Duration(20)
	secretKey := "secret"

	t.Run("Valid", func(t *testing.T) {
		payload["user_id"] = 1
		expireTime = time.Duration(20)
		secretKey = "secret"
		tk, err := authentication.CreateCustomToken(payload, expireTime, secretKey)
		assert.NoError(t, err)
		assert.NotNil(t, tk)
	})

	t.Run("Invalid expire time", func(t *testing.T) {
		expireTime = time.Duration(-20)
		secretKey = "secret"
		_, err := authentication.CreateCustomToken(payload, expireTime, secretKey)
		assert.Error(t, err)
	})

	t.Run("Invalid secret key", func(t *testing.T) {
		expireTime = time.Duration(20)
		secretKey = "    "
		_, err := authentication.CreateCustomToken(payload, expireTime, secretKey)
		assert.Error(t, err)
	})

	t.Run("Empty payload", func(t *testing.T) {
		for k := range payload {
			delete(payload, k)
		}
		expireTime = time.Duration(20)
		secretKey = "secret"
		_, err := authentication.CreateCustomToken(payload, expireTime, secretKey)
		assert.Error(t, err)
	})
}

func TestToken_ParseCustomToken(t *testing.T) {
	payload := make(map[string]interface{})
	payload["user_id"] = 1
	expireTime := time.Duration(20)
	secretKey := "secret"
	token := authentication.TestToken(payload, expireTime, secretKey)

	t.Run("Valid", func(t *testing.T) {
		jwtmc, err := authentication.ParseCustomToken(token, secretKey)
		assert.NoError(t, err)
		assert.NotNil(t, jwtmc)
	})

	t.Run("Invalid token string", func(t *testing.T) {
		jwtmc, err := authentication.ParseCustomToken(fmt.Sprintf("asdasdas%sasdasd", token), secretKey)
		assert.Error(t, err)
		assert.Nil(t, jwtmc)
	})

	t.Run("Invalid secret key", func(t *testing.T) {
		jwtmc, err := authentication.ParseCustomToken(token, "invalid secret key")
		assert.Error(t, err)
		assert.Nil(t, jwtmc)
	})

	t.Run("Token is expired", func(t *testing.T) {
		expiredToken := authentication.TestToken(payload, -20, secretKey)

		jwtmc, err := authentication.ParseCustomToken(expiredToken, secretKey)
		assert.Error(t, err)
		assert.Nil(t, jwtmc)
	})

}
