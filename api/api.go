package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"url-shortener-app/config"
	"url-shortener-app/entities"
)

func ShortenURL(inputURL string) (string, error) {
	var response entities.ApiResponse

	requestURL := config.Config.GetString("api.url")
	params := url.Values{}
	params.Set("url", inputURL)
	reqBody := bytes.NewBufferString(params.Encode())
	resp, err := http.Post(requestURL, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", errors.New(response.Error)
	}

	return response.Result_url, nil
}
