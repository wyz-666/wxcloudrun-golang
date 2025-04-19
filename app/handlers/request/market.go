package request

type ReqMarket struct {
	Product      string `json:"product" form:"product"`
	Date         string `json:"date" form:"date"`
	LowerPrice   string `json:"lowerPrice" form:"lowerPrice"`
	HigherPrice  string `json:"higherPrice" form:"higherPrice"`
	ClosingPrice string `json:"closingPrice" form:"closingPrice"`
}
