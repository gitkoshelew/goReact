package authentication

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
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
}

// CreateToken ...
func CreateToken(userid uint64) (*Token, error) {
	token := &Token{}
	token.AtExpires = time.Now().Add(time.Minute * 15).Unix() // Expire time of Access token
	token.AccessUUID = uuid.NewV4().String()                  // Create a random RFC4122 version 4 UUID a cryptographically secure for Access token

	token.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() // Expire time of Refresh token
	token.RefreshUUID = uuid.NewV4().String()                   // Create a random RFC4122 version 4 UUID a cryptographically secure for Refresh token

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = token.AccessUUID
	atClaims["user_id"] = userid
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
	accessToken := r.Header.Get("Access-Token")
	log.Printf("access token from header: %s", accessToken)
	return accessToken
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
		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}
