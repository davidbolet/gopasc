package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	pascalClient *PascalClient
}

func (s *TestSuite) SetupTest() {
	s.pascalClient = GenerateClient("http://localhost:4003")
}

func TestPascalClientSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestGetAccount() {
	account, err := s.pascalClient.GetAccount(3532)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().NotNil(account, "An account object should be returned")
	s.Assert().Equal("0000000000000000000000000000000000000000", account.Seal, "Account Seal should be 0000000000000000000000000000000000000000")
}

func (s *TestSuite) TestGetAccountInvalidParam() {
	account, err := s.pascalClient.GetAccount(-1)
	s.Assert().Nil(err, "Error should be nil")
	s.Assert().Nil(account, "Account should be nil")
}

func (s *TestSuite) TestNodeStatus() {
	nodeStatus, err := s.pascalClient.NodeStatus()
	fmt.Printf("%+v", nodeStatus)
	s.Assert().Nil(err, "Error should be nil")
	s.Assert().NotNil(nodeStatus, "NodeStatus should not be nil")
}
