package response

type ResLogin struct {
	Uuid   string `json:uuid`
	UserID string `json:"userId"`
	Token  string `json:"token"`
}
