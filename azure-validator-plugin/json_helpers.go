package main

import (
	"encoding/json"
)

func JsonHeader() map[string][]string {
	headerResponse := make(map[string][]string, 0)
	headerResponse["Content-Type"] = []string{"application/json"}
	return headerResponse
}

type jsonResponse struct {
	Plugin     string `json:"plugin"`
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
	Route      string `json:"route,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func PreparePayload(plugin string, statusCode int, message string, err string, route string) []byte {
	payload := &jsonResponse{
		Plugin:     plugin,
		Message:    message,
		StatusCode: statusCode,
		Error:      err,
		Route:      route,
	}
	body, _ := json.Marshal(payload)
	return body
}
