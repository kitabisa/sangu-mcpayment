package mcpayment

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/thanhpk/randstr"
)

// rename credential_test.toml.sample to credential_test.toml first
// fill the credential needed

type TokenizeTestSuite struct {
	suite.Suite
	tokenGateway  ITokenizationGateway
	conf          Configs
	newRegisterID string
}

// RegTokenCase case struct for register token
type RegTokenCase struct {
	Name string
	In   TokenizeRegisterReq
	Out  TokenizeRegResp
	Err  error
}

// GetDelTokenCase case struct for get and del token
type GetDelTokenCase struct {
	Name string
	In   string
	Out  TokenizeDetail
	Err  error
}

func TestTokenizeTestSuite(t *testing.T) {
	suite.Run(t, new(TokenizeTestSuite))
}

func (mc *TokenizeTestSuite) SetupTest() {
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

	mc.tokenGateway = NewTokenizationGateway(client)
	mc.newRegisterID = randstr.String(5)
	mc.conf = conf
}

func (mc *TokenizeTestSuite) TestRegisterToken() {
	testName := "Tokenize_Register:%s"
	var RegTokenTestCases = []RegTokenCase{
		{
			Name: fmt.Sprintf(testName, "OK"),
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
			Name: fmt.Sprintf(testName, "Error_Validation"),
			In: TokenizeRegisterReq{
				CallbackURL: "not-url-format",
				RegisterID:  mc.newRegisterID,
				ReturnURL:   mc.conf.ReturnURL,
			},
			Err: ErrInvalidRequest,
			Out: TokenizeRegResp{},
		},
	}

	for _, test := range RegTokenTestCases {
		resp, err := mc.tokenGateway.Register(&test.In)
		assert.Equal(mc.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(mc.T(), test.Err, err, test.Name)
		}

		if resp.Error {
			assert.Equal(mc.T(), test.Out.Data.ErrorCode, resp.Data.ErrorCode, test.Name)
		}

	}
}

func (mc *TokenizeTestSuite) TestGetToken() {
	testName := "Tokenize_Get:%s"
	var getTokenTestCases = []GetDelTokenCase{
		{
			Name: fmt.Sprintf(testName, "OK"),
			In:   mc.conf.RegisteredID,
			Err:  nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: false,
				},
			},
		},
		{
			Name: fmt.Sprintf(testName, ErrCodeNotFound),
			In:   randstr.String(20),
			Err:  nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: true,
					Data: TokenizeStatusDataResp{
						ErrorCode: ErrCodeNotFound,
					},
				},
			},
		},
	}

	for _, test := range getTokenTestCases {
		resp, err := mc.tokenGateway.Get(test.In)
		assert.Equal(mc.T(), test.Err, err, test.Name)
		assert.Equal(mc.T(), test.Out.Error, resp.Error, test.Name)

		if resp.Error {
			assert.Equal(mc.T(), test.Out.Data.ErrorCode, resp.Data.ErrorCode, test.Name)
		}
	}
}

func (mc *TokenizeTestSuite) TestDeleteToken() {
	testName := "Tokenize_Del:%s"
	var delTokenTestCases = []GetDelTokenCase{
		{
			Name: fmt.Sprintf(testName, ErrCodeNotFound),
			In:   randstr.String(20),
			Err:  nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: true,
					Data: TokenizeStatusDataResp{
						ErrorCode: ErrCodeNotFound,
					},
				},
			},
		},
		{
			Name: fmt.Sprintf(testName, "OK"),
			In:   mc.conf.RegisteredToken,
			Err:  nil,
			Out: TokenizeDetail{
				TokenizeStatusResp: TokenizeStatusResp{
					Error: false,
				},
			},
		},
	}

	for _, test := range delTokenTestCases {
		resp, err := mc.tokenGateway.Delete(test.In)
		assert.Equal(mc.T(), test.Err, err, test.Name)
		assert.Equal(mc.T(), test.Out.Error, resp.Error, test.Name)

		if resp.Error {
			assert.Equal(mc.T(), test.Out.Data.ErrorCode, resp.Data.ErrorCode, test.Name)
		}
	}
}
