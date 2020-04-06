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
	tokenGateway  TokenizationGateway
	conf          Configs
	newRegisterID string
}

type Configs struct {
	BaseURLToken    string
	XSignKey        string
	ReturnURL       string
	RegisteredID    string
	RegisteredToken string
}

// RegTokenCase case struct for register token
type RegTokenCase struct {
	SignKey string
	In      TokenizeRegisterReq
	Out     TokenizeRegResp
	Err     error
}

// GetDelTokenCase case struct for get and del token
type GetDelTokenCase struct {
	SignKey string
	In      string
	Out     TokenizeDetail
	Err     error
}

// GetConfig get config for test
func GetConfig() (Configs, error) {
	theToml, err := ioutil.ReadFile("credential_test.toml")
	if err != nil {
		return Configs{}, err
	}

	var conf Configs
	if _, err := toml.Decode(string(theToml), &conf); err != nil {
		return Configs{}, err
	}

	return conf, nil
}

func TestMcPaymentTestSuite(t *testing.T) {
	suite.Run(t, new(McPaymentTestSuite))
}

func (mc *McPaymentTestSuite) SetupTest() {
	conf, err := GetConfig()
	if err != nil {
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
	mc.conf = conf
}

func (mc *McPaymentTestSuite) TestRegisterToken() {
	var RegTokenTestCases = []RegTokenCase{
		{
			// OK
			SignKey: mc.conf.XSignKey,
			In: TokenizeRegisterReq{
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				RegisterID:  mc.newRegisterID,
				ReturnURL:   mc.conf.ReturnURL,
			},
			Err: nil,
			Out: TokenizeRegResp{
				Error: false,
			},
		},
		{
			// Error Validation
			SignKey: mc.conf.XSignKey,
			In: TokenizeRegisterReq{
				CallbackURL: "not-url-format",
				RegisterID:  mc.newRegisterID,
				ReturnURL:   mc.conf.ReturnURL,
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
				ReturnURL:   mc.conf.ReturnURL,
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
	var getTokenTestCases = []GetDelTokenCase{
		{
			SignKey: mc.conf.XSignKey,
			In:      mc.conf.RegisteredID,
			Err:     nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: false,
				},
			},
		},
		{
			SignKey: mc.conf.XSignKey,
			In:      randstr.String(20),
			Err:     nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: true,
					Data: TokenizeStatusDataResp{
						ErrorCode: "NOT_FOUND",
					},
				},
			},
		},
		{
			SignKey: randstr.String(20),
			In:      mc.conf.RegisteredID,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: true,
					Data: TokenizeStatusDataResp{
						ErrorCode: "INTERNAL_SERVER_ERROR",
					},
				},
			},
		},
	}

	for _, test := range getTokenTestCases {
		mc.tokenGateway.Client.XSignKey = test.SignKey
		resp, err := mc.tokenGateway.Get(test.In)
		assert.Equal(mc.T(), test.Err, err)
		assert.Equal(mc.T(), test.Out.Error, resp.Error)

		if resp.Error {
			assert.Equal(mc.T(), test.Out.Data.ErrorCode, resp.Data.ErrorCode)
		}
	}
}

func (mc *McPaymentTestSuite) TestDeleteToken() {
	var delTokenTestCases = []GetDelTokenCase{
		{
			SignKey: randstr.String(20),
			In:      mc.conf.RegisteredID,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: true,
					Data: TokenizeStatusDataResp{
						ErrorCode: "INTERNAL_SERVER_ERROR",
					},
				},
			},
		},
		{
			SignKey: mc.conf.XSignKey,
			In:      randstr.String(20),
			Err:     nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: true,
					Data: TokenizeStatusDataResp{
						ErrorCode: "NOT_FOUND",
					},
				},
			},
		},
		{
			SignKey: mc.conf.XSignKey,
			In:      mc.conf.RegisteredToken,
			Err:     nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: false,
				},
			},
		},
	}

	for _, test := range delTokenTestCases {
		mc.tokenGateway.Client.XSignKey = test.SignKey
		resp, err := mc.tokenGateway.Delete(test.In)
		assert.Equal(mc.T(), test.Err, err)
		assert.Equal(mc.T(), test.Out.Error, resp.Error)

		if resp.Error {
			assert.Equal(mc.T(), test.Out.Data.ErrorCode, resp.Data.ErrorCode)
		}
	}
}
