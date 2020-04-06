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

type RegTokenCase struct {
	SignKey      string
	BaseURLToken string
	In           TokenizeRegisterReq
	Out          TokenizeRegResp
	Err          error
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
	var RegTokenTestCases = []RegTokenCase{
		{
			// OK
			SignKey: mc.tokenGateway.Client.XSignKey,
			In: TokenizeRegisterReq{
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				RegisterID:  mc.newRegisterID,
				ReturnURL:   "https://google.com",
			},
			Err: nil,
			Out: TokenizeRegResp{
				Error: false,
			},
		},
		{
			// Error Validation
			SignKey: mc.tokenGateway.Client.XSignKey,
			In: TokenizeRegisterReq{
				CallbackURL: "not-url-format",
				RegisterID:  mc.newRegisterID,
				ReturnURL:   "https://google.com",
			},
			Err: ErrInvalidRequest,
			Out: TokenizeRegResp{},
		},
		{
			// Error SignKey
			SignKey: "any-sign-key",
			In: TokenizeRegisterReq{
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				RegisterID:  mc.newRegisterID,
				ReturnURL:   "https://google.com",
			},
			Err: nil,
			Out: TokenizeRegResp{
				Error: true,
				Data: TokenizeRegDataResp{
					ErrorCode: "INTERNAL_SERVER_ERROR",
				},
			},
		},
	}

	for _, test := range RegTokenTestCases {
		mc.tokenGateway.Client.XSignKey = test.SignKey
		resp, err := mc.tokenGateway.Register(&test.In)
		assert.Equal(mc.T(), test.Out.Error, resp.Error)

		if err == nil {
			assert.Equal(mc.T(), test.Err, err)
		} else {
			assert.Error(mc.T(), test.Err, err)
		}

		if resp.Error {
			assert.Equal(mc.T(), test.Out.Data.ErrorCode, resp.Data.ErrorCode)
		}

	}
}

func (mc *McPaymentTestSuite) TestGetToken() {
	resp, err := mc.tokenGateway.Get(mc.registeredID)
	assert.Equal(mc.T(), nil, err)
	assert.Equal(mc.T(), false, resp.Error)
	// TODO: add test case for fail
}

func (mc *McPaymentTestSuite) TestDeleteToken() {
	resp, err := mc.tokenGateway.Delete("31a2d102c892eaad241465ce830b301fdf5c9ab5be76b7d20c1a75c122ea4d78")
	assert.Equal(mc.T(), nil, err)
	assert.Equal(mc.T(), false, resp.Error)
	// TODO: add test case for fail
}
