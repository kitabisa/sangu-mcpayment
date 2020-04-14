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

// TokenizationGateway ...
type TokenizationGateway struct {
	Client Client
}

// Register register token
func (g *TokenizationGateway) Register(req *TokenizeRegisterReq) (resp TokenizeRegResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrInvalidRequest, err.Error())
		return
	}

	fullPath := fmt.Sprintf("%s%s", g.Client.BaseURLToken, epTokenRegister)

	reqBodyJSON, err := json.Marshal(req)
	if err != nil {
		if g.Client.LogLevel > 0 {
			g.Client.Logger.Printf(PaymentName, " Marshalling body failed: %s\n", err)
		}
		return
	}

	err = g.Client.Call(http.MethodPost, fullPath, nil, strings.NewReader(string(reqBodyJSON)), &resp)
	return
}

// Get token
func (g *TokenizationGateway) Get(registerID string) (resp TokenizeDetail, err error) {
	fullPath := fmt.Sprintf("%s%s", g.Client.BaseURLToken, fmt.Sprintf(epTokenGet, registerID))
	err = g.Client.Call(http.MethodGet, fullPath, nil, nil, &resp)
	return
}

// Delete token
func (g *TokenizationGateway) Delete(token string) (resp TokenizeDetail, err error) {
	fullPath := fmt.Sprintf("%s%s", g.Client.BaseURLToken, fmt.Sprintf(epTokenDel, token))
	err = g.Client.Call(http.MethodDelete, fullPath, nil, nil, &resp)
	return
}
