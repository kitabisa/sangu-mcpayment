package mcpayment

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
	Message            string             `json:"message"`
}

// RecrSchdDetailResp response body for recurring schedule detail data
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
	RecurNo   int     `json:"recur_no"`
	Status    string  `json:"status"`
	ChargedAt string  `json:"charged_at"`
}
