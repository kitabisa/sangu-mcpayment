package mcpayment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	epTokenRegister = "/request_tokenization"
	epTokenBase     = "/tokenization/%s"
)

// TokenizationGateway ...
type TokenizationGateway struct {
	Client Client
}

// Register register token
func (g *TokenizationGateway) Register(req TokenizeRegisterReq) (resp TokenizeRegResp, err error) {
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
	fullPath := fmt.Sprintf("%s%s", g.Client.BaseURLToken, epTokenBase, registerID)
	err = g.Client.Call(http.MethodGet, fullPath, nil, nil, &resp)
	return
}

// Delete token
func (g *TokenizationGateway) Delete(registerID string) (resp TokenizeDetail, err error) {
	fullPath := fmt.Sprintf("%s%s", g.Client.BaseURLToken, epTokenBase, registerID)
	err = g.Client.Call(http.MethodDelete, fullPath, nil, nil, &resp)
	return
}
