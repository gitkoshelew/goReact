package jwt

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"user/pkg/logging"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

var (
	// ErrUnexpectedSigningMethod ...
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	// ErrBadRequest ...
	ErrBadRequest = errors.New("bad request")
)

var (
	// EmailSecretKey ...
	EmailSecretKey = os.Getenv("EMAIL_CONFIRM_SECRET")
	logger         = logging.GetLogger()
)

// Token ...
type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUUID   string `json:"accessUuid"`
	RefreshUUID  string `json:"refreshUuid"`
	AtExpires    int64  `json:"atExpires"`
	RtExpires    int64  `json:"rtExpires"`
}

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string `json:"accessUuid"`
	UserID     uint64 `json:"userId"`
	Role       string `json:"role"`
}

// CreateToken ...
func CreateToken(userid uint64, role string) (*Token, error) {
	token := &Token{}
	token.AtExpires = time.Now().Add(time.Minute * 120).Unix() // Expire time of Access token
	token.AccessUUID = uuid.NewV4().String()                   // Create a random RFC4122 version 4 UUID a cryptographically secure for Access token

	token.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() // Expire time of Refresh token
	token.RefreshUUID = uuid.NewV4().String()                   // Create a random RFC4122 version 4 UUID a cryptographically secure for Refresh token

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = token.AccessUUID
	atClaims["user_id"] = userid
	atClaims["role"] = role
	atClaims["exp"] = token.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = token.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["role"] = role
	rtClaims["exp"] = token.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractRefreshToken ...
func ExtractRefreshToken(r *http.Request) string {
	JWTcookie, err := r.Cookie("Refresh-Token")
	if err != nil {
		log.Print("Error occured while reading cookie")
	}
	return JWTcookie.Value
}

// ExtractAccessToken ...
func ExtractAccessToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken ...
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractAccessToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// IsValid ...
func IsValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		role := fmt.Sprint(claims["role"])
		if err != nil {
			return nil, err
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
			Role:       role,
		}, nil
	}
	return nil, err
}

// CreateCustomToken ...
func CreateCustomToken(payload map[string]interface{}, expireTime time.Duration, secretKey string) (string, error) {
	if expireTime < 1 {
		logger.Debugf("Bad request expire time: %v. Err msg: %v", expireTime, ErrBadRequest)
		return "", fmt.Errorf("%v, expire time = %v", ErrBadRequest, expireTime)
	}

	checkSecret := strings.ReplaceAll(secretKey, " ", "")
	if checkSecret == "" {
		logger.Debugf("Bad request secret key: %v. Err msg: %v", secretKey, ErrBadRequest)
		return "", fmt.Errorf("%v, secret key: %v", ErrBadRequest, secretKey)
	}

	if len(payload) < 1 {
		logger.Debugf("Bad request empty payload: %v. Err msg: %v", payload, ErrBadRequest)
		return "", fmt.Errorf("%v, empty payload: %v", ErrBadRequest, payload)
	}

	claims := jwt.MapClaims{}
	for key, element := range payload {
		claims[key] = element
	}
	claims["exp"] = time.Now().Add(time.Minute * expireTime).Unix()

	logger.Debugf("JWT.MapClaims created: %v", claims)

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tk.SignedString([]byte(secretKey))

	if err != nil {
		logger.Debugf("Error during complete signed token creating. Err msg: %v", err)
		return "", err
	}
	return token, nil
}

// ParseCustomToken ...
func ParseCustomToken(tokenStr, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Debugf("Unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("%v. %v", ErrUnexpectedSigningMethod, token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
	if err != nil {
		logger.Debugf("Error during token parsing. Err msg: %v", err)
		return nil, err
	}

	_, ok := token.Claims.(jwt.Claims)
	if !ok && !token.Valid {
		logger.Debugf("Error during jwt.Claims creating. Err msg: %v", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		logger.Debugf("Error during jwt.MapClaims creating. Err msg: %v", err)
		return nil, err
	}

	return claims, nil
}
