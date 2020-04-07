package mcpayment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
)

var (
	// EpRecurring endpoint for create recurring
	EpRecurring = "/subscriptions"

	// EpRecurringID endpoint for get recurring detail
	EpRecurringID = EpRecurring + "/%s"
)

// RecurringGateway gateway for recurring
type RecurringGateway struct {
	Client Client
}

// Create call create subscription API
func (r *RecurringGateway) Create(req *RecrCreateReq) (resp RecrResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrInvalidRequest, err.Error())
		return
	}

	fullPath := fmt.Sprintf("%s%s", r.Client.BaseURLRecurring, EpRecurring)

	reqBodyJSON, err := json.Marshal(req)
	if err != nil {
		if r.Client.LogLevel > 0 {
			r.Client.Logger.Printf(PaymentName, " Marshalling body failed: %s\n", err)
		}
		return
	}

	err = r.Client.Call(http.MethodPost, fullPath, nil, strings.NewReader(string(reqBodyJSON)), &resp)
	return
}

// Get call get subscription API
func (r *RecurringGateway) Get(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s", r.Client.BaseURLRecurring, fmt.Sprintf(EpRecurringID, registerID))
	err = r.Client.Call(http.MethodGet, fullPath, nil, nil, &resp)
	return
}

// Update call update subsciption API
func (r *RecurringGateway) Update(registerID string, req *RecrUpdateReq) (resp RecrResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrInvalidRequest, err.Error())
		return
	}

	fullPath := fmt.Sprintf("%s%s", r.Client.BaseURLRecurring, fmt.Sprintf(EpRecurringID, registerID))

	reqBodyJSON, err := json.Marshal(req)
	if err != nil {
		if r.Client.LogLevel > 0 {
			r.Client.Logger.Printf(PaymentName, " Marshalling body failed: %s\n", err)
		}
		return
	}

	err = r.Client.Call(http.MethodPatch, fullPath, nil, strings.NewReader(string(reqBodyJSON)), &resp)
	return
}

// Enable call enable subsciption API
func (r *RecurringGateway) Enable(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s/enable", r.Client.BaseURLRecurring, fmt.Sprintf(EpRecurringID, registerID))
	err = r.Client.Call(http.MethodPost, fullPath, nil, nil, &resp)
	return
}

// Disable call enable subsciption API
func (r *RecurringGateway) Disable(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s/disable", r.Client.BaseURLRecurring, fmt.Sprintf(EpRecurringID, registerID))
	err = r.Client.Call(http.MethodPost, fullPath, nil, nil, &resp)
	return
}
