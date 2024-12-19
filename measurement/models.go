package measurement

type Request struct {
	Country     string `json:"country"`
	Channel     string `json:"channel"`
	Impressions uint64 `json:"impressions"`
	InventoryId string `json:"inventoryId"`
	UtcDatetime string `json:"utcDatetime"`
}
type BatchRequest struct {
	Rows []Request `json:"rows"`
}

type Response struct {
	Coverage                Coverage                `json:"coverage"`
	Policies                []Policy                `json:"policies"`
	RequestID               string                  `json:"requestId"`
	TotalEmissions          float64                 `json:"totalEmissions"`
	TotalEmissionsBreakdown TotalEmissionsBreakdown `json:"totalEmissionsBreakdown"`
	Rows                    []Row                   `json:"rows"`
}

type Coverage struct {
	AdFormats        MetricData `json:"adFormats"`
	Channels         MetricData `json:"channels"`
	MediaOwners      MetricData `json:"mediaOwners"`
	Properties       MetricData `json:"properties"`
	Sellers          MetricData `json:"sellers"`
	TotalImpressions MetricData `json:"totalImpressions"`
	TotalRows        MetricData `json:"totalRows"`
}

type MetricData struct {
	Metric         string `json:"metric"`
	Generic        int    `json:"generic,omitempty"`
	Modeled        int    `json:"modeled,omitempty"`
	Skipped        int    `json:"skipped,omitempty"`
	Unknown        int    `json:"unknown,omitempty"`
	VendorSpecific int    `json:"vendorSpecific,omitempty"`
}

type Policy struct {
	Compliant    int    `json:"compliant"`
	Noncompliant int    `json:"noncompliant"`
	Policy       string `json:"policy"`
	PolicyOwner  string `json:"policyOwner"`
}

type TotalEmissionsBreakdown struct {
	Framework string            `json:"framework"`
	Totals    EmissionBreakdown `json:"totals"`
}

type EmissionBreakdown struct {
	AdSelection       float64 `json:"adSelection"`
	CreativeDelivery  float64 `json:"creativeDelivery"`
	MediaDistribution float64 `json:"mediaDistribution"`
}

type Row struct {
	EmissionsBreakdown EmissionsBreakdown `json:"emissionsBreakdown"`
	InventoryCoverage  string             `json:"inventoryCoverage"`
	TotalEmissions     float64            `json:"totalEmissions"`
	Internal           Internal           `json:"internal"`
}

type EmissionsBreakdown struct {
	Breakdown Breakdown `json:"breakdown"`
	Framework string    `json:"framework"`
}

type Breakdown struct {
	AdSelection       Component `json:"adSelection"`
	Compensated       Component `json:"compensated"`
	CreativeDelivery  Component `json:"creativeDelivery"`
	MediaDistribution Component `json:"mediaDistribution"`
}

type Component struct {
	Breakdown map[string]EmissionDetail `json:"breakdown"`
	Total     float64                   `json:"total"`
}

type EmissionDetail struct {
	Emissions float64 `json:"emissions"`
	Provider  string  `json:"provider,omitempty"`
}

type Internal struct {
	CountryRegionGCO2PerKwh float64              `json:"countryRegionGCO2PerKwh"`
	CountryRegionCountry    string               `json:"countryRegionCountry"`
	Channel                 string               `json:"channel"`
	DeviceType              string               `json:"deviceType"`
	PropertyID              int                  `json:"propertyId"`
	PropertyInventoryType   string               `json:"propertyInventoryType"`
	PropertyName            string               `json:"propertyName"`
	BenchmarkPercentile     int                  `json:"benchmarkPercentile"`
	IsMFA                   bool                 `json:"isMFA"`
	PolicyEvaluationData    PolicyEvaluationData `json:"policyEvaluationData"`
}

type PolicyEvaluationData struct {
	PropertyID           int    `json:"propertyId"`
	IsMFA                bool   `json:"isMFA"`
	IsInventory          bool   `json:"isInventory"`
	Country              string `json:"country"`
	Channel              string `json:"channel"`
	ChannelStatus        string `json:"channelStatus"`
	BenchmarksPercentile int    `json:"benchmarksPercentile"`
}
