package response

type MonthlyPriceStats struct {
	Month     string  `json:"month"`
	AvgLow    float64 `json:"avgLow"`
	AvgHigh   float64 `json:"avgHigh"`
	AvgPrice  float64 `json:"avgPrice"`
	LowIndex  float64 `json:"lowIndex"`
	HighIndex float64 `json:"highIndex"`
	MidIndex  float64 `json:"midIndex"`
	FitPrice  float64 `json:"fitPrice"`
}

type GECMonthlyPriceStats struct {
	Month    string  `json:"month"`
	Type     string  `json:"type"`
	AvgPrice float64 `json:"avgPrice"`
}
