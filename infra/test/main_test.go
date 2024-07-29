package main

import (
	"crypto/tls"
	"net/http"
	"testing"
)

const baseURL = "http://YOUR_PUBLIC_IP"

func TestHTTPRedirect(t *testing.T) {
	resp, err := http.Get(baseURL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, resp.StatusCode)
	}

	location, err := resp.Location()
	if err != nil {
		t.Fatalf("Failed to get redirect location: %v", err)
	}

	expectedScheme := "https"
	if location.Scheme != expectedScheme {
		t.Errorf("Expected scheme %s, got %s", expectedScheme, location.Scheme)
	}
}

func TestHTTPSContent(t *testing.T) {
	httpsURL := "https://YOUR_PUBLIC_IP"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(httpsURL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedContentType := "text/html"
	if contentType := resp.Header.Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}
}
