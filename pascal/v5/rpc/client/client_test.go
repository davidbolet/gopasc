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
	fmt.Printf("%+v", account)
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
	fmt.Printf("%+v", *getaccountoperations)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestGetPendings() {
	var start = 0
	var max = 100

	getpendings, err := s.pascalClient.GetPendings(start, max)
	if len(*getpendings) == 0 {
		println("No pending operations")
	} else {
		fmt.Printf("%+v", *getpendings)
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
	fmt.Printf("%+v", *findoperation)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestFindAccounts() {
	name := "atreides"
	start := 0
	max := 100
	exact := true
	var minBalance float64 = 1
	var maxBalance float64 = 132
	//var encPubkey models.HexaString = s.myEncPubkey
	//var b58Pubkey *models.HexaString
	foundAccounts, err := s.pascalClient.FindAccounts(name, start, 0, max, exact, minBalance, maxBalance, nil, &s.myB58Key)
	fmt.Printf("%+v", *foundAccounts)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestSendTo() {
	sender := 819652
	target := 1104984
	var amount float64 = 0.01
	var fee float64 = 0.0001
	var payload models.HexaString = ""
	payloadMethod := "dest"
	pwd := ""

	sendTo, err := s.pascalClient.SendTo(sender, target, amount, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *sendTo)
	s.Assert().Nil(err, "Error has to be nil")

}

func (s *TestSuite) TestChangeKey() {
	account := 1104984
	//newEncPubkey := ""
	var newB58Pubkey models.HexaString = "3GhhbonsfmKwoT4iCgjgqNbmPi7u9CWuZPBW9jboar96Ljz3Wnfpvun9tYaNmerrrGu9n4CehDTStCZAUQ7whdKpnb4yupkwWQNfcu"
	fee := 0.0001
	var payload models.HexaString = "test"
	payloadMethod := "dest"
	pwd := ""
	ChangeKey, err := s.pascalClient.ChangeKey(account, nil, &newB58Pubkey, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *ChangeKey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestListAccountForSale() {
	accountTarget := 1104984
	accountSigner := 1104984
	price := 2000.0
	sellerAccount := 819652
	//var newB58Pubkey models.HexaString = "3GhhbonsfmKwoT4iCgjgqNbmPi7u9CWuZPBW9jboar96Ljz3Wnfpvun9tYaNmerrrGu9n4CehDTStCZAUQ7whdKpnb4yupkwWQNfcu"
	lockedUntilBlock := 454662
	fee := 0.0001
	var payload models.HexaString = "test"
	payloadMethod := "dest"
	pwd := ""
	AccoutForSale, err := s.pascalClient.ListAccountForSale(accountTarget, accountSigner, price, sellerAccount, nil, nil, fee, payload, payloadMethod, lockedUntilBlock, pwd)
	fmt.Printf("%+v", *AccoutForSale)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestDelistAccountForSale() {
	accountTarget := 1104984
	accountSigner := 1104984
	fee := 0.0001
	var payload models.HexaString = "test"
	payloadMethod := "dest"
	pwd := ""
	AccountForDelist, err := s.pascalClient.DelistAccountForSale(accountTarget, accountSigner, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *AccountForDelist)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestBuyAccount() {
	buyerAccount := 123
	accountToPurchase := 1104984
	price := 0.0001
	sellerAccount := 42
	//var newEncPubkey models.HexaString = "test"
	var newB58Pubkey models.HexaString = "dest"
	var amount float64 = 0.0001
	var fee float64 = 0.00001
	var payload models.HexaString = "asdsa"
	payloadMethod := "dest"
	pwd := ""
	BuyAccount, err := s.pascalClient.BuyAccount(buyerAccount, accountToPurchase, price, sellerAccount, nil, &newB58Pubkey, amount, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *BuyAccount)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestChangeAccountInfo() {
	accountTarget := 819652
	accountSigner := 819652
	//var newEncPubkey models.HexaString = "test"
	var newB58Pubkey models.HexaString = "3GhhborrbqL9vEmXzH1sTfZTK87gZmdXTdS92VMmzBJaYWQZuFwz4HVVwoLmZoh1ZubJzMU9Vm1QJsT5sZM48Um7nJ3PRsDjkQSSMJ"
	newName := ""
	newType := 1
	var fee float64 = 0
	var payload models.HexaString = "Testing Acc"
	payloadMethod := "aes"
	pwd := "1234"
	ChangeAccInfo, err := s.pascalClient.ChangeAccountInfo(accountTarget, accountSigner, nil, &newB58Pubkey, newName, newType, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *ChangeAccInfo)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestSignSendTo() {
	var rawOperations models.HexaString = "ASF"
	sender := 233
	target := 421
	var senderEncPubkey models.HexaString = "test"
	var senderB58Pubkey models.HexaString = "dest"
	var targetEncPubkey models.HexaString = "test"
	var targetB58Pubkey models.HexaString = "dest"

	accObject, err := s.pascalClient.GetAccount(1104984)
	s.Assert().Nil(err, "Error has to be nil")
	lastNoperation := accObject.NOperation
	var amount float64 = 0.0001
	var fee float64 = 0.00001
	var payload models.HexaString = ""
	payloadMethod := "dest"
	pwd := ""
	SignSend, err := s.pascalClient.SignSendTo(&rawOperations, sender, target, &senderEncPubkey, &senderB58Pubkey, &targetEncPubkey, &targetB58Pubkey, lastNoperation, amount, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *SignSend)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestSignChangeKey() {
	var rawOperations models.HexaString = "ASF"
	account := 2321
	//var oldEncPubkey models.HexaString = "test"
	var oldB58Pubkey models.HexaString = "dest"
	//var newEncPubkey models.HexaString = "test"
	var newB58Pubkey models.HexaString = "dest"

	accObject, _ := s.pascalClient.GetAccount(1104984)

	lastNoperation := accObject.NOperation
	var amount float64 = 0.0001
	var fee float64 = 0.00001
	var payload models.HexaString = ""
	payloadMethod := "dest"
	pwd := ""
	ChangeAccKey, err := s.pascalClient.SignChangeKey(&rawOperations, account, nil, &oldB58Pubkey, nil, &newB58Pubkey, lastNoperation, amount, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *ChangeAccKey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestSignListAccountForSale() {
	accountTarget := 523
	accountSigner := 532
	price := 43
	sellerAccount := 432
	//var newEncPubkey models.HexaString = "test"
	var newB58Pubkey models.HexaString = "dest"
	lockedUntilBlock := 54
	var fee float64 = 0.00001
	var payload models.HexaString = ""
	payloadMethod := "dest"
	pwd := ""
	ChangeAccKey, err := s.pascalClient.SignListAccountForSale(accountTarget, accountSigner, price, sellerAccount, nil, &newB58Pubkey, lockedUntilBlock, fee, payload, payloadMethod, pwd)
	fmt.Printf("%+v", *ChangeAccKey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestSignDelistAccountForSale() {
	var rawoperations models.HexaString = "as"
	//var signerEncPubkey models.HexaString = "test"
	var signerB58Pubkey models.HexaString = "dest"
	lastNoperation := 1
	ChangeAccKey, err := s.pascalClient.SignDelistAccountForSale(&rawoperations, nil, &signerB58Pubkey, lastNoperation)
	fmt.Printf("%+v", *ChangeAccKey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestSignBuyAccount() {
	buyerAccount := 123
	accountToPurchase := 1104984
	price := 0.0001
	sellerAccount := 42
	//var newEncPubkey models.HexaString = "test"
	var newB58Pubkey models.HexaString = "dest"
	var amount float64 = 0.0001
	var fee float64 = 0.00001
	var payload models.HexaString = "asdsa"
	payloadMethod := "dest"
	pwd := ""

	//var signerEncPubkey models.HexaString = "test"
	var signerB58Pubkey models.HexaString = "dest"
	lastNoperation := 13
	BuyAccount, err := s.pascalClient.SignBuyAccount(buyerAccount, accountToPurchase, price, sellerAccount, nil, &newB58Pubkey, amount, fee, payload, payloadMethod, pwd, nil, &signerB58Pubkey, lastNoperation)
	fmt.Printf("%+v", *BuyAccount)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestOperationsInfo() {
	var rawoperations models.HexaString = "asdsad"
	opInfo, err := s.pascalClient.OperationsInfo(&rawoperations)
	fmt.Printf("%+v", *opInfo)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestEncondePubkey() {
	ecNid := 714
	var x models.HexaString = "saffs"
	var y models.HexaString = "sass"
	encondePubkey, err := s.pascalClient.EncondePubkey(ecNid, x, y)
	fmt.Printf("%+v", *encondePubkey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestDecodePubkey() {
	//var encPubkey models.HexaString = "AFSS"
	var b58Pubkey models.HexaString = "AFSS"
	decodePubkey, err := s.pascalClient.DecodePubkey(nil, &b58Pubkey)
	fmt.Printf("%+v", *decodePubkey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestPayloadEncrypt() {
	var payload models.HexaString = "test"
	payloadMethod := "aes"
	//var encPubkey models.HexaString = "AFSS"
	var b58Pubkey models.HexaString = s.myB58Key
	pwd := "1234"

	decodePubkey, err := s.pascalClient.PayloadEncrypt(payload, payloadMethod, nil, &b58Pubkey, pwd)
	fmt.Printf("%+v", decodePubkey)
	s.Assert().Nil(err, "Error has to be nil")
}

func (s *TestSuite) TestPayloadDecrypt() {
	var payload models.HexaString = "c000226120"
	var pwds []string = make([]string, 1)
	pwds[0] = "1234"
	decodePubkey, err := s.pascalClient.PayloadDecrypt(payload, pwds)
	fmt.Printf("%+v", *decodePubkey)
	s.Assert().Nil(err, "Error has to be nil")
}
