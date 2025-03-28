package response

type ResLogin struct {
	UserID string `json:"uuid"`
	Token  string `json:"token"`
}
