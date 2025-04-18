package request

type ReqMarket struct {
	Date         string `json:"date" form:"date"`
	LowerPrice   string `json:"lowerPrice" form:"lowerPrice"`
	HigherPrice  string `json:"higherPrice" form:"higherPrice"`
	ClosingPrice string `json:"closingPrice" form:"closingPrice"`
}
