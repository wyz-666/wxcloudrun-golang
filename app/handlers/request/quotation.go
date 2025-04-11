package request

import "time"

type ReqQuotation struct {
	FormName    string    `json:"formName" form:"formName"`
	UserID      string    `json:"userId" form:"userId"`
	Product     string    `json:"product" form:"product"`
	Type        string    `json:"type" form:"type"`
	LowerPrice  string    `json:"lowerPrice" form:"lowerPrice"`
	HigherPrice string    `json:"higherPrice" form:"higherPrice"`
	Price       string    `json:"price" form:"price"`
	TxVolume    string    `json:"txVolume" form:"txVolume"`
	NowTime     time.Time `json:"nowTime" form:"nowTime"`
}

type ReqApproveQuotation struct {
	QID  string `json:"qid" form:"qid"`
	Type string `json:"type" form:"type"`
}
