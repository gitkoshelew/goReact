package service

import (
	"encoding/json"
	"fmt"
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

	//var result map[string]interface{}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//json.NewDecoder(resp.Body).Decode(&result)
	bodyBytesJson, err := json.Marshal(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("bodyBytes -   :", bodyBytes)

	fmt.Println("bodyBytesJson -   :", bodyBytesJson)


	return &bodyBytes, nil
}

//GetToken
func GetToken(getTokenURI string) (*oauth2.Token, error) {

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
