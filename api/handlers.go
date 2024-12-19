package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"scope3/caching-api/measurement"
)

// TODO: make these configurable
const (
	defaultCountry     = "US"
	defaultChannel     = "web"
	defaultImpressions = 1000
)

// TODO: add interface to make testing easier
type Handler struct {
	measurementClient measurement.CachingClient
}

func NewHandler(measurementClient measurement.CachingClient) Handler {
	return Handler{measurementClient: measurementClient}
}

func (h Handler) PostEmmisions(w http.ResponseWriter, r *http.Request) {
	// Convert Rerquest
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	// TODO: handle bad requests
	var batchRequest BatchRequest
	err = json.Unmarshal(reqBody, &batchRequest)
	if err != nil {
		fmt.Println(err)
	}

	measurementRequest := measurement.BatchCachingRequest{}
	for _, row := range batchRequest.Rows {
		measurementRequest.Rows = append(measurementRequest.Rows, measurement.CachingRequest{
			Request: measurement.Request{
				InventoryId: row.InventoryId,
				Country:     defaultCountry,
				Channel:     defaultChannel,
				Impressions: defaultImpressions,
				UtcDatetime: "2024-10-31", // TODO: determine date properly
			},
			EnsurePresent:  row.EnsurePresent,
			EnsureNotStale: row.EnsureNotStale,
		})
	}

	// Make the call
	result, err := h.measurementClient.Measure(measurementRequest)
	if err != nil {
		fmt.Println(err)
	}
	// TODO: handle error and return correct status code.

	// Build the response
	response := BatchResponse{}
	for _, row := range result.Rows {
		response.Rows = append(response.Rows, Response{TotalEmissions: row.TotalEmissions})
	}

	json.NewEncoder(w).Encode(response)
}
