package service

import (
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
)

//GetUserDataGit
func UserDataGit( getDataURI string, headerValue string) (map[string]interface{}, error) {

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
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	defer resp.Body.Close()
	return result, nil
}

//GetUserDataGit
func GetTokenGit(getTokenURI string) (*oauth2.Token, error) {

	req, err := http.NewRequest(http.MethodGet, getTokenURI, nil)
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
