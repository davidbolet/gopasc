package client

import (
	"github.com/davidbolet/gopasc.git/pascal/v5/rpc/models"
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
