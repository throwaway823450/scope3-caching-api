package measurement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var _ Client = &clientImpl{}

type Client interface {
	Measure(BatchRequest) (*Response, error)
}

func NewClient(endpoint string, apiKey string) Client {
	return &clientImpl{endpoint: endpoint, apiKey: apiKey}
}

type clientImpl struct {
	endpoint string
	apiKey   string
}

func (c *clientImpl) Measure(batchRequest BatchRequest) (*Response, error) {
	body, err := json.Marshal(batchRequest)
	if err != nil {
		return nil, fmt.Errorf("Error serializing request data: %v\n", err)
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Error createing new request: %v\n", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error calling endpoint: %v\n", err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v\n", err)
	}

	// Check the HTTP status code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK HTTP status: %s\nResponse body: %s\n", response.Status, string(responseBody))
	}

	// Deserialize the response into the Response struct
	var responseData Response
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, fmt.Errorf("Error deserializing response JSON: %v\n", err)
	}

	return &responseData, nil
}
