package mcpayment

// TokenizeRegResp response body for register token
type TokenizeRegResp struct {
	Error   bool                `json:"error"`
	Message string              `json:"message"`
	Data    TokenizeRegDataResp `json:"data"`
}

// TokenizeRegDataResp response inner data for register token
type TokenizeRegDataResp struct {
	SeamlessURL string `json:"seamless_url"`
	ExpiredDate string `json:"expired_date"`
	Message     string `json:"message"`
	ErrorCode   string `json:"error_code"`
}

// TokenizeStatusResp response body for error / default response on GET and DEL request
type TokenizeStatusResp struct {
	Error   bool                   `json:"error"`
	Message string                 `json:"message"`
	Data    TokenizeStatusDataResp `json:"data"`
}

// TokenizeStatusDataResp status response inner data
type TokenizeStatusDataResp struct {
	ErrorCode string                     `json:"error_code"`
	Body      TokenizeStatusDataBodyResp `json:"body"`
}

// TokenizeStatusDataBodyResp body on status response inner data
type TokenizeStatusDataBodyResp struct {
	Token string `json:"token"`
}

// TokenizeDetail token detail data from get / callback
type TokenizeDetail struct {
	Token          string  `json:"token"`
	RegisterID     string  `json:"register_id"`
	MaskedCardNo   string  `json:"masked_card_number"`
	CardHolderName string  `json:"card_holder_name"`
	CardExpDate    string  `json:"card_exp_date"`
	CardBrand      string  `json:"card_brand"`
	Status         string  `json:"status"`
	BankIssuer     string  `json:"bank_issuer"`
	Amount         float64 `json:"amount"`
	TokenizeStatusResp
}
