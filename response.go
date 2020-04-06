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
	MaskedCardNo   string `json:"masked_card_number"`
	CardHolderName string `json:"card_holder_name"`
	CardExpDate    string `json:"card_exp_date"`
	CardBrand      string `json:"card_brand"`
	Status         string `json:"status"`
	BankIssuer     string `json:"bank_issuer"`
	StatusResp
}

// RecrResp response body for recurring
type RecrResp struct {
	Error   bool           `json:"error"`
	Message string         `json:"message"`
	Data    RecrDetailResp `json:"data"`
}

// RecrDetailResp response body for recurring detail data
type RecrDetailResp struct {
	ID                 string             `json:"id"`
	Status             string             `json:"status"`
	CreatedAt          string             `json:"created at"`
	RegisterID         string             `json:"register_id"`
	Name               string             `json:"name"`
	Amount             float64            `json:"amount"`
	Token              string             `json:"token"`
	CallbackURL        string             `json:"callback_url"`
	Schedule           RecrSchdDetailResp `json:"schedule"`
	MissedChargeAction string             `json:"missed_charge_action"`
	TotalRecurrence    int                `json:"total_recurrence"`
	Transactions       RecrTrxResp        `json:"transactions"`
}

// RecrSchdDetail response body for recurring schedule detail data
type RecrSchdDetailResp struct {
	Interval     int    `json:"interval"`
	IntervalUnit string `json:"interval_unit"`
	StartTime    string `json:"start_time"`
	Previous     string `json:"previous"`
	Next         string `json:"next"`
}

// RecrTrxResp response body for recurring transactions data
type RecrTrxResp struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	RecurNo   int     `json:"recur no"`
	Status    string  `json:"status"`
	ChargedAt string  `json:"charged at"`
}
