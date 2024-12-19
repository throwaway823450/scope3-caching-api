package measurement

import (
	"fmt"
	"time"
)

const (
	defaultExpiry = 10 * time.Second
)

type CachingClient interface {
	Measure(BatchCachingRequest) (*Response, error)
}

func NewCachingClient(client Client) CachingClient {
	return &cachingClientImpl{client: client, cache: NewCache()}
}

type cachingClientImpl struct {
	client Client
	cache  *Cache
}

func (c *cachingClientImpl) Measure(batchRequest BatchCachingRequest) (*Response, error) {
	results := []Row{}
	toRefresh := make(map[int]Request)
	toRefreshAsync := []Request{}
	for i, row := range batchRequest.Rows {
		key := row.Request.InventoryId
		cachedItem, exists := c.cache.GetWithTimestamp(key)
		if !exists {
			// Set an empty placeholder
			results = append(results, Row{})
			if row.EnsurePresent {
				// It doesn't exist, but the client wants to wait for the result
				toRefresh[i] = row.Request
			} else {
				// It doesn't exist and the client doesn't need to the result now.
				// Schedule an update in the background.
				toRefreshAsync = append(toRefreshAsync, row.Request)
			}
		} else {
			isStale := time.Now().After(cachedItem.EntryTime.Add(defaultExpiry))
			if !isStale {
				// Not stale so return
				results = append(results, cachedItem.Data)
			} else {
				if row.EnsureNotStale {
					// Set an empty placeholder
					results = append(results, Row{})
					if row.EnsurePresent {
						toRefresh[i] = row.Request
					} else {
						// Client just doesn't want stale data, OK with empty data.
						// Refresh it in the background
						toRefreshAsync = append(toRefreshAsync, row.Request)
					}

				} else {
					// Client is OK with stale data
					results = append(results, cachedItem.Data)
					// Refresh it in the background
					toRefreshAsync = append(toRefreshAsync, row.Request)
				}
			}
		}
	}

	// Get data for responses that need to refreshed right now
	if len(toRefresh) > 0 {
		fmt.Printf("Refreshing %d item(s) in the main thread\n", len(toRefresh))
		var toRefreshRequest BatchRequest
		for _, request := range toRefresh {
			toRefreshRequest.Rows = append(toRefreshRequest.Rows, request)
		}
		response, _ := c.client.Measure(toRefreshRequest)
		// TODO: handle error
		responseIndex := 0
		for resultsIndex, request := range toRefresh {
			fmt.Println("Refreshed", request.InventoryId)
			results[resultsIndex] = response.Rows[responseIndex]
			c.cache.Set(request.InventoryId, response.Rows[responseIndex])
			responseIndex++
		}
	}

	finalResponse := Response{}
	finalResponse.Rows = append(finalResponse.Rows, results...)

	if len(toRefreshAsync) > 0 {
		go func() {
			fmt.Printf("Refreshing %d item(s) in the background\n", len(toRefreshAsync))
			response, _ := c.client.Measure(BatchRequest{Rows: toRefreshAsync})
			for resultsIndex, request := range toRefreshAsync {
				fmt.Println("Refreshed", request.InventoryId)
				results[resultsIndex] = response.Rows[resultsIndex]
				c.cache.Set(request.InventoryId, response.Rows[resultsIndex])
			}
		}()
	}

	return &finalResponse, nil
}
