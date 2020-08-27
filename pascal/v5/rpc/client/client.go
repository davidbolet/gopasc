package client

import (
	"fmt"

	"github.com/davidbolet/gopasc/pascal/v5/rpc/models"
	"github.com/ybbus/jsonrpc"
)

// PascalClient represents the rpc client.
type PascalClient struct {
	RPCAddress string
	rpcClient  jsonrpc.RPCClient
}

// GenerateClient returns an operative client.
func GenerateClient(address string) *PascalClient {
	client := PascalClient{address, jsonrpc.NewClient(address)}
	return &client
}

// GetAccount Returns a JSON Object with account information including pending operations not included in blockchain yet, but affecting this account.
func (client *PascalClient) GetAccount(accountNum int) (account *models.Account, err error) {
	resp, err := client.rpcClient.Call("getaccount", map[string]interface{}{
		"account": accountNum,
	})
	if err != nil {
		return nil, err
	}
	err = resp.GetObject(&account)
	return account, nil
}

// NodeStatus Returns information of the Node in a JSON Object.
func (client *PascalClient) NodeStatus() (nodeStatus *models.NodeStatus, err error) {
	resp, err := client.rpcClient.Call("nodestatus")
	if err != nil {
		return nil, err
	}
	err = resp.GetObject(&nodeStatus)
	return nodeStatus, nil
}

// GetWalletAccounts Returns a JSON array with all wallet accounts.
func (client *PascalClient) GetWalletAccounts(encPubkey *models.HexaString, b58Pubkey *models.HexaString, start int, max int) (accounts *[]models.Account, err error) {
	var toSend = make(map[string]interface{})
	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}
	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}
	resp, err := client.rpcClient.Call("getwalletaccounts", toSend)
	if err != nil {
		return nil, err
	}
	err = resp.GetObject(&accounts)
	fmt.Printf("%+v", resp.Result)
	return accounts, nil
}

// GetWalletAccountsCount Get number of available wallet accounts (total or filtered by public key)
func (client *PascalClient) GetWalletAccountsCount(encPubkey *models.HexaString, b58Pubkey *models.HexaString, start int, max int) (accountscount int64, err error) {
	var toSend = make(map[string]interface{})
	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}
	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}

	resp, err := client.rpcClient.Call("getwalletaccountscount", toSend)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%+v", resp.Result)
	accountscount, err = resp.GetInt()
	return accountscount, nil
}

//GetWalletPubkeys Returns a JSON Array with all pubkeys of the Wallet (address)
func (client *PascalClient) GetWalletPubkeys(start int, max int) (getwalletpubkeys []models.PublicKey, err error) {
	resp, err := client.rpcClient.Call("getwalletpubkeys", map[string]interface{}{
		"start": 0,
		"max":   100,
	})
	if err != nil {
		return
	}
	getwalletpubkeys = []models.PublicKey{}
	err = resp.GetObject(&getwalletpubkeys)
	return getwalletpubkeys, nil
}

//GetWalletPubkey Returns a JSON Object with a public key if found in the Wallet
func (client *PascalClient) GetWalletPubkey(encPubkey *models.HexaString, b58Pubkey *models.HexaString) (getwalletpubkey models.PublicKey, err error) {
	var toSend = make(map[string]interface{})
	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}
	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}
	resp, err := client.rpcClient.Call("getwalletpubkey", toSend)
	if err != nil {
		return
	}
	getwalletpubkey = models.PublicKey{}
	err = resp.GetObject(&getwalletpubkey)
	return getwalletpubkey, nil
}

// GetWalletCoins Returns coins balance.
func (client *PascalClient) GetWalletCoins(encPubkey *models.HexaString, b58Pubkey *models.HexaString) (getwalletcoins float64, err error) {
	var toSend = make(map[string]interface{})
	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}
	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}
	resp, err := client.rpcClient.Call("getwalletcoins", toSend)
	if err != nil {
		return
	}
	stringwalletcoins, err := resp.GetFloat()
	getwalletcoins = stringwalletcoins
	return getwalletcoins, nil
}

//GetBlock Returns a JSON Object with a block information
func (client *PascalClient) GetBlock(block int) (Block *models.Block, err error) {
	resp, err := client.rpcClient.Call("getblock", map[string]interface{}{
		"block": block,
	})
	if err != nil {
		return nil, err
	}

	err = resp.GetObject(&Block)
	return Block, nil
}

//GetBlocks Returns a JSON Array with blocks information from "start" to "end" (or "last" n blocks) Blocks are returned in DESCENDING order
func (client *PascalClient) GetBlocks(start int, end int) (Block *[]models.Block, err error) {
	resp, err := client.rpcClient.Call("getblocks", map[string]interface{}{
		"start": start,
		"end":   end,
	})
	if err != nil {
		return nil, err
	}
	Block = &[]models.Block{}
	err = resp.GetObject(&Block)
	return Block, nil
}

// GetBlockCount Returns an Integer with blockcount of node
func (client *PascalClient) GetBlockCount() (blockcount int, err error) {
	resp, err := client.rpcClient.Call("getblockcount")
	if err != nil {
		return 1, nil
	}
	err = resp.GetObject(&blockcount)
	return blockcount, nil
}

//GetBlockOperation Returns a JSON Object with an operation inside a block
func (client *PascalClient) GetBlockOperation(block int, opblock int) (Operation *models.Operations, err error) {
	resp, err := client.rpcClient.Call("getblockoperation", map[string]interface{}{
		"block":   block,
		"opblock": opblock,
	})
	if err != nil {
		return nil, err
	}
	Operation = &models.Operations{}
	err = resp.GetObject(Operation)
	return Operation, nil
}

//GetBlockOperations Returns a JSON Array with all operations of specified block Operations are returned in DESCENDING order
func (client *PascalClient) GetBlockOperations(block int, start int, max int) (Operations *[]models.Operations, err error) {
	resp, err := client.rpcClient.Call("getblockoperations", map[string]interface{}{
		"block": block,
		"start": start,
		"max":   max,
	})
	if err != nil {
		return nil, err
	}
	Operations = &[]models.Operations{}
	err = resp.GetObject(Operations)
	return Operations, nil
}

//GetAccountOperations Return a JSON Array with "Operation Object" items. Operations made over an account Operations are returned in DESCENDING order
func (client *PascalClient) GetAccountOperations(account int, depth int, start int, max int) (AccountOperations *[]models.Operations, err error) {
	resp, err := client.rpcClient.Call("getaccountoperations", map[string]interface{}{
		"account": account,
		"depth":   depth,
		"start":   start,
		"max":     max,
	})
	if err != nil {
		return nil, err
	}
	AccountOperations = &[]models.Operations{}
	err = resp.GetObject(AccountOperations)
	return AccountOperations, nil
}

//GetPendings Return a JSON Array with "Operation Object" items with operations pending to be included at the Blockchain.
func (client *PascalClient) GetPendings(start int, max int) (pendings *[]models.Operations, err error) {
	resp, err := client.rpcClient.Call("getpendings", map[string]interface{}{
		"start": start,
		"max":   max,
	})
	if err != nil {
		return nil, err
	}
	pendings = &[]models.Operations{}
	err = resp.GetObject(pendings)
	return pendings, err
}

// GetPendingsCount Return pending opertions count
func (client *PascalClient) GetPendingsCount() (pendingsCount *models.Operations, err error) { //Doesn't display correctly
	resp, err := client.rpcClient.Call("getpendingscount")
	if err != nil {
		return nil, err
	}
	err = resp.GetObject(&pendingsCount)
	return pendingsCount, nil
}

// FindOperation Return a JSON Object in "Operation Object" format.
func (client *PascalClient) FindOperation(ophash *models.HexaString) (operation *models.Operations, err error) {
	resp, err := client.rpcClient.Call("findoperation", map[string]interface{}{
		"ophash": ophash,
	})
	if err != nil {
		return nil, err
	}

	operation = &models.Operations{}
	err = resp.GetObject(operation)
	return operation, nil
}

//FindAccounts Find accounts by name/type and returns them as an array of "Account Object"
func (client *PascalClient) FindAccounts(name string, Acctype int, start int, max int, exact bool, minBalance float64, maxBalance float64, encPubkey *models.HexaString, b58Pubkey *models.HexaString) (FoundAccounts *[]models.Account, err error) {

	var toSend = make(map[string]interface{})

	if name != "" {
		toSend["name"] = name
	}
	toSend["Acctype"] = Acctype
	toSend["start"] = start
	toSend["max"] = max
	if exact {
		toSend["exact"] = exact
	}
	if minBalance != 0 {
		toSend["min_balance"] = minBalance
	}
	if maxBalance != 0 {
		toSend["max_balance"] = maxBalance
	}
	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}
	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}
	resp, err := client.rpcClient.Call("findaccounts", toSend)
	FoundAccounts = &[]models.Account{}
	err = resp.GetObject(FoundAccounts)
	return FoundAccounts, nil
}

//SendTo Executes a transaction operation from "sender" to "target"
func (client *PascalClient) SendTo(sender int, target int, amount float64, fee float64, payload models.HexaString, payloadMethod string, pwd string) (operationInfo *models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["sender"] = sender
	toSend["target"] = target
	toSend["amount"] = amount
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}
	if err != nil {
		return nil, err
	}

	resp, err := client.rpcClient.Call("sendto", toSend)

	operationInfo = &models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//ChangeKey Executes a change key operation, changing "account" public key for a new one.
func (client *PascalClient) ChangeKey(account int, newEncPubkey *models.HexaString, newB58Pubkey *models.HexaString, fee float64, payload models.HexaString, payloadMethod string, pwd string) (operationInfo *models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["account"] = account
	if newEncPubkey != nil {
		toSend["new_enc_pubkey"] = newEncPubkey

	}

	if newB58Pubkey != nil {
		toSend["new_b58_pubkey"] = newB58Pubkey

	}
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("changekey", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//ListAccountForSale Lists an account for sale (public or private).
func (client *PascalClient) ListAccountForSale(accountTarget int, accountSigner int, price float64, sellerAccount int, newEncPubkey *models.HexaString, newB58Pubkey *models.HexaString, fee float64, payload models.HexaString, payloadMethod string, lockedUntilBlock int, pwd string) (operationInfo *models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["account_target"] = accountTarget
	toSend["account_signer"] = accountSigner
	toSend["price"] = price
	toSend["seller_account"] = sellerAccount
	if newEncPubkey != nil {
		toSend["new_enc_pubkey"] = newEncPubkey

	}

	if newB58Pubkey != nil {
		toSend["new_b58_pubkey"] = newB58Pubkey

	}
	toSend["locked_until_block"] = lockedUntilBlock
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("listaccountforsale", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//DelistAccountForSale Delist an account for sale.
func (client *PascalClient) DelistAccountForSale(accountTarget int, accountSigner int, fee float64, payload models.HexaString, payloadMethod string, pwd string) (operationInfo *models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["account_target"] = accountTarget
	toSend["account_signer"] = accountSigner
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("delistaccountforsale", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//BuyAccount Buy an account previously listed for sale (public or private).
func (client *PascalClient) BuyAccount(buyerAccount int, accountToPurchase int, price float64, sellerAccount int, newEncPubkey *models.HexaString, newB58Pubkey *models.HexaString, amount float64, fee float64, payload models.HexaString, payloadMethod string, pwd string) (operationInfo *models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["buyer_account"] = buyerAccount
	toSend["account_to_purchase"] = accountToPurchase
	toSend["price"] = price
	toSend["seller_account"] = sellerAccount
	if newEncPubkey != nil {
		toSend["new_enc_pubkey"] = newEncPubkey

	}

	if newB58Pubkey != nil {
		toSend["new_b58_pubkey"] = newB58Pubkey

	}
	toSend["amount"] = amount
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("buyaccount", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//ChangeAccountInfo Signs a change account info for cold cold wallets.
func (client *PascalClient) ChangeAccountInfo(accountTarget int, accountSigner int, newEncPubkey *models.HexaString, newB58Pubkey *models.HexaString, newName string, newType int, fee float64, payload models.HexaString, payloadMethod string, pwd string) (operationInfo *models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["account_target"] = accountTarget
	toSend["account_signer"] = accountSigner
	if newEncPubkey != nil {
		toSend["new_enc_pubkey"] = newEncPubkey

	}

	if newB58Pubkey != nil {
		toSend["new_b58_pubkey"] = newB58Pubkey

	}

	if newName != "" {
		toSend["new_name"] = newName

	}

	if newType != 0 {
		toSend["new_type"] = newType

	}
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("changeaccountinfo", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//SignSendTo Creates and signs a "Send to" operation without checking information and without transfering to the network. It's usefull for "cold wallets" that are off-line (not synchronized with the network) and only holds private keys
func (client *PascalClient) SignSendTo(rawOperations *models.HexaString, sender int, target int, SenderEncPubkey *models.HexaString, SenderB58Pubkey *models.HexaString, TargetEncPubkey *models.HexaString, TargetB58Pubkey *models.HexaString, lastNOperation int, amount float64, fee float64, payload models.HexaString, payloadMethod string, pwd string) (rawOperationInfo *models.RawOperations, err error) {
	var toSend = make(map[string]interface{})

	if rawOperations != nil {
		toSend["rawoperations"] = rawOperations

	}

	toSend["sender"] = sender
	toSend["target"] = target

	if SenderEncPubkey != nil {
		toSend["sender_enc_pubkey"] = SenderEncPubkey

	}
	if SenderB58Pubkey != nil {
		toSend["sender_b58_pubkey"] = SenderB58Pubkey

	}
	if TargetEncPubkey != nil {
		toSend["target_enc_pubkey"] = TargetEncPubkey

	}
	if TargetB58Pubkey != nil {
		toSend["target_b58_pubkey"] = TargetB58Pubkey

	}

	toSend["last_n_operation"] = lastNOperation
	toSend["amount"] = amount
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("signsendto", toSend)

	if err != nil {
		return nil, err
	}

	rawOperationInfo = &models.RawOperations{}
	err = resp.GetObject(rawOperationInfo)
	return rawOperationInfo, err
}

//SignChangeKey Creates and signs a "Change key" operation without checking information and without transfering to the network. It's usefull for "cold wallets" that are off-line (not synchronized with the network) and only holds private keys
func (client *PascalClient) SignChangeKey(rawOperations *models.HexaString, account int, OldEncPubkey *models.HexaString, OldB58Pubkey *models.HexaString, NewEncPubkey *models.HexaString, NewB58Pubkey *models.HexaString, lastNOperation int, amount float64, fee float64, payload models.HexaString, payloadMethod string, pwd string) (rawOperationInfo *models.RawOperations, err error) {
	var toSend = make(map[string]interface{})

	if rawOperations != nil {
		toSend["rawoperations"] = rawOperations

	}

	toSend["account"] = account
	if OldEncPubkey != nil {
		toSend["old_enc_pubkey"] = OldEncPubkey

	}
	if OldB58Pubkey != nil {
		toSend["old_b58_pubkey"] = OldB58Pubkey

	}
	if NewEncPubkey != nil {
		toSend["new_enc_pubkey"] = NewEncPubkey

	}
	if NewB58Pubkey != nil {
		toSend["new_b58_pubkey"] = NewB58Pubkey

	}

	toSend["last_n_operation"] = lastNOperation
	toSend["amount"] = amount
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("signchangekey", toSend)

	if err != nil {
		return nil, err
	}

	rawOperationInfo = &models.RawOperations{}
	err = resp.GetObject(rawOperationInfo)
	return rawOperationInfo, err
}

//SignListAccountForSale Signs a List Account For Sale operation useful for offline, cold wallets.
func (client *PascalClient) SignListAccountForSale(accountTarget int, accountSigner int, price int, sellerAccount int, NewEncPubkey *models.HexaString, NewB58Pubkey *models.HexaString, lockedUntilBlock int, fee float64, payload models.HexaString, payloadMethod string, pwd string) (rawOperationInfo *models.RawOperations, err error) {
	var toSend = make(map[string]interface{})

	toSend["account_target"] = accountTarget
	toSend["account_signer"] = accountSigner
	toSend["price"] = price
	toSend["seller_account"] = sellerAccount

	if NewEncPubkey != nil {
		toSend["new_enc_pubkey"] = NewEncPubkey

	}
	if NewB58Pubkey != nil {
		toSend["new_b58_pubkey"] = NewB58Pubkey

	}

	toSend["locked_until_block"] = lockedUntilBlock
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	resp, err := client.rpcClient.Call("signlistaccountforsale", toSend)

	if err != nil {
		return nil, err
	}

	rawOperationInfo = &models.RawOperations{}
	err = resp.GetObject(rawOperationInfo)
	return rawOperationInfo, err
}

//SignDelistAccountForSale  Signs a List an account for sale (public or private) for cold wallets
func (client *PascalClient) SignDelistAccountForSale(rawOperations *models.HexaString, signerB58Pubkey *models.HexaString, signerEncPubkey *models.HexaString, lastNoperation int) (rawOperationInfo *models.RawOperations, err error) {
	var toSend = make(map[string]interface{})

	if rawOperations != nil {
		toSend["rawoperations"] = rawOperations
	}

	if signerB58Pubkey != nil {
		toSend["signer_b58_pubkey"] = signerB58Pubkey

	}
	if signerEncPubkey != nil {
		toSend["signer_enc_pubkey"] = signerEncPubkey

	}

	resp, err := client.rpcClient.Call("signdelistaccountforsale", toSend)

	if err != nil {
		return nil, err
	}

	rawOperationInfo = &models.RawOperations{}
	err = resp.GetObject(rawOperationInfo)
	return rawOperationInfo, err
}

//SignBuyAccount Signs a buy operation for cold wallets.
func (client *PascalClient) SignBuyAccount(buyerAccount int, accountToPurchase int, price float64, sellerAccount int, newEncPubkey *models.HexaString, newB58Pubkey *models.HexaString, amount float64, fee float64, payload models.HexaString, payloadMethod string, pwd string, signerEncPubkey *models.HexaString, signerB58Pubkey *models.HexaString, lastNoperation int) (rawOperationInfo *models.RawOperations, err error) {
	var toSend = make(map[string]interface{})

	toSend["buyer_account"] = buyerAccount
	toSend["account_to_purchase"] = accountToPurchase
	toSend["price"] = price
	toSend["seller_account"] = sellerAccount

	if newB58Pubkey != nil {
		toSend["new_b58_pubkey"] = newB58Pubkey
	}

	if newEncPubkey != nil {
		toSend["new_enc_pubkey"] = newEncPubkey

	}

	toSend["amount"] = amount
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod
	toSend["pwd"] = fee

	if signerB58Pubkey != nil {
		toSend["signer_b58_pubkey"] = signerB58Pubkey

	}
	if signerEncPubkey != nil {
		toSend["signer_enc_pubkey"] = signerEncPubkey

	}

	toSend["last_n_operation"] = lastNoperation

	resp, err := client.rpcClient.Call("signbuyaccount", toSend)

	if err != nil {
		return nil, err
	}

	rawOperationInfo = &models.RawOperations{}
	err = resp.GetObject(rawOperationInfo)
	return rawOperationInfo, err
}

//SignChangeAccountInfo Signs a change account info for cold cold wallets.
func (client *PascalClient) SignChangeAccountInfo(accountTarget int, accountSigner int, newEncPubkey *models.HexaString, newB58Pubkey *models.HexaString, newName string, newType int, fee float64, payload models.HexaString, payloadMethod string, pwd string, rawoperations *models.HexaString, signerEncPubkey *models.HexaString, signerB58Pubkey *models.HexaString, lastNoperation int) (operationInfo *models.RawOperations, err error) {
	var toSend = make(map[string]interface{})

	toSend["account_target"] = accountTarget
	toSend["account_signer"] = accountSigner
	if newEncPubkey != nil {
		toSend["new_enc_pubkey"] = newEncPubkey

	}

	if newB58Pubkey != nil {
		toSend["new_b58_pubkey"] = newB58Pubkey

	}

	if newName != "" {
		toSend["new_name"] = newName

	}

	if newType != 0 {
		toSend["new_type"] = newType

	}
	toSend["fee"] = fee
	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}

	if rawoperations != nil {
		toSend["rawoperations"] = rawoperations

	}

	if signerEncPubkey != nil {
		toSend["signer_enc_pubkey"] = signerEncPubkey

	}

	if signerB58Pubkey != nil {
		toSend["signer_b58_pubkey"] = signerB58Pubkey

	}

	toSend["last_n_operation"] = lastNoperation

	resp, err := client.rpcClient.Call("signchangeaccountinfo", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &models.RawOperations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//OperationsInfo Returns information stored in a rawoperations param (obtained calling signchangekey or signsendto)
func (client *PascalClient) OperationsInfo(rawoperations *models.HexaString) (operationInfo *[]models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["rawoperations"] = rawoperations

	resp, err := client.rpcClient.Call("operationsinfo", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &[]models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

//ExecuteOperations Executes operations included in rawopertions param and transfers to the network. Raw operations can include "Send to" oprations or "Change key" operations.
func (client *PascalClient) ExecuteOperations(rawoperations *models.HexaString) (operationInfo *[]models.Operations, err error) {
	var toSend = make(map[string]interface{})

	toSend["rawoperations"] = rawoperations

	resp, err := client.rpcClient.Call("executeoperations", toSend)

	if err != nil {
		return nil, err
	}

	operationInfo = &[]models.Operations{}
	err = resp.GetObject(operationInfo)
	return operationInfo, err
}

// NODE STATUS?????

//EncondePubkey Encodes a public key based on params information
func (client *PascalClient) EncondePubkey(ecNid int, x models.HexaString, y models.HexaString) (encondedKey *models.HexaString, err error) {
	var toSend = make(map[string]interface{})

	toSend["ec_nid"] = ecNid
	toSend["x"] = x
	toSend["y"] = y

	resp, err := client.rpcClient.Call("encondepubkey", toSend)

	if err != nil {
		return nil, err
	}

	err = resp.GetObject(&encondedKey)
	return encondedKey, err
}

//DecodePubkey Decodes an encoded public key
func (client *PascalClient) DecodePubkey(encPubkey *models.HexaString, b58Pubkey *models.HexaString) (decodedKey *models.PublicKey, err error) {
	var toSend = make(map[string]interface{})

	if encPubkey != nil {
		toSend["enc_pubkey"] = encPubkey
	}

	if b58Pubkey != nil {
		toSend["b58_pubkey"] = b58Pubkey
	}

	resp, err := client.rpcClient.Call("decodepubkey", toSend)

	if err != nil {
		return nil, err
	}

	err = resp.GetObject(&decodedKey)
	return decodedKey, err
}

//PayloadEncrypt Encrypt a text "paylad" using "payload_method"
func (client *PascalClient) PayloadEncrypt(payload string, payloadMethod string, encPubkey *models.HexaString, b58Pubkey *models.HexaString, pwd string) (payloadEncrypted *models.HexaString, err error) {
	var toSend = make(map[string]interface{})

	toSend["payload"] = payload
	toSend["payload_method"] = payloadMethod

	if encPubkey != nil {
		toSend["enc_pubkey"] = encPubkey
	}

	if b58Pubkey != nil {
		toSend["b58_pubkey"] = b58Pubkey
	}

	if payloadMethod == "aes" {
		toSend["pwd"] = pwd
	}
	resp, err := client.rpcClient.Call("payloadencrypt", toSend)

	if err != nil {
		return nil, err
	}

	stringEncrypted, err := resp.GetString()
	payloadEncrypted1 := models.HexaString(stringEncrypted)
	return &payloadEncrypted1, err
}

//PayloadDecrypt Returns a HEXASTRING with decrypted text (a payload) using private keys in the wallet or a list of Passwords (used in "aes" encryption)
func (client *PascalClient) PayloadDecrypt(payload string, pwds []string) (payloadDecrypt *models.DecryptResult, err error) {
	var toSend = make(map[string]interface{})

	toSend["payload"] = payload
	toSend["pwds"] = pwds

	resp, err := client.rpcClient.Call("payloaddecrypt", toSend)

	if err != nil {
		return nil, err
	}

	payloadDecrypt = &models.DecryptResult{}
	err = resp.GetObject(payloadDecrypt)
	return payloadDecrypt, err
}

//GetConnections Returns a JSON Array with Connection Objects
func (client *PascalClient) GetConnections() (connections *[]models.Connection, err error) {

	resp, err := client.rpcClient.Call("getconnections")

	if err != nil {
		return nil, err
	}

	connections = &[]models.Connection{}
	err = resp.GetObject(connections)
	return connections, err
}

//AddNewKey Creates a new Private key and sotres it on the wallet, returning an enc_pubkey value
func (client *PascalClient) AddNewKey(ecNid int, name string) (newKey *models.PublicKey, err error) {
	var toSend = make(map[string]interface{})

	toSend["ec_nid"] = ecNid
	toSend["name"] = name

	resp, err := client.rpcClient.Call("addnewkey", toSend)

	if err != nil {
		return nil, err
	}
	newKey = &models.PublicKey{}
	err = resp.GetObject(newKey)
	return newKey, err
}

//Lock Locks the Wallet if it has a password, otherwise wallet cannot be locked
func (client *PascalClient) Lock() (locked bool, err error) {

	resp, err := client.rpcClient.Call("lock")

	locked, err = resp.GetBool()
	return locked, err
}

//UnLock Unlocks a locked Wallet using "pwd" param
func (client *PascalClient) UnLock(pwd string) (locked bool, err error) {
	var toSend = make(map[string]interface{})
	toSend["pwd"] = pwd
	resp, err := client.rpcClient.Call("unlock", toSend)

	locked, err = resp.GetBool()
	return locked, err
}

//SetWalletPassword Changes the password of the Wallet. (Must be previously unlocked) Note: If pwd param is empty string, then wallet will be not protected by password
func (client *PascalClient) SetWalletPassword(pwd string) (newPassword bool, err error) {
	var toSend = make(map[string]interface{})
	toSend["pwd"] = pwd
	resp, err := client.rpcClient.Call("setwalletpassword", toSend)

	newPassword, err = resp.GetBool()
	return newPassword, err
}

//StopNode Stops the node and the server. Closes all connections
func (client *PascalClient) StopNode() (nodeStop bool, err error) {

	resp, err := client.rpcClient.Call("stopnode")

	nodeStop, err = resp.GetBool()
	return nodeStop, err
}

//StartNode Starts the node and the server. Starts connection process
func (client *PascalClient) StartNode() (nodeStart bool, err error) {

	resp, err := client.rpcClient.Call("startnode")

	nodeStart, err = resp.GetBool()
	return nodeStart, err
}

//SignMessage Signs a digest message using a public key
func (client *PascalClient) SignMessage(digest string, encPubkey *models.HexaString, b58Pubkey *models.HexaString) (ToRecieve map[string]models.HexaString, err error) {
	var toSend = make(map[string]interface{})
	toSend["digest"] = digest

	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}

	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}

	resp, err := client.rpcClient.Call("signmessage", toSend)
	var toReceive = make(map[string]models.HexaString)

	err = resp.GetObject(&toReceive)
	return toReceive, err
}

//VerifySign Verify if a digest message is signed by a public key
func (client *PascalClient) VerifySign(digest string, encPubkey *models.HexaString, b58Pubkey *models.HexaString, signature models.HexaString) (receive map[string]models.HexaString, err error) {
	var toSend = make(map[string]interface{})
	toSend["digest"] = digest

	if encPubkey != nil {
		toSend["enc_pubkey"] = *encPubkey
	}

	if b58Pubkey != nil {
		toSend["b58_pubkey"] = *b58Pubkey
	}

	toSend["signature"] = signature

	resp, err := client.rpcClient.Call("verifysign", toSend)

	var toReceive = make(map[string]models.HexaString)
	err = resp.GetObject(&toReceive)
	return toReceive, err
}
