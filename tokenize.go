package mcpayment

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
)

var (
	epTokenRegister = "/request_tokenization"
)

// TokenizationGateway ...
type TokenizationGateway struct {
	Client Client
}

// Register register token
func (g *TokenizationGateway) Register(req TokenizeRegisterReq) (resp TokenizeRegResp, err error) {
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		err = errors.New(err.Error())
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
	fullPath := fmt.Sprintf("%s/%s", g.Client.BaseURLToken, registerID)
	fmt.Println(fullPath)
	err = g.Client.Call(http.MethodGet, fullPath, nil, nil, &resp)
	return
}

// Delete token
func (g *TokenizationGateway) Delete(token string) (resp TokenizeDetail, err error) {
	fullPath := fmt.Sprintf("%s/%s", g.Client.BaseURLToken, token)
	fmt.Println(fullPath)
	err = g.Client.Call(http.MethodDelete, fullPath, nil, nil, &resp)
	return
}
