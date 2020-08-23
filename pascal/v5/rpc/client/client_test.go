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
