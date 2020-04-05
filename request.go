package mcpayment

// TokenizeRegisterReq body request for register token
type TokenizeRegisterReq struct {
	RegisterID  string    `json:"register_id" valid:"required,length(1|1024)"`
	CallbackURL string    `json:"callback_url" valid:"required,url"`
	ReturnURL   string    `json:"return_url" valid:"required,url"`
	IsTrx       bool      `json:"is_transaction" valid:"required"`
	TrxDetail   TrxDetail `json:"transaction" valid:"-"`
}

// TransactionDetail body request for transactiond detail
type TrxDetail struct {
	Amount float64 `json:"amount" valid:"-"`
	Desc   string  `json:"description" valid:"-"`
}

// TokenizeGetReq param for get tokenize
type TokenizeGetReq struct {
	RegisterID string `json:"register_id" valid:"required"`
}

// TokenizeDelReq param for delete tokenize
type TokenizeDelReq struct {
	RegisterID string `json:"register_id" valid:"required"`
}
