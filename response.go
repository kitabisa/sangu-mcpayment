package mcpayment

// TokenizeRegResp response body for register token
type TokenizeRegResp struct {
	Error bool                `json:"error"`
	Data  TokenizeRegDataResp `json:"data"`
}

// TokenizeRegDataResp response inner data for register token
type TokenizeRegDataResp struct {
	SeamlessURL string `json:"seamless_url"`
	Message     string `json:"message"`
	ErrorCode   string `json:"error_code"`
}

// StatusResp response body for error / default response
type StatusResp struct {
	Error bool           `json:"error"`
	Code  int            `json:"code"`
	Data  StatusDataResp `json:"data"`
}

// StatusDataResp status response inner data
type StatusDataResp struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorcode"`
}

// TokenizeDetail token detail data from get / callback
type TokenizeDetail struct {
	Token          string `json:"token"`
	RegisterID     string `json:"register_id"`
	MaskedCardNo   string `json:"masked_card_no"`
	CardHolderName string `json:"card_holder_name"`
	CardExpDate    string `json:"card_exp_date"`
	CardBrand      string `json:"card_brand"`
	Status         string `json:"status"`
	BankIssuer     string `json:"bank_issuer"`
	StatusResp
}
