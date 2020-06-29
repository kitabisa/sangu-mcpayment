package mcpayment

import (
	"gopkg.in/guregu/null.v3"
)

// RecrCreateReq body request for create recurring
type RecrCreateReq struct {
	RegisterID         string           `json:"register_id" valid:"required,length(1|100)"`
	Name               string           `json:"name" valid:"required,length(1|100)"`
	Amount             float64          `json:"amount" valid:"required,range(1|999999999999999)"`
	Token              string           `json:"token" valid:"required,length(1|500)"`
	CallbackURL        string           `json:"callback_url" valid:"required,url"`
	Schedule           RecrSchCreateReq `json:"schedule" valid:"required"`
	MissedChargeAction *string          `json:"missed_charge_action,omitempty" valid:"-,in(ignore|stop)"`
	TotalRecurrence    *int             `json:"total_recurrence,omitempty" valid:"-,range(1|2147483647)"`
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
	Token              string           `json:"token" valid:"required,length(1|500)"`
	CallbackURL        string           `json:"callback_url" valid:"required,url"`
	Schedule           RecrSchUpdateReq `json:"schedule" valid:"required"`
	MissedChargeAction *string          `json:"missed_charge_action" valid:"-,in(ignore|stop)"`
	TotalRecurrence    *int             `json:"total_recurrence" valid:"-"`
}

// RecrSchUpdateReq body request detail for update recurring schedule
type RecrSchUpdateReq struct {
	Interval     int    `json:"interval" valid:"range(1|365)"`
	IntervalUnit string `json:"interval_unit" valid:"in(day|week|month)"`
}

// RecrCallbackReq callback body request
type RecrCallbackReq struct {
	ID           string      `json:"id"`
	RegisterID   string      `json:"register_id"`
	RecurNo      int         `json:"recur_no"`
	Status       string      `json:"status"`
	Amount       float64     `json:"amount"`
	Message      null.String `json:"message"`
	CreatedAt    string      `json:"created_at"`
	SignatureKey string      `json:"signature_key"`
}
