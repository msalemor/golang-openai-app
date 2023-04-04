package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AzureOpenAIRequest(text string, app App) (string, error) {
	// Define the request body
	requestBody := RequestBody{
		Prompt:      text,
		MaxTokens:   300,
		N:           1,
		Temperature: 0.7,
	}

	// Marshal the request body into a JSON string
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request with the request body
	req, err := http.NewRequest("POST", app.Endpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}

	// Set the API key as an authorization header
	req.Header.Set("api-key", app.APIKey)

	// Set the content type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var textCompletion TextCompletion
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &textCompletion)

	return textCompletion.Choices[0].Text, nil
}
