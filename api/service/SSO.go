package service

import (
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

//GetUserData
func GetUserData(getDataURI string, headerValue string) (*[]byte, error) {

	req, err := http.NewRequest(http.MethodGet, getDataURI, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", headerValue)
	var c http.Client
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &bodyBytes, nil
}

//GetToken Exchange  code for an access token
func GetToken(getTokenURI string) (*oauth2.Token, error) {

	req, err := http.NewRequest(http.MethodPost, getTokenURI, nil)
	if err != nil {

		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	var c http.Client
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	token := &oauth2.Token{}

	if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}
