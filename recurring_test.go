package mcpayment

import (
	"github.com/stretchr/testify/suite"
)

type RecurringTestSuite struct {
	suite.Suite
	recGateway    RecurringGateway
	conf          Configs
	newRegisterID string
}
