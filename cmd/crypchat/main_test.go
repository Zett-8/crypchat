package main

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	go main()

	_, err := http.Get("http://localhost:8080")

	if err != nil {
		t.Errorf("failed to start server")
	}
}
