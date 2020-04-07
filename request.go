package mcpayment

// TokenizeRegisterReq body request for register token
type TokenizeRegisterReq struct {
	RegisterID  string    `json:"register_id" valid:"required,length(1|100)"`
	CallbackURL string    `json:"callback_url" valid:"required,url"`
	ReturnURL   string    `json:"return_url" valid:"required,url"`
	IsTrx       bool      `json:"is_transaction,omitempty" valid:"required"`
	TrxDetail   TrxDetail `json:"transaction,omitempty" valid:"-"`
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

// RecrCreateReq body request for create recurring
type RecrCreateReq struct {
	RegisterID         string           `json:"register_id" valid:"required,length(1|100)"`
	Name               string           `json:"name" valid:"required,length(1|100)"`
	Amount             float64          `json:"amount" valid:"required,range(1|999999999999999)"`
	Token              string           `json:"token" valid:"required,length(1|100)"`
	CallbackURL        string           `json:"callback_url" valid:"required,url"`
	Schedule           RecrSchCreateReq `json:"schedule" valid:"required"`
	MissedChargeAction string           `json:"missed_charge_action,omitempty" valid:"optional,in(ignore|stop)"`
	TotalRecurrence    int              `json:"total_recurrence,omitempty" valid:"optional,range(1|2147483647)"`
}

// RecrSchCreateReq body request detail for create recurring schedule
type RecrSchCreateReq struct {
	Interval     int    `json:"interval" valid:"range(1|365)"`
	IntervalUnit string `json:"interval_unit" valid:"in(day|week|month)"`
	StartTime    string `json:"start_time" valid:"rfc3339"`
}

// RecrUpdateReq body request for update recurring
type RecrUpdateReq struct {
	Name               string           `json:"name" valid:"required,length(1|100)"`
	Amount             float64          `json:"amount" valid:"required,range(1|999999999999999)"`
	Token              string           `json:"token" valid:"required,length(1|100)"`
	CallbackURL        string           `json:"callback_url" valid:"required,url"`
	Schedule           RecrSchUpdateReq `json:"schedule" valid:"required"`
	MissedChargeAction string           `json:"missed_charge_action" valid:"-"`
	TotalRecurrence    int              `json:"total_recurrence" valid:"-"`
}

// RecrSchUpdateReq body request detail for update recurring schedule
type RecrSchUpdateReq struct {
	Interval     int    `json:"interval" valid:"range(1|365)"`
	IntervalUnit string `json:"interval_unit" valid:"in(day|week|month)"`
}
