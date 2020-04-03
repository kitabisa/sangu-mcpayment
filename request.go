package mcpayment

// TokenizeRegisterReq body request for register token
type TokenizeRegisterReq struct {
	RegisterID  string `json:"register_id" valid:"required"`
	CallbackURL string `json:"callback_url" valid:"required"`
}

// TokenizeGetReq param for get tokenize
type TokenizeGetReq struct {
	RegisterID string `json:"register_id" valid:"required"`
}

// TokenizeDelReq param for delete tokenize
type TokenizeDelReq struct {
	RegisterID string `json:"register_id" valid:"required"`
}
