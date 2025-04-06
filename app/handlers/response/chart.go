package response

type MonthlyPriceStats struct {
	Month    string  `json:"month"`
	AvgLow   float64 `json:"avgLow"`
	AvgHigh  float64 `json:"avgHigh"`
	AvgPrice float64 `json:"avgPrice"`
}

type GECMonthlyPriceStats struct {
	Month    string  `json:"month"`
	Type     string  `json:"type"`
	AvgPrice float64 `json:"avgPrice"`
}
