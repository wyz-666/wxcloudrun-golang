package request

type ReqUser struct {
	Account         string `json:"account" form:"account"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Type            int    `json:"type" form:"type"`
	Name            string `json:"name" form:"name"`
	Company         string `json:"company" form:"company"`
	Phone           string `json:"phone" form:"phone"`
	Email           string `json:"email" form:"email"`
}

type ReqLogin struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type ReqApproveUser struct {
	UserID string `json:"userId" form:"userId"`
}
