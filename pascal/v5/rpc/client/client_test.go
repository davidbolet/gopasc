package client

import (
	"fmt"
	"testing"

	"github.com/davidbolet/gopasc/pascal/v5/rpc/models"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	pascalClient *PascalClient
	myB58Key     models.HexaString
	myEncPubkey  models.HexaString
}

func (s *TestSuite) SetupTest() {
	s.pascalClient = GenerateClient("http://localhost:4003")
	s.myB58Key = "3GhhborrbqL9vEmXzH1sTfZTK87gZmdXTdS92VMmzBJaYWQZuFwz4HVVwoLmZoh1ZubJzMU9Vm1QJsT5sZM48Um7nJ3PRsDjkQSSMJ"
	s.myEncPubkey = "CA022000A436D45ADEBFC40AE7899339BF37C28CDAD29AF7B516A2E4A250471F484B2EA52000F0940F383942D8E1003CDA543381C21AC38BAE264B4E4BC98680ED7A1EDFD8AC"
}

func TestPascalClientSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestGetAccount() {
	account, err := s.pascalClient.GetAccount(819652)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().NotNil(account, "An account object should be returned")
	s.Assert().Equal(uint16(0), account.Type, "Account type should be 0")
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

func (s *TestSuite) TestGetWalletAccountsEncKey() {
	var myEncPubkey models.HexaString = s.myEncPubkey
	accounts, err := s.pascalClient.GetWalletAccounts(&myEncPubkey, nil, 0, 100)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().NotNil(accounts, "An account object should be returned")
}

func (s *TestSuite) TestGetWalletAccountsB58Key() {
	var myB58Key models.HexaString = s.myB58Key
	accounts, err := s.pascalClient.GetWalletAccounts(nil, &myB58Key, 0, 100)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().NotNil(accounts, "An account object should be returned")
}

func (s *TestSuite) TestGetWalletAccountsCountEncKey() {
	var myEncPubkey models.HexaString = s.myEncPubkey
	accountscount, err := s.pascalClient.GetWalletAccountsCount(&myEncPubkey, nil, 0, 10)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().Equal(int64(1), accountscount, "There are 0 accounts")
}

func (s *TestSuite) TestGetWalletAccountsCountB58Key() {
	var myB58Key models.HexaString = s.myB58Key
	accountscount, err := s.pascalClient.GetWalletAccountsCount(nil, &myB58Key, 0, 10)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().Equal(int64(1), accountscount, "There are 0 accounts")
}

func (s *TestSuite) TestGetWalletPubkeys() {
	getwalletpubkeys, err := s.pascalClient.GetWalletPubkeys(0, 100)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().NotNil(getwalletpubkeys, "A public key should be returned")

}

func (s *TestSuite) TestGetWalletPubkey() {
	var myB58Key models.HexaString = s.myB58Key
	//var myEncPubkey models.HexaString = s.myEncPubkey
	getwalletpubkeys, err := s.pascalClient.GetWalletPubkey(nil, &myB58Key)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().NotNil(getwalletpubkeys, "A public key should be returned")
}

func (s *TestSuite) TestGetWalletCoinsNoKey() {
	getwalletcoins, err := s.pascalClient.GetWalletCoins(nil, nil)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().Greater(getwalletcoins, float64(0), "Coins must not be 0")
}

func (s *TestSuite) TestGetWalletCoinsNonExistingKey() {
	var myB58Key models.HexaString = "3GhhborrbqL9vEmXzH1sTfZTK87gZmdXTdS92VMmzBJaYWQZuFwz4HVVwoLmZoh1ZubJzMU9Vm1QJsT5sZM48Um7nJ3PRsDjkQSSMJ"
	var myEncPubkey models.HexaString = s.myEncPubkey
	getwalletcoins, err := s.pascalClient.GetWalletCoins(&myEncPubkey, &myB58Key)
	s.Assert().Nil(err, "Error has to be nil")
	s.Assert().Equal(getwalletcoins, float64(0), "Coins must be 0")
	s.Assert().NotNil(getwalletcoins, "A public key should be returned")
}

func (s *TestSuite) TestGetBlock() {
	var block = 453489
	getblock, err := s.pascalClient.GetBlock(block)
	fmt.Printf("%+v", getblock)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetBlocks() {
	//var last = 5
	var start = 453486
	var end = 453489
	getblocks, err := s.pascalClient.GetBlocks(start, end)
	fmt.Printf("%+v", getblocks)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetBlockCount() {
	getblockcount, err := s.pascalClient.GetBlockCount()
	fmt.Printf("%+v", getblockcount)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetBlockOperation() {
	var block = 453486
	var operationBlock = 0
	getblockoperation, err := s.pascalClient.GetBlockOperation(block, operationBlock)
	fmt.Printf("%+v", getblockoperation)
	s.Assert().Nil(err, "Error has to be nil")

}

func (s *TestSuite) TestGetBlockOperations() {
	var block = 453486
	var start = 0
	var max = 100

	getblockoperations, err := s.pascalClient.GetBlockOperations(block, start, max)
	fmt.Printf("%+v", getblockoperations)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetAccountOperations() {
	var account = 0
	var depth = 100
	var start = 0
	var max = 100

	getaccountoperations, err := s.pascalClient.GetAccountOperations(account, depth, start, max)
	fmt.Printf("%+v", getaccountoperations)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetPendings() {
	var start = 0
	var max = 100

	getpendings, err := s.pascalClient.GetPendings(start, max)
	if len(*getpendings) == 0 {
		println("No pending operations")
	} else {
		fmt.Printf("%+v", getpendings)
	}
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetPendingsCount() {
	getpendingscount, err := s.pascalClient.GetPendingsCount()
	fmt.Printf("%+v", getpendingscount)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestFindOperation() {
	var ophash models.HexaString
	ophash = "71EB0600C4810C0002000000A348E5B249B15AEB916679456B2D6C87F6CBB647"
	findoperation, err := s.pascalClient.FindOperation(&ophash)
	fmt.Printf("%+v", findoperation)
	s.Assert().Nil(err, "Error has to be nil")
}
