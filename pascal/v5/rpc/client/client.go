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
	return pendings, nil
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
