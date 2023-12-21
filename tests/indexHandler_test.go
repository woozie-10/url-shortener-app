package tests

import (
	"net/http"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	url := "http://localhost:2311"
	resp, err := http.Get(url)

	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received status code %d, expected 200", resp.StatusCode)
	}
	t.Log("Test completed successfully")
}
