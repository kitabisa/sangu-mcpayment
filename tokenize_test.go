package mcpayment

import (
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/thanhpk/randstr"
)

// rename credential_test.toml.sample to credential_test.toml first
// fill the credential needed

type McPaymentTestSuite struct {
	suite.Suite
	tokenGateway    TokenizationGateway
	newRegisterID   string
	returnURL       string
	registeredID    string
	registeredToken string
}

type configs struct {
	BaseURLToken    string
	XSignKey        string
	ReturnURL       string
	RegisteredID    string
	RegisteredToken string
}

func TestMcPaymentTestSuite(t *testing.T) {
	suite.Run(t, new(McPaymentTestSuite))
}

func (mc *McPaymentTestSuite) SetupTest() {
	theToml, err := ioutil.ReadFile("credential_test.toml")
	if err != nil {
		mc.T().Log(err)
		mc.T().FailNow()
	}

	var conf configs
	if _, err := toml.Decode(string(theToml), &conf); err != nil {
		mc.T().Log(err)
		mc.T().FailNow()
	}

	client := NewClient()
	client.BaseURLToken = conf.BaseURLToken
	client.XSignKey = conf.XSignKey
	client.IsEnvProduction = false
	client.LogLevel = 3

	mc.tokenGateway = TokenizationGateway{Client: client}
	mc.newRegisterID = randstr.String(5)
	mc.returnURL = conf.ReturnURL
	mc.registeredID = conf.RegisteredID
	mc.registeredToken = conf.RegisteredToken
}

func (mc *McPaymentTestSuite) TestRegisterToken() {
	req := TokenizeRegisterReq{
		CallbackURL: "https://mcpayment.free.beeceptor.com",
		RegisterID:  mc.newRegisterID,
		ReturnURL:   "https://google.com",
	}

	// TODO: add test case for fail
	resp, err := mc.tokenGateway.Register(req)

	assert.Equal(mc.T(), nil, err)
	assert.Equal(mc.T(), false, resp.Error)
}

func (mc *McPaymentTestSuite) TestGetToken() {
	// TODO: still error
	resp, err := mc.tokenGateway.Get(randstr.String(5))
	assert.Equal(mc.T(), nil, err)
	assert.Equal(mc.T(), false, resp.Error)
}

func (mc *McPaymentTestSuite) TestDeleteToken() {
	// TODO: still error
	resp, err := mc.tokenGateway.Delete(randstr.String(5))
	assert.Equal(mc.T(), nil, err)
	assert.Equal(mc.T(), false, resp.Error)
}
