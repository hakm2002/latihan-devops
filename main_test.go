package main

import (
	"io"
	"net/http"
	"testing"
)

// func TestMain(m *testing.M) {
// 	go func() {
// 		if err := http.ListenAndServe(":8081", nil); err != nil {
// 			log.Fatalf("Server failed to start: %v", err)
// 		}
// 	}()
// 	m.Run()
// }

func TestRootEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from GitLab CI/CD!!"
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyString := string(bodyBytes)
	if bodyString != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", bodyString, expected)
	}
}

func TestTestEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from test endpoint"
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyString := string(bodyBytes)
	if bodyString != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", bodyString, expected)
	}
}
