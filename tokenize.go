package mcpayment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
)

var (
	epTokenRegister = "/request_tokenization"
	epTokenGet      = "/detail/%s"
	epTokenDel      = "/delete/%s"
)

// ITokenizationGateway interface for TokenizationGateway
type ITokenizationGateway interface {
	Register(req *TokenizeRegisterReq) (resp TokenizeRegResp, err error)
	Get(registerID string) (resp TokenizeDetail, err error)
	Delete(token string) (resp TokenizeDetail, err error)
}

// tokenizationGateway ...
type tokenizationGateway struct {
	client Client
}

func NewTokenizationGateway(client Client) ITokenizationGateway {
	return &tokenizationGateway{
		client: client,
	}
}

// Register register token
func (g *tokenizationGateway) Register(req *TokenizeRegisterReq) (resp TokenizeRegResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrInvalidRequest, err.Error())
		return
	}

	fullPath := fmt.Sprintf("%s%s", g.client.BaseURLToken, epTokenRegister)

	reqBodyJSON, err := json.Marshal(req)
	if err != nil {
		if g.client.LogLevel > 0 {
			g.client.Logger.Printf(PaymentName, " Marshalling body failed: %s\n", err)
		}
		return
	}

	err = g.client.Call(http.MethodPost, fullPath, nil, strings.NewReader(string(reqBodyJSON)), &resp)
	return
}

// Get token
func (g *tokenizationGateway) Get(registerID string) (resp TokenizeDetail, err error) {
	fullPath := fmt.Sprintf("%s%s", g.client.BaseURLToken, fmt.Sprintf(epTokenGet, registerID))
	err = g.client.Call(http.MethodGet, fullPath, nil, nil, &resp)
	return
}

// Delete token
func (g *tokenizationGateway) Delete(token string) (resp TokenizeDetail, err error) {
	fullPath := fmt.Sprintf("%s%s", g.client.BaseURLToken, fmt.Sprintf(epTokenDel, token))
	err = g.client.Call(http.MethodDelete, fullPath, nil, nil, &resp)
	return
}
