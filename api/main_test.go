package main

import (
	"bytes"
	"encoding/json"
	"github.com/csci4950tgt/api/models"
	"net/http"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	_, err := http.Get("http://127.0.0.1:8080/api")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTicket(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	url := "http://127.0.0.1:8080/api/tickets"

	// Test that GET method works
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		t.Error("HTTP error, reason " + resp.Status)
	}

	// Test that POST method works
	request := &models.Ticket{URL: "https://www.google.com", ScreenShot: []models.ScreenShot{models.ScreenShot{Width: 1920, Height: 1080, Filename: "screenshot.png"}}}
	requestJson, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	resp, err = http.Post(url, "application/json", bytes.NewBuffer(requestJson))

	statusCode = resp.StatusCode
	if statusCode != http.StatusOK {
		t.Error("HTTP error, reason " + resp.Status)
	}

	// Test invalid method - DELETE
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err = client.Do(req)
	statusCode = resp.StatusCode
	if statusCode != http.StatusMethodNotAllowed {
		t.Error("Expected http status code: 405, actual code: " + resp.Status)
	}
}

func TestGetTicket(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://127.0.0.1:8080/tickets/2000000")
	if err != nil {
		t.Error(err)
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusNotFound {
		t.Error("Expected http status code: 404, actual code: " + resp.Status)
	}
}
