package mcpayment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
)

var (
	epCreate  = "/create"
	epGet     = "/detail/%s"
	epUpdate  = "/update/%s"
	epEnable  = epUpdate + "/enable"
	epDisable = epUpdate + "/disable"
	epFinish  = epUpdate + "/finish"
)

// IRecurringGateway interface for recurringGateway
type IRecurringGateway interface {
	Create(req *RecrCreateReq) (resp RecrResp, err error)
	Get(registerID string) (resp RecrResp, err error)
	Update(registerID string, req *RecrUpdateReq) (resp RecrResp, err error)
	Enable(registerID string) (resp RecrResp, err error)
	Disable(registerID string) (resp RecrResp, err error)
	Finish(registerID string) (resp RecrResp, err error)
	ValidateSignKey(req RecrCallbackReq) bool
}

// RecurringGateway gateway for recurring
type recurringGateway struct {
	client Client
}

// NewRecurringGateway create instance of IRecurringGateway
func NewRecurringGateway(client Client) IRecurringGateway {
	return &recurringGateway{
		client: client,
	}
}

// Create call create subscription API
func (r *recurringGateway) Create(req *RecrCreateReq) (resp RecrResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrInvalidRequest, err.Error())
		return
	}

	fullPath := fmt.Sprintf("%s%s", r.client.BaseURLRecurring, epCreate)

	reqBodyJSON, err := json.Marshal(req)
	if err != nil {
		if r.client.LogLevel > 0 {
			r.client.Logger.Printf(PaymentName, " Marshalling body failed: %s\n", err)
		}
		return
	}

	err = r.client.Call(http.MethodPost, fullPath, nil, strings.NewReader(string(reqBodyJSON)), &resp)
	return
}

// Get call get subscription API
func (r *recurringGateway) Get(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s", r.client.BaseURLRecurring, fmt.Sprintf(epGet, registerID))
	err = r.client.Call(http.MethodGet, fullPath, nil, nil, &resp)
	return
}

// Update call update subsciption API
func (r *recurringGateway) Update(registerID string, req *RecrUpdateReq) (resp RecrResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrInvalidRequest, err.Error())
		return
	}

	fullPath := fmt.Sprintf("%s%s", r.client.BaseURLRecurring, fmt.Sprintf(epUpdate, registerID))

	reqBodyJSON, err := json.Marshal(req)
	if err != nil {
		if r.client.LogLevel > 0 {
			r.client.Logger.Printf(PaymentName, " Marshalling body failed: %s\n", err)
		}
		return
	}

	err = r.client.Call(http.MethodPatch, fullPath, nil, strings.NewReader(string(reqBodyJSON)), &resp)
	return
}

// Enable call enable subsciption API
func (r *recurringGateway) Enable(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s", r.client.BaseURLRecurring, fmt.Sprintf(epEnable, registerID))
	err = r.client.Call(http.MethodPost, fullPath, nil, nil, &resp)
	return
}

// Disable call enable subsciption API
func (r *recurringGateway) Disable(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s", r.client.BaseURLRecurring, fmt.Sprintf(epDisable, registerID))
	err = r.client.Call(http.MethodPost, fullPath, nil, nil, &resp)
	return
}

// Finish call finish subsciption API
func (r *recurringGateway) Finish(registerID string) (resp RecrResp, err error) {
	fullPath := fmt.Sprintf("%s%s", r.client.BaseURLRecurring, fmt.Sprintf(epFinish, registerID))
	err = r.client.Call(http.MethodPost, fullPath, nil, nil, &resp)
	return
}

// ValidateSignKey validate signature key on callback request
func (r *recurringGateway) ValidateSignKey(req RecrCallbackReq) bool {
	return validateSignatureKeyRecurringTransaction(r.client.XSignKey, req.RegisterID, req.SignatureKey, req.Amount)
}
