package mcpayment

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/thanhpk/randstr"
)

type RecurringTestSuite struct {
	suite.Suite
	recGateway    IRecurringGateway
	conf          Configs
	newRegisterID string
}

func TestRecurringTestSuite(t *testing.T) {
	suite.Run(t, &RecurringTestSuite{})
}

func (r *RecurringTestSuite) SetupTest() {
	conf, err := GetConfig()
	if err != nil {
		r.T().Log(err)
		r.T().FailNow()
	}

	client := NewClient()
	client.BaseURLRecurring = conf.BaseURLRecurring
	client.XSignKey = conf.XSignKey
	client.IsEnvProduction = false
	client.LogLevel = 3

	r.recGateway = NewRecurringGateway(client)
	r.newRegisterID = randstr.String(5)
	r.conf = conf
}

// CaseCreateToken struct for test case create
type CaseCreateToken struct {
	Name string
	In   RecrCreateReq
	Out  RecrResp
	Err  error
}

// CaseUpdateToken struct for test case update
type CaseUpdateToken struct {
	Name    string
	In      RecrUpdateReq
	InRegID string
	Out     RecrResp
	Err     error
}

// CaseToken struct for other test case
type CaseToken struct {
	Name string
	In   string
	Out  RecrResp
	Err  error
}

func (r *RecurringTestSuite) TestCreate() {
	testName := "Recurring_Create:%s"
	var casesCreateToken = []CaseCreateToken{
		/* as this test case is not working on staging, but should be applied on prod
		{
			Name:    fmt.Sprintf(testName, "fail_starttime_less_than_today"),
			In: RecrCreateReq{
				RegisterID:  randstr.String(20),
				Name:        randstr.String(20),
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchCreateReq{
					Interval:     1,
					IntervalUnit: "month",
					StartTime:    time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339),
				},
			},
			Out: RecrResp{
				Error: true,
			},
			Err: nil,
		},*/
		{
			Name: fmt.Sprintf(testName, "fail_starttime_format"),
			In: RecrCreateReq{
				RegisterID:  randstr.String(20),
				Name:        randstr.String(20),
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchCreateReq{
					Interval:     1,
					IntervalUnit: "month",
					StartTime:    "1234567890",
				},
			},
			Out: RecrResp{},
			Err: ErrInvalidRequest,
		},
		{
			Name: fmt.Sprintf(testName, "fail_required"),
			In: RecrCreateReq{
				Name:        randstr.String(20),
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchCreateReq{
					Interval:     1,
					IntervalUnit: "month",
					StartTime:    time.Now().UTC().Add(5 * time.Second).Format(time.RFC3339),
				},
			},
			Out: RecrResp{
				Error: false,
			},
			Err: ErrInvalidRequest,
		},
		{
			Name: fmt.Sprintf(testName, "success"),
			In: RecrCreateReq{
				RegisterID:  randstr.String(20),
				Name:        randstr.String(20),
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchCreateReq{
					Interval:     1,
					IntervalUnit: "month",
					StartTime:    time.Now().UTC().Add(5 * time.Second).Format(time.RFC3339),
				},
			},
			Out: RecrResp{
				Error: false,
			},
			Err: nil,
		},
	}

	for _, test := range casesCreateToken {
		resp, err := r.recGateway.Create(&test.In)
		assert.Equal(r.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(r.T(), test.Err, err, test.Name)
		}
	}
}

func (r *RecurringTestSuite) TestGet() {
	nameTest := "Recurring_Get:%s"
	var casesToken = []CaseToken{
		{
			Name: fmt.Sprintf(nameTest, "OK"),
			In:   r.conf.RegisteredID,
			Out: RecrResp{
				Error: false,
			},
			Err: nil,
		},
		{
			Name: fmt.Sprintf(nameTest, "Not Found"),
			In:   randstr.String(20),
			Out: RecrResp{
				Error: true,
			},
			Err: nil,
		},
	}

	for _, test := range casesToken {
		resp, err := r.recGateway.Get(test.In)
		assert.Equal(r.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(r.T(), test.Err, err, test.Name)
		}
	}
}

func (r *RecurringTestSuite) TestUpdate() {
	testName := "Recurring_Update:%s"
	var casesUpdateToken = []CaseUpdateToken{
		{
			Name: fmt.Sprintf(testName, "fail_not_found"),
			In: RecrUpdateReq{
				Name:        randstr.String(20),
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchUpdateReq{
					Interval:     1,
					IntervalUnit: "month",
				},
			},
			InRegID: randstr.String(20),
			Out: RecrResp{
				Error: true,
			},
			Err: nil,
		},
		{
			Name: fmt.Sprintf(testName, "fail_required"),
			In: RecrUpdateReq{
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchUpdateReq{
					Interval:     1,
					IntervalUnit: "month",
				},
			},
			InRegID: r.conf.RegisteredID,
			Out:     RecrResp{},
			Err:     ErrInvalidRequest,
		},
		{
			Name: fmt.Sprintf(testName, "OK"),
			In: RecrUpdateReq{
				Name:        randstr.String(20),
				Amount:      10000,
				Token:       r.conf.RegisteredToken,
				CallbackURL: "https://mcpayment.free.beeceptor.com",
				Schedule: RecrSchUpdateReq{
					Interval:     1,
					IntervalUnit: "month",
				},
				MissedChargeAction: "ignore",
			},
			InRegID: r.conf.RegisteredID,
			Out: RecrResp{
				Error: false,
			},
			Err: nil,
		},
	}

	for _, test := range casesUpdateToken {
		resp, err := r.recGateway.Update(test.InRegID, &test.In)
		assert.Equal(r.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(r.T(), test.Err, err, test.Name)
		}
	}
}

func (r *RecurringTestSuite) TestEnable() {
	nameTest := "Recurring_Enable:%s"
	var casesToken = []CaseToken{
		{
			Name: fmt.Sprintf(nameTest, "OK"),
			In:   r.conf.RegisteredID,
			Out: RecrResp{
				Error: false,
			},
			Err: nil,
		},
		{
			Name: fmt.Sprintf(nameTest, "Not Found"),
			In:   randstr.String(20),
			Out: RecrResp{
				Error: true,
			},
			Err: nil,
		},
	}

	for _, test := range casesToken {
		resp, err := r.recGateway.Enable(test.In)
		assert.Equal(r.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(r.T(), test.Err, err, test.Name)
		}
	}
}

func (r *RecurringTestSuite) TestDisable() {
	nameTest := "Recurring_Disable:%s"
	var casesToken = []CaseToken{
		{
			Name: fmt.Sprintf(nameTest, "OK"),
			In:   r.conf.RegisteredID,
			Out: RecrResp{
				Error: false,
			},
			Err: nil,
		},
		{
			Name: fmt.Sprintf(nameTest, "Not Found"),
			In:   randstr.String(20),
			Out: RecrResp{
				Error: true,
			},
			Err: nil,
		},
	}

	for _, test := range casesToken {
		resp, err := r.recGateway.Disable(test.In)
		assert.Equal(r.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(r.T(), test.Err, err, test.Name)
		}
	}
}

func (r *RecurringTestSuite) TestFinish() {
	nameTest := "Recurring_Finish:%s"
	var casesToken = []CaseToken{
		{
			Name: fmt.Sprintf(nameTest, "OK"),
			In:   r.conf.RegisteredID,
			Out: RecrResp{
				Error: false,
			},
			Err: nil,
		},
		{
			Name: fmt.Sprintf(nameTest, "Not Found"),
			In:   randstr.String(20),
			Out: RecrResp{
				Error: true,
			},
			Err: nil,
		},
	}

	for _, test := range casesToken {
		resp, err := r.recGateway.Finish(test.In)
		assert.Equal(r.T(), test.Out.Error, resp.Error, test.Name)

		if err == nil {
			assert.Equal(r.T(), test.Err, err, test.Name)
		}
	}
}
