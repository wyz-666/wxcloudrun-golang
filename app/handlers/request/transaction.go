package request

type ReqTransaction struct {
	UserID   string `json:"userId" form:"userId"`
	Project  string `json:"project" form:"project"`
	Type     string `json:"type" form:"type"`
	Price    string `json:"price" form:"price"`
	TxVolume string `json:"txVolume" form:"txVolume"`
}
