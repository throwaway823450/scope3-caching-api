package api

type Request struct {
	InventoryId string `json:"inventoryId"`
}

type BatchRequest struct {
	Rows []Request `json:"rows"`
}

type Response struct {
	TotalEmissions float64 `json:"totalEmissions"`
}

type BatchResponse struct {
	Rows []Response `json:"rows"`
}
