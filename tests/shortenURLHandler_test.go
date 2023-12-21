package tests

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestShortenURLHandler(t *testing.T) {
	hostURL := "http://localhost:2311"
	params := url.Values{}
	params.Set("url", "https://www.google.com/")
	reqBody := bytes.NewBufferString(params.Encode())
	resp, err := http.Post(hostURL, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Received status code %d, expected 200", resp.StatusCode)
	}

	t.Log("Test completed successfully")
}
