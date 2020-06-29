package mcpayment

// TokenizeRegisterReq body request for register token
type TokenizeRegisterReq struct {
	RegisterID  string    `json:"register_id" valid:"required,length(1|100)"`
	CallbackURL string    `json:"callback_url" valid:"required,url"`
	ReturnURL   string    `json:"return_url" valid:"required,url"`
	IsTrx       bool      `json:"is_transaction,omitempty" valid:"-"`
	TrxDetail   TrxDetail `json:"transaction,omitempty" valid:"-"`
}

// TrxDetail body request for transactiond detail
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

// TokenizeCallbackReq body request on callback tokenization
type TokenizeCallbackReq struct {
	Token          string  `json:"token"`
	RegisterID     string  `json:"register_id"`
	MaskedCardNo   string  `json:"masked_card_number"`
	CardHolderName string  `json:"card_holder_name"`
	CardExpDate    string  `json:"card_exp_date"`
	CardBrand      string  `json:"card_brand"`
	Status         string  `json:"status"`
	BankIssuer     string  `json:"bank_issuer"`
	Amount         float64 `json:"amount"`
	SignatureKey   string  `json:"signature_key"`
	ErrorCode      string  `json:"error_code"`
	Message        string  `json:"message"`
}
