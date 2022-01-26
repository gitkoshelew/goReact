package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TestToken ...
func TestToken(payload map[string]string, expireTime time.Duration, secretKey string) string {

	claims := jwt.MapClaims{}
	for key, element := range payload {
		claims[key] = element
	}
	claims["exp"] = time.Now().Add(time.Minute * expireTime).Unix()

	logger.Debugf("JWT.MapClaims created: %v", claims)

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tk.SignedString([]byte(secretKey))

	if err != nil {
		return ""
	}
	return token
}
